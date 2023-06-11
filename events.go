package matchmaker

import (
	"io"
	"strconv"
	bin "github.com/0xzer/matchmaker/binary"
	"github.com/0xzer/matchmaker/util"

	"google.golang.org/protobuf/proto"
	tpb "google.golang.org/protobuf/types/known/timestamppb"
	wpb "google.golang.org/protobuf/types/known/wrapperspb"
)

var HubbleInstrumentEvents = util.PreAuthHubbleEvents{
	LoginFacebookTap: "f554b19e-dfcd",
	LoginGoogleTap: "4de5c7d5-a28e",
	LoginSMSTap: "f353c681-565a",
	SplashImpression: "4a354f6c-c875",
	LoginImpression: "5375cf05-3088",
	WelcomeBackImpression: "9a0d4099-f8a1",
	LoginPushTap: "677ba945-c261",
	WelcomeBackTextTap: "53e9081c-5571",
	PushAuthDeviceCheckImpression: "a3af5a55-71cf",
	Default: "1b7aaa2d-ccb7",
}

type EventPublishHeaders struct {
	ContentType string `json:"content-type" header:"content-type"`
	UserAgent string `json:"user-agent" header:"user-agent"`
}

func (c *Client) PublishEvents() (*bin.AppPublishInitialData, error) {
	eventData := &bin.InitialData{
		Data: c.EventContainer,
	}
	data, _ := proto.Marshal(eventData)
	url := util.DATA_URL_V3+"/publish/app/binary"
	headers := EventPublishHeaders{
		ContentType: "application/x-protobuf",
		UserAgent: "okhttp/4.9.1",
	}
	res, err := c.PostRequest(url, data, headers)
	if err != nil {
		return &bin.AppPublishInitialData{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return &bin.AppPublishInitialData{}, err2
	}
	res.Body.Close()
	var eventResponse bin.AppPublishInitialData
	err = proto.Unmarshal(body, &eventResponse)
    if err != nil {
		return &bin.AppPublishInitialData{}, err
    }
	return &eventResponse, err
}

func (c *Client) AddNewEvent(PlatformEvent *bin.AppPlatformEvent) {
	session := c.GetEventSession()
	ev := &bin.Event{
		Id: wpb.String(util.RandomUUIDv4()),
		CreateTime: tpb.Now(),
		EventSession: session,
		Value: &bin.Event_Platform{
			Platform: &bin.AppPlatformEvent{
				Value: PlatformEvent.Value,
			},
		},
	}
	c.EventContainer.Events = append(c.EventContainer.Events, ev)
	//log.Println(c.EventHandler)
	//d, _ := proto.Marshal(c.EventContainer)
	//os.WriteFile("payload.bin", d, os.ModePerm)
}

func (c *Client) GetEventSession() *bin.AppEventSession {
	versions := c.App.GetSemanticVersions()
	appVersion, _ := strconv.Atoi(c.App.AppVersion)
	app := &bin.AppAttributes{
		Name: &bin.AppName{
			Value: &bin.AppName_Tinder{},
		},
		Version: &bin.SemanticVersion{
			Major: wpb.Int32(versions.Major),
			Minor: wpb.Int32(versions.Minor),
			Patch: wpb.Int32(versions.Patch),
			Build: wpb.Int32(int32(appVersion)),
			Release: wpb.Bool(true),
		},
		Session: &bin.AppSession{
			SessionId: wpb.String(c.Device.Session.SessionId),
			SessionTimeElapsed: wpb.Int64(int64(c.Device.Session.TimeElapsedSeconds())),
		},
	}
	s := &bin.AppEventSession{
		App: app,
		Auth: &bin.AuthAttributes{
			AuthId: wpb.String(c.Device.DeviceId),
		},
		Device: &bin.DeviceAttributes{
			DeviceId: wpb.String(c.Device.PersistentDeviceId),
			DeviceNetworkProvider: &bin.NetworkProvider{
				Name: wpb.String(c.Device.DataProvider),
				Value: &bin.NetworkProvider_Type{
					Type: bin.NetworkType_NETWORK_TYPE_WIFI,
				},
			},
			Manufacturer: wpb.String(c.Device.Manufacturer),
			Memory: &bin.MemoryAttributes{
				TotalRamBytes: wpb.Int64(2833965056),
				FreeRamBytes: wpb.Int64(1135984640),
			},
			PersistentId: wpb.String(c.Device.PersistentDeviceId),
			Platform: &bin.DevicePlatform{
				OsVersion: wpb.String(c.Device.DeviceOsVersion),
				Type: &bin.DevicePlatform_Android{
					Android: &bin.Android{
						DeviceId: wpb.String(c.Device.PersistentDeviceId),
						InstanceId: wpb.String(c.Device.DeviceId),
						IsRooted: wpb.Bool(false),
						OsVersion: wpb.String(c.Device.DeviceOsVersion),
						PlatformVariant: bin.PlatformVariant_PLATFORM_VARIANT_GOOGLE,
						StoreVariant: bin.StoreVariant_STORE_VARIANT_PLAY_STORE,
					},
				},
			},
			Power: &bin.PowerAttributes{
				BatteryLevelPercent: wpb.Double(100),
				LowerPowerModeEnabled: wpb.Bool(false),
			},
		},
		User: &bin.UserAttributes{},
	}
	return s
}

func (c *Client) AddInitialEvents() {
	sessionEvent, _ := util.NewJsonEvent("Session.New", c.App, c.Device, &util.JsonEventExtra{
		SessionStartMethod: 1,
		SessionType: "app",
	})
	c.AddNewEvent(sessionEvent)
	accountIntroEvent, _ := util.NewJsonEvent("Account.Intro", c.App, c.Device, &util.JsonEventExtra{
		AdvertisingID: c.Device.AdvertisingId,
		Version: "authV2",
		LocaleCountry: c.App.LocaleCountry,
	})
	c.AddNewEvent(accountIntroEvent)
	identityInteractEvent, _ := util.NewJsonEvent("Identity.Interact", c.App, c.Device, &util.JsonEventExtra{
		StepName: "token_type",
		AdvertisingID: c.Device.AdvertisingId,
		Decision: "interactive",
		Origin: "client",
		AuthSessionId: c.Device.AuthSession.SessionId,
	})
	c.AddNewEvent(identityInteractEvent)
	hubbleInstrumentEvent := util.HubbleInstrument(HubbleInstrumentEvents.LoginImpression, bin.HubbleInstrumentType_HUBBLE_INSTRUMENT_TYPE_IMPRESSION)
	c.AddNewEvent(hubbleInstrumentEvent)
	firstDrawEvent := util.DurationMeasure("firstDraw", "", bin.DurationMeasure_CATEGORY_STARTUP, c.App.AppStart)
	c.AddNewEvent(firstDrawEvent)
	firstActivityDrawEvent := util.DurationMeasure("firstActivityDraw", "login", bin.DurationMeasure_CATEGORY_STARTUP, c.App.AppStart)
	c.AddNewEvent(firstActivityDrawEvent)
	identityMFAEvent, _ := util.NewJsonEvent("Identity.MFA", c.App, c.Device, &util.JsonEventExtra{
		StepName: "intro_screen",
		ActionContext: "google_facebook_sms",
		FunnelSessionId: c.Device.AuthSession.SessionId, // funnelsession is the same as auth sessionid
		FunnelVersion: "google_signin_v1.0",
		ActionName: "view",
		FunnelName: "identity_mfa",
		AdvertisingID: c.Device.AdvertisingId,
	})
	c.AddNewEvent(identityMFAEvent)
}