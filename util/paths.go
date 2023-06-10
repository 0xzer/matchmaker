package util

import (
	"fmt"
	"net/url"
)


var MEDIA_URL_V1 = "https://imageupload.gotinder.com"
var DATA_URL_V3 = "https://data.gotinder.com/v3"
var BASE_URL_V1 = "https://api.gotinder.com"
var BASE_URL_V2 = "https://api.gotinder.com/v2"
var BASE_URL_V3 = "https://api.gotinder.com/v3"
var STATIC_URL_V1 = "https://static-assets.gotinder.com"
var KEEPALIVE_SOCKET_URL = "wss://keepalive.gotinder.com/ws"

/* Static V1 Paths*/
var PROFANITY_FILTER_V1 = STATIC_URL_V1+"/ays/areyousure_v1.tsv"


/* V2 Paths */
var BUCKETS_V2 = BASE_URL_V2+"/buckets"
var PROFILE_V2 = BASE_URL_V2+"/profile?include=offerings%2Cfeature_access%2Cpaywalls%2Caccount%2Cuser%2Cboost%2Ccampaigns%2Ccompliance%2Cemail_settings%2Cexperiences%2Cinstagram%2Clikes%2Conboarding%2Ctravel%2Cplus_control%2Cpurchase%2Creadreceipts%2Cspotify%2Csuper_likes%2Ctinder_u%2Ctop_photo%2Ctutorials%2Conboarding%2Cavailable_descriptors%2Cprofile_meter"

var MATCHES_V2_URL = BASE_URL_V2+"/matches"
type MATCHES_V2 struct {}
func (m MATCHES_V2) BuildUrl(count string, pageToken *string) string {
	params := url.Values{}
	if pageToken != nil {
		params.Add("page_token", *pageToken)
	}
	params.Add("count", count)
	uri := MATCHES_V2_URL+"?"+params.Encode()
	return uri
}

var MATCHES_MESSAGES_V2_URL = BASE_URL_V2+"/matches/%s/messages"
type MATCHES_MESSAGES_V2 struct {}
func (m MATCHES_MESSAGES_V2) BuildUrl(matchId string, count string, pageToken *string) string {
	params := url.Values{}
	if pageToken != nil {
		params.Add("page_token", *pageToken)
	}
	params.Add("count", count)
	uri := fmt.Sprintf(MATCHES_MESSAGES_V2_URL, matchId)+"?"+params.Encode()
	return uri
}

var USER_INFO_V1_URL = BASE_URL_V1+"/user/%s"
type USER_INFO_V1 struct {}
func (u USER_INFO_V1) BuildUrl(userId string) string {
	return fmt.Sprintf(USER_INFO_V1_URL, userId)
}

var SEND_MESSAGE_V1_URL = BASE_URL_V1+"/user/matches/%s"
type SEND_MESSAGE_V1 struct {}
func (m SEND_MESSAGE_V1) BuildUrl(matchId string) string {
	return fmt.Sprintf(SEND_MESSAGE_V1_URL, matchId)
}

var SEEN_MATCH_V2_URL = BASE_URL_V2+"/seen/%s"
type SEEN_MATCH_V2 struct {}
func (s SEEN_MATCH_V2) BuildUrl(matchId string) string {
	return fmt.Sprintf(SEEN_MATCH_V2_URL, matchId)
}

var TENOR_TRENDING_V1_URL = BASE_URL_V1+"/tenor/trending?locale=%s"
type TENOR_TRENDING_V1 struct {}
func (t TENOR_TRENDING_V1) BuildUrl(locale string) string {
	return fmt.Sprintf(TENOR_TRENDING_V1_URL, locale)
}

var TENOR_SEARCH_V1_URL = BASE_URL_V1+"/tenor/search"
type TENOR_SEARCH_V1 struct {}
func (t TENOR_SEARCH_V1) BuildUrl(query string, locale string) string {
	params := url.Values{}
	params.Add("query", query)
	params.Add("locale", locale)
	uri := TENOR_SEARCH_V1_URL+"?"+params.Encode()
	return uri
}

var GET_UPDATES_V1_URL = BASE_URL_V1+"/updates?is_boosting=false&boost_cursor=0"

/* V3 Paths */
var AUTH_LOGIN_V3 = BASE_URL_V3+"/auth/login"