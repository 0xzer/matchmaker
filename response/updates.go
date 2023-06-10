package response

import "time"

type UpdatesResponse struct {
	Matches           []Match      `json:"matches,omitempty"`
	Blocks            []string        `json:"blocks,omitempty"`
	Inbox             []any        `json:"inbox,omitempty"`
	LikedMessages     []LikedMessage        `json:"liked_messages,omitempty"`
	HarassingMessages []any        `json:"harassing_messages,omitempty"`
	Lists             []any        `json:"lists,omitempty"`
	Goingout          []any        `json:"goingout,omitempty"`
	DeletedLists      []any        `json:"deleted_lists,omitempty"`
	Matchmaker        []any        `json:"matchmaker,omitempty"`
	Squads            []any        `json:"squads,omitempty"`
	LastActivityDate  time.Time    `json:"last_activity_date,omitempty"`
	PollInterval      PollInterval `json:"poll_interval,omitempty"`
}

type PollInterval struct {
	Standard   int `json:"standard,omitempty"`
	Persistent int `json:"persistent,omitempty"`
}