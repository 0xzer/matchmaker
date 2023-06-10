package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"github.com/0xzero/matchmaker"
	"github.com/0xzero/matchmaker/debug"
	"github.com/0xzero/matchmaker/payload"
	"github.com/0xzero/matchmaker/util"
)

var cli *matchmaker.Client

func main() {
	dev, _ := os.ReadFile("device.json")
	var firstDevice *util.Device
	json.Unmarshal(dev, &firstDevice)
	//prox := "http://ip:port"
	logger := debug.NewLogger()
	// for no logging just pass zerolog.Logger{}
	cli = matchmaker.NewClient(firstDevice, logger, nil)
	cli.SetEventHandler(evHandler)
	c := make(chan os.Signal)
	input := make(chan string)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer close(input)
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			line := strings.TrimSpace(scan.Text())
			if len(line) > 0 {
				input <- line
			}
		}
	}()
	for {
		select {
		case <-c:
			log.Println("Interrupt received, exiting")
			return
		case cmd := <-input:
			if len(cmd) == 0 {
				log.Println("Stdin closed, exiting")
				return
			}
			args := strings.Fields(cmd)
			cmd = args[0]
			args = args[1:]
			go handleCmd(strings.ToLower(cmd), args)
		}
	}
}

func handleCmd(cmd string, args []string) {
	switch cmd {
	case "savedevice":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: savedevice <fileName>")
			return
		}
		path := args[0]
		err := cli.Device.SaveAsJson(path)
		if err != nil {
			cli.Logger.Err(err)
		} else {
			cli.Logger.Info().Msg("Saved device...")
		}
	case "sendsms":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: sendsms <phone_number>")
			return
		}
		phoneNumber := args[0]
		smsSent, err := cli.Authenticator.SendSMS(phoneNumber)
		if err != nil {
			cli.Logger.Err(err)
			return
		}
		cli.Logger.Info().Any("response", smsSent).Msg("SendSMS")
	case "verifyotp":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: verifyotp <otp>")
			return
		}
		otp := args[0]
		phoneVerified, err := cli.Authenticator.VerifyOTP(otp)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("response", phoneVerified).Msg("VerifyOTP")
	case "validateemailotp":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: validateemailotp <otp>")
		}
		otp := args[0]
		emailValidated, err := cli.Authenticator.ValidateEmailOTP(otp)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("response", emailValidated).Msg("ValidateEmailOTP")
	case "getprofile":
		profile, err := cli.Api.GetProfile()
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", profile).Msg("GetProfile")
	case "getchats":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: getchats <count> <page_token>")
			return
		}
		count := args[0]
		var pageToken *string
		if len(args) > 1 {
			pageToken = &args[1]
		}
		chats, err := cli.Api.GetMatches(count, pageToken)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", chats).Msg("GetChats")
	case "getmessages":
		if len(args) < 2 {
			cli.Logger.Error().Msg("Usage: getmessages <matchid> <count> <page_token>")
			return
		}
		matchId := args[0]
		count := args[1]
		var pageToken *string
		if len(args) > 2 {
			pageToken = &args[2]
		}
		messages, err := cli.Api.GetMessages(matchId, count, pageToken)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", messages).Msg("GetMessages")
	case "getuser":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: getuser <userid>")
			return
		}
		userId := args[0]
		userInfo, err := cli.Api.GetUser(userId)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", userInfo).Msg("GetUser")
	case "sendmessage":
		if len(args) < 2 {
			cli.Logger.Error().Msg("Usage: sendmessage <matchid> <...message>")
			return
		}
		matchId := args[0]
		message := args[1:]
		messagePayload := payload.MessagePayload{
			Message: strings.Join(message, " "),
		}
		messageSent, err := cli.Api.SendMessage(matchId, messagePayload)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", messageSent).Msg("SendMessage")
	case "connect":
		err := cli.Connect()
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
	case "disconnect":
		err := cli.Disconnect(nil)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
		}
	case "starttyping":
		if len(args) < 2 {
			cli.Logger.Error().Msg("Usage: starttyping <userid> <matchid>")
			return
		}
		userId := args[0]
		matchId := args[1]
		err := cli.Api.StartTyping(userId, matchId)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Debug().Msg("Started Typing!")
	case "trendinggifs":
		res, err := cli.Api.TrendingGifs()
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", res).Msg("TrendingGifs")
	case "searchgifs":
		if len(args) < 1 {
			cli.Logger.Error().Msg("Usage: searchgif <query>")
			return
		}
		query := args[0]
		res, err := cli.Api.SearchGifs(query)
		if err != nil {
			cli.Logger.Error().Msg(err.Error())
			return
		}
		cli.Logger.Info().Any("data", res).Msg("SearchGifs")
	}
}

func evHandler(rawEvt interface{}) {
	switch evt := rawEvt.(type) {
	case matchmaker.Event_NewMessage:
		msg := evt.Message
		cli.Logger.Info().
		Any("from", msg.From).
		Any("from_me", evt.IsMe()).
		Any("content", msg.Message).
		Any("message_id", msg.ID).
		Any("type", msg.Type).
		Msg("New Message!")
	case matchmaker.Event_StartedTyping:
		cli.Logger.Info().Any("data", evt).Msg("Match Started Typing!")
	case matchmaker.Event_Closed:
		cli.Logger.Info().Any("closeMsg", evt.Reason).Msg("Client connection closed.")
	case matchmaker.Event_ConnectionFromAnotherClient:
		cli.Logger.Info().Any("data", evt).Msg("Connected from another client/phone")
	case matchmaker.Event_ClientReady:
		cli.Logger.Info().Any("data", evt).Msg("Client is connected and ready.")
	case matchmaker.Event_Block:
		cli.Logger.Info().Any("data", evt).Msg("New block!")
	case matchmaker.Event_AddedImage:
		cli.Logger.Info().Any("data", evt).Msg("Added new image to profile!")
	case matchmaker.Event_NewMatch:
		cli.Logger.Info().Any("data", evt).Msg("New Match!")
	case matchmaker.Event_LikedMessage:
		cli.Logger.Info().Any("data", evt).Msg("New Liked Message!")
	default:
		cli.Logger.Info().Any("data", rawEvt).Msg("UnknownEvent")
	}
}