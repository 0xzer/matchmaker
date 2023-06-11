package payload

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	wpb "google.golang.org/protobuf/types/known/wrapperspb"
	bin "github.com/0xzer/matchmaker/binary"
)

type AuthPayloads struct {}

func (p *AuthPayloads) SendSMS(number string) protoreflect.ProtoMessage {
	return &bin.AuthGatewayRequest{
		Factor: &bin.AuthGatewayRequest_Phone{
			Phone: &bin.Phone{
				Phone: number,
			},
		},
	}
}

func (p *AuthPayloads) VerifyOTP(otp string, phone string) protoreflect.ProtoMessage {
	return &bin.AuthGatewayRequest{
		Factor: &bin.AuthGatewayRequest_PhoneOtp{
			PhoneOtp: &bin.PhoneOtp{
				Phone: wpb.String(phone),
				Otp: otp,
			},
		},
	}
}

func (p *AuthPayloads) ValidateEmailOTP(otp string, refreshToken string) protoreflect.ProtoMessage {
	return &bin.AuthGatewayRequest{
		Factor: &bin.AuthGatewayRequest_EmailOtp{
			EmailOtp: &bin.EmailOtp{
				Otp: otp,
				RefreshToken: wpb.String(refreshToken),
			},
		},
	}
}

func (p *AuthPayloads) RefreshAuth(refreshToken string) protoreflect.ProtoMessage {
	return &bin.AuthGatewayRequest{
		Factor: &bin.AuthGatewayRequest_RefreshAuth{
			RefreshAuth: &bin.RefreshAuth{
				RefreshToken: refreshToken,
			},
		},
	}
}