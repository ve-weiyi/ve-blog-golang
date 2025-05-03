package stompws

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

// Client 表示一个 WebSocket 客户端连接
type Client struct {
	conn      *websocket.Conn
	send      chan []byte
	topics    map[string]bool
	ClientID  string
	sessionID string
	mu        sync.Mutex

	disconnected bool
	ctx          context.Context
	cancel       context.CancelFunc
	logger       logx.Logger
}

// Send 安全地发送消息到客户端
func (c *Client) Send(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.disconnected {
		return ErrClientDisconnected
	}

	select {
	case c.send <- message:
		return nil
	case <-time.After(WriteWait):
		return ErrSendTimeout
	case <-c.ctx.Done():
		return ErrServerShuttingDown
	}
}

func (c *Client) Disconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.disconnected {
		c.disconnected = true
		c.cancel()
		close(c.send)
		c.conn.Close()
	}
}

func (c *Client) writePump(b *Broker) {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.Disconnect()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			log.Printf("--------send-------->:\n%v", string(message))

			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				b.logger.Error("Write error",
					logx.Field("clientID", c.ClientID),
					logx.Field("err", err))
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				b.logger.Error("Ping error",
					logx.Field("clientID", c.ClientID),
					logx.Field("err", err))
				return
			}

		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) readPump(b *Broker) {
	defer func() {
		b.unregister <- c
		c.Disconnect()
	}()

	c.conn.SetReadLimit(MaxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(PongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})

	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				b.logger.Error("Read error",
					logx.Field("clientID", c.ClientID),
					logx.Field("err", err))
			}
			break
		}

		log.Printf("<--------read--------:\n%v", string(message))

		switch messageType {
		case websocket.TextMessage:
			// Handle text message
		case websocket.PingMessage:
			log.Printf("Received ping message from client %s", c.ClientID)
		default:
			b.logger.Error("Unsupported message type")
		}

		b.stats.IncrementMessagesReceived()

		if len(message) > MaxMessageSize {
			b.stats.IncrementErrors()
			errorFrame := fmt.Sprintf("ERROR\nmessage:%s\n\n\x00", ErrMessageTooLarge.Error())
			c.Send([]byte(errorFrame))
			continue
		}

		stompFrame, err := parseStompFrame(string(message))
		if err != nil {
			b.stats.IncrementErrors()
			b.logger.Error("STOMP parse error",
				logx.Field("clientID", c.ClientID),
				logx.Field("err", err))
			errorFrame := fmt.Sprintf("ERROR\nmessage:%s\n\n\x00", ErrInvalidFrame.Error())
			c.Send([]byte(errorFrame))
			continue
		}

		ctx := context.WithValue(c.ctx, "clientID", c.ClientID)
		ctx = context.WithValue(ctx, "sessionID", c.sessionID)

		switch stompFrame.Command {
		case "CONNECT":
			b.handleConnect(ctx, c, stompFrame)
		case "SUBSCRIBE":
			b.handleSubscribe(ctx, c, stompFrame)
		case "UNSUBSCRIBE":
			b.handleUnsubscribe(ctx, c, stompFrame)
		case "SEND":
			b.handleSend(ctx, c, stompFrame)
		case "ACK":
			b.handleAck(ctx, c, stompFrame)
		case "NACK":
			b.handleNack(ctx, c, stompFrame)
		case "DISCONNECT":
			b.handleDisconnect(ctx, c, stompFrame)
			return
		default:
			b.handlePing(ctx, c, stompFrame)
			//b.stats.IncrementErrors()
			//b.logger.Errorw("Unhandled STOMP command",
			//	logx.Field("clientID", c.ClientID),
			//	logx.Field("command", stompFrame.Command))
			//errorFrame := fmt.Sprintf("ERROR\nmessage:Unsupported command '%s'\n\n\x00", stompFrame.Command)
			//c.Send([]byte(errorFrame))
		}
	}
}
