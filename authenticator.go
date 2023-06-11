package matchmaker

import (
	"errors"
	"io"
	"log"
	"time"

	bin "github.com/0xzer/matchmaker/binary"
	"github.com/0xzer/matchmaker/payload"
	response "github.com/0xzer/matchmaker/response"
	"github.com/0xzer/matchmaker/util"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	ErrAlreadyRefreshing = errors.New("already attempting to refresh")
	ErrInvalidRefreshToken = errors.New("invalid refreshToken")
)

type Authenticator struct {
	client *Client
	payloads *payload.AuthPayloads
	phone string
	refreshToken string
	refreshing bool
}

func (a *Authenticator) isRefreshing() bool {
	return a.refreshing
}

func (a *Authenticator) SetRefreshing(v bool) {
	a.refreshing = v
}

func (a *Authenticator) SetPhone(number string) {
	a.phone = number
}

func (a *Authenticator) SetRefreshToken(refreshToken string) {
	a.refreshToken = refreshToken
}

func (a *Authenticator) Reset() {
	a = &Authenticator{
		client: a.client,
		payloads: &payload.AuthPayloads{},
		refreshing: false,
	}
}

func (a *Authenticator) UpdateDeviceData(loginResult *bin.LoginResult) {
	if loginResult != nil {
		a.SetRefreshToken(loginResult.GetRefreshToken())
		a.client.Device.UserSession = util.UserSession{
			UserId: loginResult.GetUserId(),
			AuthTokenTTL: loginResult.GetAuthTokenTtl().GetValue(),
			RefreshToken: loginResult.GetRefreshToken(),
			XAuthToken: loginResult.GetAuthToken(),
			SessionId: util.RandomUUIDv4(),
			TimeElapsed: 0,
			CreatedAt: time.Now(),
		}
		err := a.client.cacheCurrentUser()
		if err != nil {
			a.client.Logger.Error().Err(err).Msg("Failed to cache current user...")
		}
		a.client.cacheLastActivityDate()
		a.client.Logger.Debug().Any("data", a.client.Device.UserSession).Msg("Successfully Updated UserSession Data")
	} else {
		a.client.Logger.Error().Msg("Invalid LoginResult")
	}
}

func (a *Authenticator) sendAuthRequest(message protoreflect.ProtoMessage) (*bin.AuthGatewayResponse, error) {
	var headers util.Headers
	headers.Build(*a.client.Device, a.client.App)
	headers.SetContentType("proto")
	payload, _ := bin.EncodeProtoMessage(message)
	//a.client.Log.Infof("LoginRequest", "payload", message)
	res, err := a.client.PostRequest(util.AUTH_LOGIN_V3,payload,headers)
	if err != nil {
		return &bin.AuthGatewayResponse{}, err
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	var responseBody bin.AuthGatewayResponse
	err = bin.DecodeProtoMessage(body, &responseBody)
	if err != nil {
		return &bin.AuthGatewayResponse{}, err
	}
	requestFailed := responseBody.GetError()
	if requestFailed != nil {
		return &bin.AuthGatewayResponse{}, &response.AuthError{Message:requestFailed.GetMessage(),Code:requestFailed.GetCode()}
	}
	//a.client.Log.Infof("LoginRequest", "response", &responseBody)
	return &responseBody, nil
}

// Phone numbers should always start with the country code representative of their originating country.
func (a *Authenticator) SendSMS(number string) (response.SentSMSResponse, error) {
	payload := a.payloads.SendSMS(number)
	res, err := a.sendAuthRequest(payload)
	if err != nil {
		return response.SentSMSResponse{}, err
	}
	phoneState := res.GetValidatePhoneOtpState()
	success := phoneState.GetSmsSent().GetValue()
	if success {
		a.phone = number
	}
	return response.SentSMSResponse{Number: phoneState.Phone, Success: success}, nil
}

// there is a possibility for this to not return the loginresult, that is if the device is not recognized and needs email validation aswell.
func (a *Authenticator) VerifyOTP(otp string) (response.VerifyOTPResponse, error) {
	payload := a.payloads.VerifyOTP(otp, a.phone)
	res, err := a.sendAuthRequest(payload)
	if err != nil {
		return response.VerifyOTPResponse{Success:false}, err
	}
	emailState := res.GetValidateEmailOtpState()
	if emailState != nil {
		refreshToken := emailState.GetRefreshToken().GetValue()
		a.SetRefreshToken(refreshToken)
		return response.VerifyOTPResponse{Success:true,EmailHint:emailState.GetMaskedEmail(),NeedsEmailValidation:true}, nil
	}
	loginResult := res.GetLoginResult()
	if loginResult != nil {
		a.UpdateDeviceData(loginResult)
		return response.VerifyOTPResponse{
			UserId: loginResult.GetUserId(),
			AuthTokenTtl: loginResult.GetAuthTokenTtl().GetValue(),
			RefreshToken: loginResult.GetRefreshToken(),
			AuthToken: loginResult.GetAuthToken(),
			NeedsEmailValidation: false,
			EmailHint: "",
			Success: true,
		}, nil
	}
	return response.VerifyOTPResponse{Success:false}, nil
}

func (a *Authenticator) ValidateEmailOTP(otp string) (response.ValidateEmailOTPResponse, error) {
	payload := a.payloads.ValidateEmailOTP(otp, a.refreshToken)
	res, err := a.sendAuthRequest(payload)
	if err != nil {
		return response.ValidateEmailOTPResponse{Success:false}, err
	}
	loginResult := res.GetLoginResult()
	a.UpdateDeviceData(loginResult)
	a.refreshToken = loginResult.GetRefreshToken()
	return response.ValidateEmailOTPResponse{
		AuthToken: loginResult.GetAuthToken(),
		RefreshToken: loginResult.GetRefreshToken(),
		UserId: loginResult.GetUserId(),
		AuthTokenTtl: loginResult.GetAuthTokenTtl().GetValue(),
		Success: true,
	}, nil
}

func (a *Authenticator) RefreshAuth() (*bin.LoginResult, error) {
	if a.isRefreshing() {
		a.client.Logger.Error().Any("refreshToken", a.refreshToken).Msg(ErrAlreadyRefreshing.Error())
		return nil, ErrAlreadyRefreshing
	}
	a.SetRefreshing(true)
	payload := a.payloads.RefreshAuth(a.refreshToken)
	res, err := a.sendAuthRequest(payload)
	if err != nil {
		a.client.Logger.Error().Any("refreshToken", a.refreshToken).Msg(err.Error())
		a.SetRefreshing(false)
		return nil, err
	}
	loginResult := res.GetLoginResult()
	if loginResult == nil {
		a.client.Logger.Error().Any("refreshToken", a.refreshToken).Msg("The refresh token is invalid.")
		return loginResult,ErrInvalidRefreshToken
	}
	a.UpdateDeviceData(loginResult)
	a.SetRefreshing(false)
	return loginResult, nil
}