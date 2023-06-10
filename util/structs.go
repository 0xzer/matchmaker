package util

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type App struct {
	AppVersion string
	TinderVersion string
	Platform string
	PlatformVariant string
	StoreVariant string
	Language string
	UserAgent string
	XSupportedImageFormats string
	AppStart int64
	PushEnabled bool
	LocaleCountry string
	LastActivityDate time.Time
}

type SemanticVersions struct {
	Major int32
	Minor int32
	Patch int32
}

func (a *App) GetFullAppVersion() string {
	return a.TinderVersion + " ("+a.Platform+")"
}

func (a *App) GetFullLocale() string {
	return a.Language+"_"+a.LocaleCountry
}

func (a *App) GetPlatformNum() int {
	if a.Platform == "android" {
		return 2
	}
	if a.Platform == "ios" {
		return 1
	}
	return 0
}

func (a *App) GetSemanticVersions() SemanticVersions {
	versionSplit := strings.Split(a.TinderVersion, ".")
	major, _ := strconv.Atoi(versionSplit[0])
	minor, _ := strconv.Atoi(versionSplit[1])
	patch, _ := strconv.Atoi(versionSplit[2])
	return SemanticVersions{
		Major: int32(major),
		Minor: int32(minor),
		Patch: int32(patch),
	}
}

type Device struct {
	DeviceId string
	OsVersion string
	PersistentDeviceId string
	Manufacturer string
	Model string
	DeviceOsVersion string
	DataProvider string
	AdvertisingId string
	AuthSession Session
	Session Session
	UserSession UserSession
}

func (d *Device) SaveAsJson(path string) error {
	jsonData, err := json.Marshal(d)
	if err != nil {
		log.Println("Failed to save device", err.Error())
		return err
	}
	err = os.WriteFile(path, jsonData, os.ModePerm.Perm())
	if err != nil {
		log.Println("Failed to save device", err.Error())
	}
	return err
}

type Headers struct {
	AppVersion string `header:"App-Version"`
	ContentType string `header:"Content-Type"`
	InstallId string `header:"Install-ID"`
	OsVersion string `header:"OS-Version"`
	PersistentDeviceId string `header:"Persistent-Device-ID"`
	Platform  string `header:"Platform"`
	PlatformVariant string `header:"Platform-Variant"`
	StoreVariant string `header:"Store-Variant"`
	TinderVersion string `header:"Tinder-Version"`
	UserAgent string `header:"User-Agent"`
	XSupportedImageFormats string `header:"X-Supported-Image-Formats"`
	AppSessionId string `header:"App-Session-Id"`
	AppSessionTimeElapsed string `header:"App-Session-Time-Elapsed"`
	UserSessionId string `header:"User-Session-Id"`
	UserSessionTimeElapsed string `header:"User-Session-Time-Elapsed"`
	XAuthToken string `header:"X-Auth-Token"`
	AdvertisingId string `header:"Advertising-Id"`
	FunnelSessionId string `header:"Funnel-Session-Id"`
}

func (h *Headers) Build(device Device, app *App) {
	h.InstallId = device.DeviceId
	h.PersistentDeviceId = device.PersistentDeviceId
	h.Platform = app.Platform
	h.PlatformVariant = app.PlatformVariant
	h.StoreVariant = app.StoreVariant
	h.TinderVersion = app.TinderVersion
	h.UserAgent = app.UserAgent
	h.XSupportedImageFormats = app.XSupportedImageFormats
	h.OsVersion = device.OsVersion
	h.AppVersion = app.AppVersion
	h.FunnelSessionId = device.AuthSession.SessionId
	h.AdvertisingId = device.AdvertisingId

	if len(device.UserSession.SessionId) > 0 {
		h.XAuthToken = device.UserSession.XAuthToken
		h.UserSessionId = device.UserSession.SessionId
		h.UserSessionTimeElapsed = device.UserSession.TimeElapsedString()
	}

	if len(device.Session.SessionId) > 0 {
		h.AppSessionTimeElapsed = device.Session.TimeElapsedString()
		h.AppSessionId = device.Session.SessionId
	}
}

func (h *Headers) SetContentType(contentType string) {
	if contentType == "json" {
		h.ContentType = "application/json; charset=UTF-8"
	} else if contentType == "proto" {
		h.ContentType = "application/x-protobuf"
	}
}

type JsonEventPayload struct {
	Schema string `json:"schema,omitempty"`
	Event  map[string]interface{} `json:"event,omitempty"`
	Extra JsonEventExtra `json:"extra,omitempty"`
}

