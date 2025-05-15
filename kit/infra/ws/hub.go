// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

// Metrics 用于收集broker指标
type Metrics struct {
	Connections      int64
	MessagesSent     int64
	MessagesReceived int64
	Errors           int64
}

type BroadcastMessage struct {
	Message []byte
	Client  *Client
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	Metrics *Metrics

	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan BroadcastMessage
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Metrics:    &Metrics{},
		clients:    make(map[*Client]bool),
		broadcast:  make(chan BroadcastMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.Metrics.Connections++
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			h.Metrics.Connections--
		case bMsg := <-h.broadcast:
			h.Metrics.MessagesSent++
			h.Broadcast(bMsg.Message)
		}
	}
}

func (h *Hub) Broadcast(msg []byte) error {
	for client := range h.clients {
		select {
		case client.send <- msg:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
	return nil
}
