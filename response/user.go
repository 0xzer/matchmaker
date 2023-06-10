package response

import "time"

type UserStruct struct {
	ID        string    `json:"_id,omitempty"`
	Bio       string    `json:"bio,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Gender    int       `json:"gender,omitempty"`
	Name      string    `json:"name,omitempty"`
	PingTime  time.Time `json:"ping_time,omitempty"`
	Photos    []Photos `json:"photos,omitempty"`
	HideDistance bool `json:"hide_distance,omitempty"`
	HideAge      bool `json:"hide_age,omitempty"`
}

type UserResponse struct {
	Status  int     `json:"status,omitempty"`
	Results Results `json:"results,omitempty"`
}

type Results struct {
	CommonFriends       []any         `json:"common_friends,omitempty"`
	CommonFriendCount   int           `json:"common_friend_count,omitempty"`
	SpotifyTopArtists   []any         `json:"spotify_top_artists,omitempty"`
	DistanceMi          int           `json:"distance_mi,omitempty"`
	ConnectionCount     int           `json:"connection_count,omitempty"`
	CommonConnections   []any         `json:"common_connections,omitempty"`
	Bio                 string        `json:"bio,omitempty"`
	BirthDate           time.Time     `json:"birth_date,omitempty"`
	Name                string        `json:"name,omitempty"`
	Jobs                []any         `json:"jobs,omitempty"`
	Schools             []any         `json:"schools,omitempty"`
	Teasers             []any         `json:"teasers,omitempty"`
	Gender              int           `json:"gender,omitempty"`
	ShowGenderOnProfile bool          `json:"show_gender_on_profile,omitempty"`
	BirthDateInfo       string        `json:"birth_date_info,omitempty"`
	PingTime            time.Time     `json:"ping_time,omitempty"`
	Badges              []any         `json:"badges,omitempty"`
	Photos              []Photos      `json:"photos,omitempty"`
	UserInterests       UserInterests `json:"user_interests,omitempty"`
	CommonLikes         []any         `json:"common_likes,omitempty"`
	CommonLikeCount     int           `json:"common_like_count,omitempty"`
	CommonInterests     []any         `json:"common_interests,omitempty"`
	SNumber             int64         `json:"s_number,omitempty"`
	ID                  string        `json:"_id,omitempty"`
	IsTinderU           bool          `json:"is_tinder_u,omitempty"`
}

type MetaResponse struct {
	Meta struct {
		Status int `json:"status,omitempty"`
	} `json:"meta,omitempty"`
}