func (payload JsonEventPayload) MarshalJSON() ([]byte, error) {
	eventMap := make(map[string]interface{})

	for key, value := range payload.Event {
		eventMap[key] = value
	}

	extraDataBytes, err := json.Marshal(payload.Extra)
	if err != nil {
		return nil, err
	}
	extraMap := make(map[string]interface{})
	err = json.Unmarshal(extraDataBytes, &extraMap)
	if err != nil {
		return nil, err
	}

	for key, value := range extraMap {
		eventMap[key] = value
	}

	tempPayload := struct {
		Schema string                 `json:"schema,omitempty"`
		Event  map[string]interface{} `json:"event,omitempty"`
	}{
		Schema: payload.Schema,
		Event:  eventMap,
	}

	return json.Marshal(tempPayload)
}

type BaseEventData struct {
	DataProvider          string  `json:"dataProvider,omitempty"`
	AuthID                string  `json:"authId,omitempty"`
	AndroidDeviceID       string  `json:"androidDeviceId,omitempty"`
	PersistentID          string  `json:"persistentId,omitempty"`
	PlatformVariant       string  `json:"platformVariant,omitempty"`
	Model                 string  `json:"model,omitempty"`
	Platform              int     `json:"platform,omitempty"`
	IsDebugBuild          bool    `json:"isDebugBuild"`
	OsVersion             string  `json:"osVersion,omitempty"`
	AppVersion            string  `json:"appVersion,omitempty"`
	InstanceID            string  `json:"instanceId,omitempty"`
	IsPushEnabled         bool    `json:"isPushEnabled,omitempty"`
	Language              string  `json:"language,omitempty"`
	AppSessionTimeElapsed float64 `json:"appSessionTimeElapsed,omitempty"`
	AppBuild              int     `json:"appBuild,omitempty"`
	StoreVariant          string  `json:"storeVariant,omitempty"`
	AppSessionID          string  `json:"appSessionId,omitempty"`
	Manu                  string  `json:"manu,omitempty"`
	Ts                    int64   `json:"ts,omitempty"`
}

type JsonEventExtra struct {
	AdvertisingID         string  `json:"advertisingId,omitempty"`
	Version               string  `json:"version,omitempty"`
	SessionType           string  `json:"sessionType,omitempty"`
	SessionStartMethod    int     `json:"sessionStartMethod,omitempty"`
	LocaleCountry         string  `json:"localeCountry,omitempty"`
	Origin                string  `json:"origin,omitempty"`
	Decision              string  `json:"decision,omitempty"`
	StepName              string  `json:"stepName,omitempty"`
	AuthSessionId         string  `json:"authSessionId,omitempty"`
	ActionContext         string  `json:"actionContext,omitempty"`
	FunnelSessionId       string  `json:"funnelSessionId,omitempty"`
	ActionName            string  `json:"actionName,omitempty"`
	FunnelVersion         string  `json:"funnelVersion,omitempty"`
	FunnelName            string  `json:"funnelName,omitempty"`
}

func (data *BaseEventData) Build(app *App, device *Device) map[string]interface{} {
	tempMap := make(map[string]interface{}, 0)
	device.Session.UpdateTimeElapsed()
	tempMap["appSessionTimeElapsed"] = device.Session.TimeElapsed
	tempMap["dataProvider"] = device.DataProvider
	tempMap["authId"] = device.DeviceId
	tempMap["androidDeviceId"] = device.PersistentDeviceId
	tempMap["persistentId"] = device.PersistentDeviceId
	tempMap["platformVariant"] = app.PlatformVariant
	tempMap["ts"] = GetTimestamp().UnixMilli
	tempMap["platform"] = app.GetPlatformNum()
	tempMap["storeVariant"] = app.StoreVariant
	tempMap["isPushEnabled"] = app.PushEnabled
	appBuild, _ := strconv.Atoi(app.AppVersion)
	tempMap["appBuild"] = appBuild
	tempMap["language"] = app.Language
	tempMap["appSessionId"] = device.Session.SessionId
	tempMap["osVersion"] = device.DeviceOsVersion
	tempMap["appVersion"] = app.GetFullAppVersion()
	tempMap["instanceId"] = device.DeviceId
	tempMap["isDebugBuild"] = false
	tempMap["model"] = device.Model
	tempMap["manu"] = device.Manufacturer
	return tempMap
}

type PreAuthHubbleEvents struct {
	LoginFacebookTap string
	LoginGoogleTap string
	LoginSMSTap string
	SplashImpression string
	LoginImpression string
	WelcomeBackImpression string
	LoginPushTap string
	WelcomeBackTextTap string
	PushAuthDeviceCheckImpression string
	BanImpression string
	Default string
}

type CurrentUser struct {
	UserId string
	Phone string
	Email string
	Name string
	BirthDate string
	CreateDate string
}