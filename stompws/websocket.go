package stompws

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: 5 * time.Second,
}

// HandleWebSocket 处理 WebSocket 连接
func HandleWebSocket(b *Broker, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		b.logger.Errorw("WebSocket upgrade failed",
			logx.Field("err", err))
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	client := &Client{
		conn:      conn,
		send:      make(chan []byte, DefaultMessageQueueSize),
		topics:    make(map[string]bool),
		ClientID:  r.RemoteAddr,
		sessionID: fmt.Sprintf("sess-%d", time.Now().UnixNano()),
		ctx:       ctx,
		cancel:    cancel,
		logger:    b.logger,
	}

	b.register <- client

	go client.writePump(b)
	go client.readPump(b)
}
