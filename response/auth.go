package response

//import bin "tinderlib/binary"

type SentSMSResponse struct {
	Number string
	Success bool
}

type VerifyOTPResponse struct {
	EmailHint string
	NeedsEmailValidation bool
	AuthToken string
	RefreshToken string
	UserId string
	AuthTokenTtl int64
	Success bool
}

type ValidateEmailOTPResponse struct {
	AuthToken string
	RefreshToken string
	UserId string
	AuthTokenTtl int64
	Success bool
}

type AuthError struct {
	Message string
	Code int32
}

func (e *AuthError) Error() string {
	return e.Message
}