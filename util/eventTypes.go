package util

import (
	"time"
	bin "github.com/0xzero/matchmaker/binary"
	dpb "google.golang.org/protobuf/types/known/durationpb"
	wpb "google.golang.org/protobuf/types/known/wrapperspb"
)

/*
 Type is oneof: "firstActivityDraw", "firstDraw", "coldStart"

 Subtype is oneof: "login" or ""

 appStart is the timestamp of when the app was first started, stored in Client
*/
func DurationMeasure(Type string, Subtype string, Category bin.DurationMeasure_Category, appStart int64) *bin.AppPlatformEvent {
	durationMeasure := &bin.DurationMeasure{
		Category: Category,
		Type: wpb.String(Type),
		Duration: dpb.New(time.Duration(appStart)),
	}
	if Subtype != "" {
		durationMeasure.Subtype = wpb.String(Subtype)
	}
	platformEvent := &bin.AppPlatformEvent{
		Value: &bin.AppPlatformEvent_DurationMeasure{
			DurationMeasure: durationMeasure,
		},
	}
	return platformEvent
}

func HubbleInstrument(ID string, Type bin.HubbleInstrumentType) *bin.AppPlatformEvent {
	hubbleInstrument := &bin.HubbleInstrument{
		Id: wpb.String(ID),
		Type: Type,
	}
	platformEvent := &bin.AppPlatformEvent{
		Value: &bin.AppPlatformEvent_HubbleInstrument{
			HubbleInstrument: hubbleInstrument,
		},
	}
	return platformEvent
}

func NewJsonEvent(eventName string, app *App, device *Device, extra *JsonEventExtra) (*bin.AppPlatformEvent, error) {
	event := &bin.JsonEvent{
		Name: wpb.String(eventName),
	}
	var baseEventData BaseEventData
	interfaceData := baseEventData.Build(app, device)
	fullPayload := JsonEventPayload{
		Schema: eventName,
		Event: interfaceData,
		Extra: *extra,
	}
	payload, err := fullPayload.MarshalJSON()
	if err != nil {
		return &bin.AppPlatformEvent{}, err
	}
	event.Value = wpb.String(string(payload))
	platformEvent := &bin.AppPlatformEvent{
		Value: &bin.AppPlatformEvent_JsonEvent{
			JsonEvent: event,
		},
	}
	return platformEvent, nil
}