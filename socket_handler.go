package matchmaker

import (
	bin "github.com/0xzer/matchmaker/binary"
)

func (c *Client) HandleSocketMessage(msgType int, data []byte) {
	if msgType == 2 { // make sure msg is binary before attempting to decode it by proto
		eventData := &bin.SocketEvents{}
		err2 := bin.DecodeProtoMessage(data, eventData)
		if err2 != nil {
			c.Logger.Err(err2).Msg("Failed to decode proto socket event")
			return
		}
		//c.Log.Infof("SocketMessage", "data", eventData)
		if c.eventHandler != nil {
			c.HandleUpdateEvent(eventData)
		} else {
			c.Logger.Error().Msg("EventHandler not set, skipping event emit")
		}
	} else {
		c.Logger.Error().Any("MsgType", msgType).Any("data", data).Msg("Unknown Event Payload")
	}
}

func (c *Client) HandleUpdateEvent(ev *bin.SocketEvents) {
	switch ev.Data.(type) {
	case *bin.SocketEvents_StartTyping:
		res := c.EventResponses.NewStartedTyping(ev.GetStartTyping())
		c.eventHandler(res)
	case *bin.SocketEvents_AppAction:
		c.HandleAppActionEvent(ev.GetAppAction())
	case *bin.SocketEvents_ConnectFromAnotherClient:
		res := c.EventResponses.NewConnectionFromAnotherClient(ev.GetConnectFromAnotherClient())
		c.eventHandler(res)
	case *bin.SocketEvents_AddedImage:
		res := c.EventResponses.NewAddedImage(ev.GetAddedImage())
		c.eventHandler(res)
	default:
		c.Logger.Info().Any("data", ev).Msg("UnknownEvent")
	}
}