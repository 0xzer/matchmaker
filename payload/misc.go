package payload

type Buckets struct {
	DeviceId string `json:"device_id"`
	Experiments []string `json:"experiments"`
}

type GetUpdates struct {
	Nudge            bool      `json:"nudge,omitempty"`
	LastActivityDate string `json:"last_activity_date,omitempty"`
}