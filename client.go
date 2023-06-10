package matchmaker

import (
	"log"
	"net/http"
	"net/url"
	"time"

	bin "github.com/0xzero/matchmaker/binary"
	"github.com/0xzero/matchmaker/payload"
	"github.com/0xzero/matchmaker/util"
	"github.com/rs/zerolog"
	wpb "google.golang.org/protobuf/types/known/wrapperspb"
)

type Proxy func(*http.Request) (*url.URL, error)
type EventHandler func(evt interface{})
type Client struct {
	App *util.App
	Device *util.Device
	Api *API
	KeepAliveSocket *KeepAliveSocket
	EventResponses *EventResponses
	eventHandler EventHandler
	Authenticator *Authenticator
	EventContainer *bin.AppEventData
	http *http.Client
	Logger *zerolog.Logger
	proxy Proxy
	CurrentUser *util.CurrentUser
}


/*
	NewClient is a constructor function for creating a new instance of the Client struct.

	Parameters:
		- device: a pointer to a util.Device instance. If this is nil, a new device template is created. If the device is not nil, the AuthSession, Session, and UserSession of the device are reset. If the UserSession contains a UserId or RefreshToken, these are also reset and the RefreshToken is assigned to the authenticator.
		- logger: an instance of zerolog.Logger that will be used for logging information and errors in the client. Pass an empty zerolog.Logger instance to disable logging completely.
		- proxy: a string pointer representing a proxy address. If this is not nil, the client will try to set this proxy address. If the proxy address cannot be set, an error is logged and the function continues.

	The function constructs a new Client instance with the following properties:
		- App: an instance of util.App with predefined properties.
		- KeepAliveSocket: an empty instance of KeepAliveSocket.
		- Device: the passed in device or a new device template if the passed in device was nil.
		- Api: an empty instance of API.
		- EventResponses: an empty instance of EventResponses.
		- Authenticator: the authenticator constructed earlier in the function.
		- EventContainer: an instance of bin.AppEventData with a BatchId set to a new UUIDv4 and an empty events array.
		- http: an instance of http.Client.
		- Logger: a pointer to the passed in logger.

	After creating the Client instance, the function assigns the Client instance to the client properties of Api, KeepAliveSocket, Authenticator, and EventResponses.

	Then the function logs the initialized device and checks if a proxy was passed in. If so, it tries to set the proxy. If setting the proxy fails, an error is logged and the program closes.

	Next, the function sends a buckets request. If this fails, an error is logged and the program closes.
	
	The function then tries to publish events. If this fails, an error is logged and the program closes.

	Then the function tries to cache the current user. If this fails, an error is logged, and the program closes.

	The function then caches the last activity date and returns the Client instance.

	Return:
		- a pointer to the newly constructed Client instance.
*/
func NewClient(device *util.Device, logger zerolog.Logger, proxy *string) *Client {
	authenticator := &Authenticator{refreshing:false}
	if device == nil {
		device = util.NewDeviceTemplate()
	} else {
		device.AuthSession.Reset()
		device.Session.Reset()
		if len(device.UserSession.UserId) > 0 {
			device.UserSession.Reset()
		}
		if len(device.UserSession.RefreshToken) > 0 {
			authenticator.refreshToken = device.UserSession.RefreshToken
		}
	}
	api := &API{socketPayloads:&payload.SocketPayloads{}}
	keepAliveSocket := &KeepAliveSocket{}
	eventResponses := &EventResponses{}
	client := &Client{
		App: &util.App{
			AppVersion: "4467",
			TinderVersion: "14.9.0",
			StoreVariant: "Play-Store",
			Platform: "android",
			PlatformVariant: "Google-Play",
			UserAgent: "Tinder Android Version 14.9.0",
			XSupportedImageFormats: "webp",
			Language: "en",
			LocaleCountry: "GB",
			AppStart: time.Now().Unix(),
		},
		KeepAliveSocket: keepAliveSocket,
		Device: device,
		Api: api,
		EventResponses: eventResponses,
		Authenticator: authenticator,
		EventContainer: &bin.AppEventData{
			BatchId: wpb.String(util.RandomUUIDv4()),
			Events: make([]*bin.Event, 0),
		},
		http: &http.Client{},
		Logger: &logger,
	}
	api.client = client
	keepAliveSocket.client = client
	authenticator.client = client
	eventResponses.client = client
	client.Logger.Debug().Any("device", client.Device).Msg("Initialized TinderClient ("+client.App.UserAgent+")")
	if proxy != nil {
		proxErr := client.SetProxy(*proxy)
		if proxErr != nil {
			client.Logger.Fatal().Err(proxErr).Msg("Failed to set proxy")
		}
	}
	buckets, bucketErr := api.SendBuckets()
	if bucketErr != nil {
		client.Logger.Fatal().Err(bucketErr).Msg("Failed to send buckets request")
	}
	client.Logger.Info().Any("status", buckets.Meta.Status).Msg("Buckets Status")
	eventPublishResponse, err := client.PublishEvents()
	if err != nil {
		client.Logger.Fatal().Any("res", eventPublishResponse).Err(err).Msg("Failed to publish events.")
	}
	unauthorizedErr := client.cacheCurrentUser()
	if unauthorizedErr != nil {
		client.Logger.Fatal().Err(unauthorizedErr).Msg("Failed to fetch user information, try logging in again")
	}
	client.Logger.Info().Any("response", eventPublishResponse.Response).Msg("EventPublish")
	client.cacheLastActivityDate()
	return client
}

