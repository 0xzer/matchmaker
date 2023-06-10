package util

import (
	"fmt"
	"math"
	"time"
)

type UserSession struct {
	UserId string
	AuthTokenTTL int64
	RefreshToken string
	XAuthToken string `json:"x-auth-token" header:"X-Auth-Token"`
	SessionId string`json:"userSessionId" header:"User-Session-Id"`
	TimeElapsed float64 `json:"userSessionTimeElapsed" header:"User-Session-Time-Elapsed"`
	CreatedAt time.Time
}

func (s *UserSession) ResetTime() {
	s.CreatedAt = time.Now()
}

func (s *UserSession) UpdateTimeElapsed() {
	elapsed := time.Since(s.CreatedAt).Seconds()
	s.TimeElapsed = math.Round(elapsed*1000) / 1000
}

func (s *UserSession) TimeElapsedString() string {
	s.UpdateTimeElapsed()
	seconds := int(s.TimeElapsed)
	milliseconds := int((s.TimeElapsed - float64(seconds)) * 1000)
	return fmt.Sprintf("%d.%03d", seconds, milliseconds)
}

func (s *UserSession) Reset() {
	s.SessionId = RandomUUIDv4()
	s.TimeElapsed = 0
	s.CreatedAt = time.Now()
}

type Session struct {
	SessionId    string    `json:"appSessionId" header:"App-Session-Id"`
	TimeElapsed  float64   `json:"appSessionTimeElapsed" header:"App-Session-Time-Elapsed"`
	CreatedAt    time.Time
}

func (s *Session) Reset() {
	*s = NewSession()
}

func (s *Session) ResetTime() {
	s.CreatedAt = time.Now()
}

func NewSession() Session {
	return Session{
		SessionId:   RandomUUIDv4(),
		TimeElapsed: 0,
		CreatedAt:   time.Now(),
	}
}

func (s *Session) UpdateTimeElapsed() {
	elapsed := time.Since(s.CreatedAt).Seconds()
	s.TimeElapsed = math.Round(elapsed*1000) / 1000
}

func (s *Session) TimeElapsedString() string {
	s.UpdateTimeElapsed()
	seconds := int(s.TimeElapsed)
	milliseconds := int((s.TimeElapsed - float64(seconds)) * 1000)
	return fmt.Sprintf("%d.%03d", seconds, milliseconds)
}

func (s *Session) TimeElapsedSeconds() int {
	return int(s.TimeElapsed)
}