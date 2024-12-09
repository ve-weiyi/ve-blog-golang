package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type ClientManager interface {
	// 注册
	RegisterClient(client *Client) (err error)
	// 注销
	UnRegisterClient(client *Client) (err error)
	// 在线人数
	GetOnlineCount() int

	// 广播
	PushBroadcast(msg any) (err error)
	// 单发
	PushMessage(clientId string, msg any) (err error)
}

// DefaultClientManager
type DefaultClientManager struct {
	Clients map[string]*Client
}

func NewDefaultClientManager() *DefaultClientManager {
	return &DefaultClientManager{
		Clients: make(map[string]*Client),
	}
}

func (m *DefaultClientManager) RegisterClient(client *Client) (err error) {
	if _, ok := m.Clients[client.ClientId]; ok {
		return
	}

	m.Clients[client.ClientId] = client
	return nil
}

func (m *DefaultClientManager) UnRegisterClient(client *Client) (err error) {
	if _, ok := m.Clients[client.ClientId]; ok {
		delete(m.Clients, client.ClientId)
		_ = client.WsConn.Close()
		client = nil
	}
	return nil
}

func (m *DefaultClientManager) GetOnlineCount() int {
	return len(m.Clients)
}

func (m *DefaultClientManager) PushBroadcast(msg any) (err error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for _, client := range m.Clients {
		_ = client.WsConn.WriteMessage(websocket.TextMessage, body)
	}
	return
}

func (m *DefaultClientManager) PushMessage(clientId string, msg any) (err error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if client, ok := m.Clients[clientId]; ok {
		_ = client.WsConn.WriteMessage(websocket.TextMessage, body)
	}
	return nil
}
