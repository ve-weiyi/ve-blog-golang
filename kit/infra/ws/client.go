// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub  *Hub
	conn *websocket.Conn
	send chan []byte

	ClientID    string
	ConnectedAt time.Time
	Claims      map[string]string

	handleConnect    func() error
	handleDisconnect func() error
	handleMessage    func(message []byte) error
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*Client, error) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{"stomp", "mqtt", "v12.stomp", "v11.stomp", "v10.stomp"},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn:        conn,
		send:        make(chan []byte, maxMessageSize),
		ClientID:    r.RemoteAddr,
		ConnectedAt: time.Now(),
		handleConnect: func() error {
			return nil
		},
		handleDisconnect: func() error {
			return nil
		},
		handleMessage: func(message []byte) error {
			return nil
		},
	}

	return client, nil
}

func (c *Client) SetHandleConnect(handleConnect func() error) {
	c.handleConnect = handleConnect
}

func (c *Client) SetHandleDisconnect(handleDisconnect func() error) {
	c.handleDisconnect = handleDisconnect
}

func (c *Client) SetHandleMessage(handleMessage func(message []byte) error) {
	c.handleMessage = handleMessage
}

func (c *Client) Connect(h *Hub) error {
	c.Hub = h
	c.Hub.register <- c

	go c.readPump()
	go c.writePump()
	return c.handleConnect()
}

func (c *Client) Disconnect() error {
	c.Hub.unregister <- c
	close(c.send)
	c.conn.Close()
	return c.handleDisconnect()
}

func (c *Client) readPump() {
	defer func() {
		c.Disconnect()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	},
	)
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			log.Printf("error: %v", err)
			break
		}

		fmt.Println("<-------------------")
		fmt.Println(string(message))
		fmt.Println("<-------------------")
		// 处理消息
		c.handleMessage(message)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			fmt.Println("------------------->")
			fmt.Println(string(message))
			fmt.Println("------------------->")

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("error: %v", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) Send(msg []byte) error {
	c.send <- msg
	return nil
}
