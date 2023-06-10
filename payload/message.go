package payload

import "encoding/json"

type MessagePayload struct {
	Message string `json:"message,omitempty"`
	Type string `json:"type,omitempty"` // "gif"
	GifId string `json:"gif_id,omitempty"`
}

func (m *MessagePayload) ToJSON() ([]byte, error) {
	data, err := json.Marshal(m)
	return data,err
}