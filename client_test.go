package matchmaker_test

import (
	"testing"
	"github.com/0xzero/matchmaker"
	"github.com/0xzero/matchmaker/debug"
)

var client *matchmaker.Client
// go test client_test.go -v
func TestClient(t *testing.T) {
	// nil for new device
	// alternatively pass an empty instance of zerolog.Logger{} to disable logging
	client = matchmaker.NewClient(nil, debug.NewLogger(), nil)
	client.SetEventHandler(evHandler)

	smsSent, err := client.Authenticator.SendSMS("someNumber...")
	if err != nil {
		client.Logger.Fatal().Err(err).Msg("Failed to send sms")
	}
    client.Logger.Info().Any("data", smsSent).Msg("Sent sms")
	verifiedOTP, err2 := client.Authenticator.VerifyOTP("OTP...") // there is a possibility for this to not return the loginresult, that is if the device is not recognized and needs email validation aswell.
    if err2 != nil {
		client.Logger.Fatal().Err(err2).Msg("Failed to verify OTP")
	}
    client.Logger.Info().Any("data", verifiedOTP).Msg("Verified OTP")

    socketErr := client.Connect()
	if socketErr != nil {
		client.Logger.Fatal().Err(socketErr).Msg("Failed to connect to socket")
	}
}

func evHandler(rawEvt interface{}) {
	switch evt := rawEvt.(type) {
        case matchmaker.Event_NewMessage:
            msg := evt.Message
            client.Logger.Info().
            Any("from", msg.From).
            Any("from_me", evt.IsMe()).
            Any("content", msg.Message).
            Any("message_id", msg.ID).
            Any("type", msg.Type).
            Msg("New Message!")
		case matchmaker.Event_ClientReady:
			client.Logger.Info().Any("data",evt).Msg("Client is ready and connected!")
    }
}