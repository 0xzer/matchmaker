package matchmaker

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	bin "github.com/0xzer/matchmaker/binary"
	"github.com/0xzer/matchmaker/payload"
	"github.com/0xzer/matchmaker/response"
	"github.com/0xzer/matchmaker/util"
)

var (
	ErrUnauthorized = errors.New("invalid authToken, try refreshing it")
)

/*
  The API struct which contains all the functions to interact with the TinderAPI
*/
type API struct {
    client *Client
	socketPayloads *payload.SocketPayloads
}

func (a *API) SendAPIRequest(url string, contentType string, method string, payload *[]byte) (*http.Response, error) {
	var headers util.Headers
	headers.Build(*a.client.Device, a.client.App)
	headers.SetContentType("json")
	var res *http.Response
	var err error
	if method == "GET" {
		res, err = a.client.GetRequest(url,headers)
	} else if method == "POST" {
		res, err = a.client.PostRequest(url,*payload,headers)
	}
	if res.StatusCode == 401 {
		return nil, ErrUnauthorized
	}
	return res, err
}

/*
  Tells the TinderAPI that the device is ready to use the API and if it should grant access to any experiments,
  this also returns every single path their API has to offer. Guessing that's just how their infrastructure works
*/
func (a *API) SendBuckets() (response.BucketResponse, error) {
	durationEvent := util.DurationMeasure("coldStart", "", bin.DurationMeasure_CATEGORY_STARTUP, a.client.App.AppStart)
	a.client.AddNewEvent(durationEvent)
	a.client.Device.Session = util.NewSession()
	payload, _ := json.Marshal(payload.Buckets{DeviceId:a.client.Device.DeviceId,Experiments:[]string{}})
	res, err := a.SendAPIRequest(util.BUCKETS_V2,"json","POST",&payload)
	if err != nil {
		return response.BucketResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.BucketResponse{}, err2
	}
	res.Body.Close()
	var responseBody response.BucketResponse
	json.Unmarshal(body,&responseBody)
	if responseBody.Data.Buckets.PushAuth == "enabled" {
		a.client.App.PushEnabled = true
	} else {
		a.client.App.PushEnabled = false
	}
	a.client.AddInitialEvents()
	return responseBody,err
}

func (a *API) GetProfile() (response.ProfileResponse, error) {
	res, err := a.SendAPIRequest(util.PROFILE_V2,"json","GET",nil)
	if err != nil {
		return response.ProfileResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.ProfileResponse{}, err2
	}
	res.Body.Close()
	var responseBody response.ProfileResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

// Count is max 100 and min 1
func (a *API) GetMatches(count string, pageToken *string) (response.MatchesResponse, error) {
	uri := util.MATCHES_V2{}
	res, err := a.SendAPIRequest(uri.BuildUrl(count, pageToken),"json","GET",nil)
	if err != nil {
		return response.MatchesResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.MatchesResponse{}, err2
	}
	res.Body.Close()
	var responseBody response.MatchesResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) GetMessages(matchId string, count string, pageToken *string) (response.MessagesResponse, error) {
	uri := util.MATCHES_MESSAGES_V2{}
	res, err := a.SendAPIRequest(uri.BuildUrl(matchId, count, pageToken), "json", "GET", nil)
	if err != nil {
		return response.MessagesResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.MessagesResponse{}, err2
	}
	res.Body.Close()
	var responseBody response.MessagesResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) GetUser(userId string) (response.UserResponse, error) {
	uri := util.USER_INFO_V1{}
	res, err := a.SendAPIRequest(uri.BuildUrl(userId), "json", "GET", nil)
	if err != nil {
		return response.UserResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.UserResponse{}, err2
	}
	res.Body.Close()
	var responseBody response.UserResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) TrendingGifs() (response.GifResponse, error) {
	uri := util.TENOR_TRENDING_V1{}
	res, err := a.SendAPIRequest(uri.BuildUrl(a.client.App.GetFullLocale()), "json", "GET", nil)
	if err != nil {
		return response.GifResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.GifResponse{}, err2
	}
	res.Body.Close()
	var responseBody response.GifResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) SearchGifs(query string) (response.GifSearchResponse, error) {
	uri := util.TENOR_SEARCH_V1{}
	res, err := a.SendAPIRequest(uri.BuildUrl(query, a.client.App.GetFullLocale()), "json", "GET", nil)
	if err != nil {
		return response.GifSearchResponse{}, err
	}
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		return response.GifSearchResponse{}, err2
	}
	log.Println(string(body))
	res.Body.Close()
	var responseBody response.GifSearchResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) SendMessage(matchId string, messagePayload payload.MessagePayload) (response.Message, error) {
	uri := util.SEND_MESSAGE_V1{}
	payload, err := messagePayload.ToJSON()
	if err != nil {
		return response.Message{}, err
	}
	res, err2 := a.SendAPIRequest(uri.BuildUrl(matchId), "json", "POST", &payload)
	if err2 != nil {
		return response.Message{}, err2
	}
	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return response.Message{}, err3
	}
	res.Body.Close()
	var responseBody response.Message
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) SeenMatch(matchId string) (response.MetaResponse, error) {
	uri := util.SEEN_MATCH_V2{}
	res, err2 := a.SendAPIRequest(uri.BuildUrl(matchId), "json", "POST", nil)
	if err2 != nil {
		return response.MetaResponse{}, err2
	}
	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return response.MetaResponse{}, err3
	}
	res.Body.Close()
	var responseBody response.MetaResponse
	json.Unmarshal(body,&responseBody)
	return responseBody, nil
}

func (a *API) StartTyping(userId string, matchId string) error {
	if !a.client.KeepAliveSocket.IsConnected() {
		a.client.Logger.Err(ErrSocketClosed).Msg("Socket is closed, can't send starttyping payload")
		return ErrSocketClosed
	}
	payload, err := a.socketPayloads.StartTypingPayload(userId, matchId)
	if err != nil {
		a.client.Logger.Err(err).Msg("Failed to parse startTyping payload")
		return err
	}
	err = a.client.KeepAliveSocket.SendData(2, payload)
	if err != nil {
		a.client.Logger.Err(err).Msg("Failed to send startTyping payload to socket.")
		return err
	}
	return nil
}

func (a *API) GetUpdates(lastActivityDate time.Time, nudge bool) (response.UpdatesResponse, error) {
	date := util.FormatDateActivity(lastActivityDate)
	payload := payload.GetUpdates{
		LastActivityDate: date,
		Nudge: nudge,
	}
	data, err1 := json.Marshal(payload)
	if err1 != nil {
		return response.UpdatesResponse{}, err1
	}
	res, err2 := a.SendAPIRequest(util.GET_UPDATES_V1_URL, "json", "POST", &data)
	if err2 != nil {
		return response.UpdatesResponse{}, err2
	}
	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return response.UpdatesResponse{}, err3
	}
	res.Body.Close()
	var responseBody response.UpdatesResponse
	json.Unmarshal(body,&responseBody)
	//a.client.Log.Successf("GetUpdates", "response", responseBody)
	if nudge {
		a.client.App.LastActivityDate = responseBody.LastActivityDate	
	}
	return responseBody, nil
}