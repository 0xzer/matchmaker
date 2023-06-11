package matchmaker

import (
	"time"

	bin "github.com/0xzer/matchmaker/binary"
	"github.com/0xzer/matchmaker/response"
)


type EventResponses struct {
	client *Client
}

type Ts struct {
	Nanos int32 `json:"nanos,omitempty"`
	Seconds int64 `json:"seconds,omitempty"`
}

type ServerTimestamps struct {
	Timestamp1 Ts `json:"timestamp1,omitempty"`
	Timestamp2 Ts `json:"timestamp2,omitempty"`
}

type Event_NewMessage struct {
	client *Client
	Message response.Message
	/*
		Additional fields that aren't included in the message itself
	*/
	MatchId                	string    `json:"_id,omitempty"`
	LastActivityDate  		time.Time `json:"last_activity_date,omitempty"`
	ReadReceipt				response.ReadReceipt `json:"readreceipt,omitempty"`
	HasShownInitialInterest bool     `json:"has_shown_initial_interest,omitempty"`
	IsNewMessage            bool     `json:"is_new_message,omitempty"`
	IsArchived       		bool   `json:"is_archived,omitempty"`
	Seen response.Seen `json:"seen,omitempty"`
}

func (m *Event_NewMessage) IsMe() bool {
	return m.Message.From == m.client.CurrentUser.UserId
}

func (e *EventResponses) NewMessage(message Event_NewMessage) Event_NewMessage {
	message.client = e.client
	return message
}

type Event_StartedTyping struct {
	client *Client
	MatchId string `json:"match_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Timestamp Ts `json:"timestamp,omitempty"`
}

func (e *EventResponses) NewStartedTyping(data *bin.StartTyping) Event_StartedTyping {
	return Event_StartedTyping{
		client: e.client,
		MatchId: data.MatchId,
		UserId: data.UserId,
		Timestamp: Ts{
			Nanos: data.Timestamp.Nanos,
			Seconds: data.Timestamp.Seconds,
		},
	}
}

type Event_Closed struct {
	Reason *string `json:"reason,omitempty"`
}

type Event_ConnectionFromAnotherClient struct {
	Timestamp Ts `json:"timestamp,omitempty"`
}

func (e *EventResponses) NewConnectionFromAnotherClient(data *bin.ConnectedOnAnotherClient) Event_ConnectionFromAnotherClient {
	return Event_ConnectionFromAnotherClient{
		Timestamp: Ts{
			Nanos: data.Timestamp.Nanos,
			Seconds: data.Timestamp.Seconds,
		},
	}
}

type Event_AddedImage struct {
	Id          string `json:"id,omitempty"`
	OriginalUrl string `json:"original_url,omitempty"`
	Images      []response.ProcessedFiles `json:"images,omitempty"`
	ImagesAdded int32  `json:"images_added,omitempty"`
}

func (e *EventResponses) NewAddedImage(data *bin.AddedImage) Event_AddedImage {
	images := make([]response.ProcessedFiles, 0)
	for _, image := range data.Images {
		images = append(images, response.ProcessedFiles{URL: image.Url,Width: int(image.Width),Height: int(image.Height)})
	}
	return Event_AddedImage{
		Id: data.Id,
		OriginalUrl: data.OriginalUrl,
		Images: images,
		ImagesAdded: data.ImagesAdded,
	}
}

type Event_ClientReady struct {
	Connected bool `json:"connected,omitempty"`
}

type Event_Block struct {
	client *Client
	MatchId string `json:"match_id,omitempty"`
}

func (e *EventResponses) NewBlock(matchId string) Event_Block {
	return Event_Block{
		client: e.client,
		MatchId: matchId,
	}
}

type Event_NewMatch struct {
	client *Client
	Match response.Match `json:"match,omitempty"`
}

func (e *EventResponses) NewMatch(match response.Match) Event_NewMatch {
	return Event_NewMatch{
		client: e.client,
		Match: match,
	}
}

type Event_LikedMessage struct {
	client *Client
	MessageID string    `json:"message_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	LikerID   string    `json:"liker_id,omitempty"`
	MatchID   string    `json:"match_id,omitempty"`
	IsLiked   bool      `json:"is_liked,omitempty"`
}

func (m *Event_LikedMessage) IsMe() bool {
	return m.LikerID == m.client.CurrentUser.UserId
}

func (e *EventResponses) NewLikedMessage(data response.LikedMessage) Event_LikedMessage {
	return Event_LikedMessage{
		client: e.client,
		MessageID: data.MessageID,
		UpdatedAt: data.UpdatedAt,
		LikerID: data.LikerID,
		MatchID: data.MatchID,
		IsLiked: data.IsLiked,
	}
}