# MatchMaker

MatchMaker is a lightweight, easy-to-use Go library for interacting with Tinder's API.

## Features

- Simple, idiomatic interface for interacting with Tinder's API. [here](https://github.com/0xzer/matchmaker/tree/main/mdtest)
- Full Go module compatibility and support for Semantic Versioning.
- Helpful debug logging to assist with troubleshooting.

## Installation

Use the [package manager](https://golang.org/dl/) to install matchmaker.

```bash
go get github.com/0xzer/matchmaker
```

# Usage
For more usage examples you can check [here](https://github.com/0xzer/matchmaker/tree/main/mdtest), it should contain everything.
```go
package matchmakerexploring

import (
	"github.com/0xzer/matchmaker"
	"github.com/0xzer/matchmaker/debug"
)

var client *matchmaker.Client
func main() {
    // nil for new device
    // alternatively pass an empty instance of zerolog.Logger{} to disable logging
    client = matchmaker.NewClient(nil, debug.NewLogger(), nil)
    client.SetEventHandler(evHandler)

    smsSent, err := client.Authenticator.SendSMS("someNumber...")
    if err != nil {
       client.Logger.Fatal().Err(err).Msg("Failed to send sms")
    }
    client.Logger.Info().Any("data", smsSent).Msg("Sent sms")
    // there is a possibility for this to not return the loginresult, that is if the device is not recognized and needs email validation aswell.
    verifiedOTP, err2 := client.Authenticator.VerifyOTP("OTP...")
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
```

# Documentation
See the **[GoDoc page for this](https://pkg.go.dev/github.com/0xzer/matchmaker)**.
