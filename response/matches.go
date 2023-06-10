package response

import "time"

type MatchesResponse struct {
	Meta struct {
		Status int `json:"status,omitempty"`
	} `json:"meta,omitempty"`
	Data struct {
		Matches []Match `json:"matches,omitempty"`
		NextPageToken string `json:"next_page_token,omitempty"`
	} `json:"data,omitempty"`
}

type Match struct {
	Seen Seen `json:"seen,omitempty"`
	ID_                string    `json:"_id,omitempty"`
	ID                string    `json:"id,omitempty"`
	Closed            bool      `json:"closed,omitempty"`
	CommonFriendCount int       `json:"common_friend_count,omitempty"`
	CommonLikeCount   int       `json:"common_like_count,omitempty"`
	CreatedDate       time.Time `json:"created_date,omitempty"`
	Dead              bool      `json:"dead,omitempty"`
	LastActivityDate  time.Time `json:"last_activity_date,omitempty"`
	MessageCount      int       `json:"message_count,omitempty"`
	Messages          []Message `json:"messages,omitempty"`
	Participants            []string `json:"participants,omitempty"`
	Pending                 bool     `json:"pending,omitempty"`
	IsSuperLike             bool     `json:"is_super_like,omitempty"`
	IsBoostMatch            bool     `json:"is_boost_match,omitempty"`
	IsSuperBoostMatch       bool     `json:"is_super_boost_match,omitempty"`
	IsPrimetimeBoostMatch   bool     `json:"is_primetime_boost_match,omitempty"`
	IsExperiencesMatch      bool     `json:"is_experiences_match,omitempty"`
	IsFastMatch             bool     `json:"is_fast_match,omitempty"`
	IsPreferencesMatch      bool     `json:"is_preferences_match,omitempty"`
	IsMatchmakerMatch       bool     `json:"is_matchmaker_match,omitempty"`
	IsOpener                bool     `json:"is_opener,omitempty"`
	HasShownInitialInterest bool     `json:"has_shown_initial_interest,omitempty"`
	IsNewMessage            bool     `json:"is_new_message,omitempty"`
	Person           UserStruct `json:"person,omitempty"`
	Following        bool `json:"following,omitempty"`
	FollowingMoments bool `json:"following_moments,omitempty"`
	Readreceipt	ReadReceipt `json:"readreceipt,omitempty"`
	LikedContent struct {
		ByCloser struct {
			UserID      string `json:"user_id,omitempty"`
			Type        string `json:"type,omitempty"`
			IsSwipeNote bool   `json:"is_swipe_note,omitempty"`
		} `json:"by_closer,omitempty"`
	} `json:"liked_content,omitempty"`
	SubscriptionTier string `json:"subscription_tier,omitempty"`
	IsArchived       bool   `json:"is_archived"`
	SuperLiker       string `json:"super_liker,omitempty"`
}

type Seen struct {
	MatchSeen bool `json:"match_seen,omitempty"`
	LastSeenMsgID string `json:"last_seen_msg_id,omitempty"`
}

type ReadReceipt struct {
	Enabled bool `json:"enabled,omitempty"`
}