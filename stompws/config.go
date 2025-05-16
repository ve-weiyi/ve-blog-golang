package stompws

import (
	"time"

	"github.com/go-stomp/stomp/v3"
)

// Interface for authenticating STOMP clients.
type Authenticator interface {
	// Authenticate based on the given login and passcode, either of which might be nil.
	// Returns true if authentication is successful, false otherwise.
	Authenticate(login, passcode string) bool
}

type Config struct {
	Authenticator Authenticator // Authenticates login/passcodes. If nil no authentication is performed
	HeartBeatTime time.Duration // Preferred value for heart-beat read/write timeout, if zero, then DefaultHeartBeat.
	Log           stomp.Logger
}

func (c *Config) HeartBeat() time.Duration {
	if c.HeartBeatTime == time.Duration(0) {
		return 10 * time.Second
	}
	return c.HeartBeatTime
}

func (c *Config) Authenticate(login, passcode string) bool {
	if c.Authenticator != nil {
		return c.Authenticator.Authenticate(login, passcode)
	}

	// no authentication defined
	return true
}

func (c *Config) Logger() stomp.Logger {
	return c.Log
}
