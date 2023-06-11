package matchmaker

import (
	"errors"
	"net/http"
	"time"
	"github.com/0xzer/matchmaker/util"
	"github.com/gorilla/websocket"
)

var (
	ErrSocketClosed      = errors.New("socket is closed")
	ErrSocketAlreadyOpen = errors.New("socket is already open")
)

type KeepAliveSocket struct {
	conn *websocket.Conn
	client *Client
	reconnectAttempts int32
}

func NewKeepAliveSocket(client *Client) *KeepAliveSocket {
	return &KeepAliveSocket{
		conn: nil,
		client: client,
	}
}

func (socket *KeepAliveSocket) Connect() error {
	if socket.conn != nil {
		socket.client.Logger.Err(ErrSocketAlreadyOpen).Msg("Socket already open")
		return ErrSocketAlreadyOpen
	}
	dialer := websocket.Dialer{}
	if socket.client.proxy != nil {
		dialer.Proxy = socket.client.proxy
	}
	var headers util.Headers
	headers.Build(*socket.client.Device, socket.client.App)
	wsHeaders := &http.Header{}
	SetHeaders(wsHeaders, headers)

	conn, _, err := dialer.Dial(util.KEEPALIVE_SOCKET_URL,*wsHeaders)
	if err != nil {
		socket.client.Logger.Err(err).Msg("Failed to connect to socket")
		return err
	}
	conn.SetCloseHandler(socket.CloseHandler)
	socket.conn = conn
	go socket.beginReadStream()
	return nil
}

func (s *KeepAliveSocket) beginReadStream() {
	s.client.Logger.Debug().Any("connected", true).Msg("Reading Data From Socket")
	for {
		msgType, msg, err := s.conn.ReadMessage()
		if err != nil {
			s.client.Logger.Err(err).Msg("Failed to read data from socket")
			return
		}
		s.client.HandleSocketMessage(msgType, msg)
	}
}

func (s *KeepAliveSocket) SendData(msgType int, data []byte) error {
	s.client.Logger.Debug().Any("data", data).Msg("Attempting to send data to websocket")
	err := s.conn.WriteMessage(msgType, data)
	if err != nil {
		s.client.Logger.Err(err).Msg("Failed to send data to socket")
		return err
	}
	s.client.Logger.Debug().Any("data", data).Msg("Successfully sent data to websocket")
	return nil
}

func (s *KeepAliveSocket) IsConnected() bool {
	return s.conn != nil
}

func (s *KeepAliveSocket) CloseHandler(code int, text string) error {
	s.client.Logger.Info().Any("code", code).Any("text", text).Msg("Socket Closed")
	s.conn = nil
	s.reconnectAttempts = 0
	s.client.Logger.Info().Msg("Attempting to reconnect to socket...")
	reconnectedErr := s.Connect()
	for reconnectedErr != nil {
		s.reconnectAttempts += 1
		s.client.Logger.Err(reconnectedErr).Msg("Reconnection Error")
		s.client.Logger.Info().Any("attempts", s.reconnectAttempts).Msg("Failed to reconnect, attempting to again in 10 seconds")
		time.Sleep(time.Second*10)
		reconnectedErr = s.Connect()
	}
	return nil
}