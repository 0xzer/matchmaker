package response

import "time"

type MessagesResponse struct {
	Meta struct {
		Status int `json:"status,omitempty"`
	} `json:"meta,omitempty"`
	Data struct {
		Messages []Message `json:"messages,omitempty"`
		NextPageToken string `json:"next_page_token,omitempty"`
	} `json:"data,omitempty"`
}

type Message struct {
	ID          string    `json:"_id,omitempty"`
	Match_Id     string    `json:"match_id,omitempty"`
	SentDate    time.Time `json:"sent_date,omitempty"`
	Message     string    `json:"message,omitempty"`
	To          string    `json:"to,omitempty"`
	From        string    `json:"from,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	Timestamp   int64     `json:"timestamp,omitempty"`
	MatchId    string    `json:"matchId,omitempty"`
	Type        string    `json:"type,omitempty"`
	ContactCard struct {
		ContactID   string `json:"contact_id,omitempty"`
		ContactType string `json:"contact_type,omitempty"`
		Deeplink    string `json:"deeplink,omitempty"`
	} `json:"contact_card,omitempty"`
}

type LikedMessage struct {
	MessageID string    `json:"message_id,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	LikerID   string    `json:"liker_id,omitempty"`
	MatchID   string    `json:"match_id,omitempty"`
	IsLiked   bool      `json:"is_liked,omitempty"`
}