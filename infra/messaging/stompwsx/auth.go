package stompwsx

import (
	"context"
	"fmt"
	"time"

	"github.com/go-stomp/stomp/v3/frame"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/stompws/server/client"
)

// tokenValidateTimeout 限定连接认证时读取权威令牌的最长等待时间。
const tokenValidateTimeout = 2 * time.Second

// Authenticator 用 tokenx 校验 STOMP CONNECT 帧里的 access token。
//
// allowGuest 为真时，login 与 passcode 均为空的连接按游客放行；
// 为假时两者必填。两种模式下 client 头都必填，它是连接的设备标识。
type Authenticator struct {
	tokens     tokenx.Service
	allowGuest bool
}

func NewAuthenticator(tokens tokenx.Service, allowGuest bool) *Authenticator {
	return &Authenticator{
		tokens:     tokens,
		allowGuest: allowGuest,
	}
}

// Authenticate implements the Authenticator interface.
func (a *Authenticator) Authenticate(c *client.Conn, f *frame.Frame) (string, string, error) {
	login := f.Header.Get("login")
	passcode := f.Header.Get("passcode")
	clientID := f.Header.Get("client")

	if clientID == "" {
		return "", "", fmt.Errorf("stomp auth failed: missing header: 'client'")
	}

	if a.allowGuest && login == "" && passcode == "" {
		return clientID, "", nil
	}
	if login == "" {
		return "", "", fmt.Errorf("stomp auth failed: missing header: 'login'")
	}
	if passcode == "" {
		return "", "", fmt.Errorf("stomp auth failed: missing header: 'passcode'")
	}

	ctx, cancel := context.WithTimeout(context.Background(), tokenValidateTimeout)
	defer cancel()

	claims, err := a.tokens.Validate(ctx, passcode)
	if err != nil {
		return "", "", fmt.Errorf("stomp auth failed: %v", err)
	}
	if claims.UserID != login || claims.DeviceID != clientID {
		return "", "", fmt.Errorf("stomp auth failed: token identity mismatch")
	}

	return clientID, login, nil
}
