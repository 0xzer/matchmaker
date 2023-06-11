package matchmaker

import (
	bin "github.com/0xzer/matchmaker/binary"
)

func (c *Client) HandleAppActionEvent(actionEvent *bin.AppAction) {
	if c.eventHandler == nil {
		c.Logger.Error().Msg("EventHandler not set, skipping event emit")
		return
	}
	actionType := actionEvent.GetActionType()
	if actionType == 0 {
		// just ignore this because it doesnt make any sense
		// read case 0 for more information
		return
	}
	lastActivityDate := c.App.LastActivityDate
	updates, err := c.Api.GetUpdates(lastActivityDate, true)
	//c.Log.Infof("Updates", "data", updates)
	if err != nil {
		c.Logger.Err(err).Msg("Failed to get updates from tinder")
		return
	}
	switch actionType {
		case 0:
			// this indicates that the user has seen your message
			// but it does not return anything that can be used to identify the user
		case 1: // New match
			 for _, match := range updates.Matches {
				if len(match.Person.ID) > 0 {
					newMatch := c.EventResponses.NewMatch(match)
					c.eventHandler(newMatch)
				}
			 }
		case 2: // New message (sent/received this triggers if u send a message yourself)
		    for _, match := range updates.Matches {
				for _, message := range match.Messages {
					newMsg := Event_NewMessage{
						Message: message,
						MatchId: match.ID_,
						LastActivityDate: match.LastActivityDate,
						Seen: match.Seen,
						ReadReceipt: match.Readreceipt,
						HasShownInitialInterest: match.HasShownInitialInterest,
						IsNewMessage: match.IsNewMessage,
						IsArchived: match.IsArchived,
					}
					evMessage := c.EventResponses.NewMessage(newMsg)
					c.eventHandler(evMessage)
					//c.Log.Successf("New Message", "data", evMessage)
				}
			}
		case 3: // liked a message
			 for _, message := range updates.LikedMessages {
				evMessage := c.EventResponses.NewLikedMessage(message)
				c.eventHandler(evMessage)
			 }
		case 4: // unmatch (same thing as the block feature)
			for _, block := range updates.Blocks {
				evMessage := c.EventResponses.NewBlock(block)
				c.eventHandler(evMessage)
			}
			//c.Logger.Debug().Any("socketdata", actionEvent).Msg("Unmatch")
	}
}