func (c *Client) SetEventHandler(evHandler EventHandler) {
	c.eventHandler = evHandler
}

func (c *Client) Connect() error {
	err := c.KeepAliveSocket.Connect()
	if err != nil {
		c.Logger.Error().Any("errMsg", err.Error()).Msg("Failed to connect to socket")
	}
	if c.eventHandler != nil && err == nil {
		c.eventHandler(Event_ClientReady{Connected: true})
	}
	return err
}

func (c *Client) Disconnect(reason *string) error {
	err := c.KeepAliveSocket.conn.Close()
	if err != nil {
		c.Logger.Error().Any("errMsg", err.Error()).Msg("Failed to close socket connection")
		return err
	}
	if c.eventHandler != nil && err == nil {
		c.KeepAliveSocket = NewKeepAliveSocket(c)
		c.eventHandler(Event_Closed{Reason: reason})
	}
	return err
}

func (c *Client) cacheLastActivityDate() {
	if len(c.Device.UserSession.XAuthToken) > 0 {
		res, err := c.Api.GetUpdates(time.Now(), false)
		if err != nil {
			c.Logger.Error().Any("full", err).Msg(err.Error())
			return
		}
		c.App.LastActivityDate = res.LastActivityDate
		c.Logger.Debug().Any("value", res.LastActivityDate).Msg("Updated LastActivityDate Cache")
	}
}

func (c *Client) cacheCurrentUser() error {
	if len(c.Device.UserSession.XAuthToken) > 0 {
		profile, err := c.Api.GetProfile()
		if err != nil {
			c.Logger.Error().Any("full", err).Msg(err.Error())
			return err
		}
		c.CurrentUser = &util.CurrentUser{
			UserId: profile.Data.User.ID,
			Phone: profile.Data.Account.AccountPhoneNumber,
			Email: profile.Data.Account.AccountEmail,
			Name: profile.Data.User.Name,
			BirthDate: profile.Data.User.BirthDate.String(),
			CreateDate: profile.Data.User.CreateDate.String(),
		}
		c.Logger.Debug().Any("data", &c.CurrentUser).Msg("Logged In As")
	}
	return nil
}

// http://host:port
func (c *Client) SetProxy(proxy string) error {
	proxyParsed, err := url.Parse(proxy)
	if err != nil {
		log.Fatal(err)
	}
	proxyUrl := http.ProxyURL(proxyParsed)
	c.http.Transport = &http.Transport{
		Proxy: proxyUrl,
	}
	c.proxy = proxyUrl
	c.Logger.Debug().Any("proxy", proxyParsed.Host).Msg("SetProxy")
	return nil
}