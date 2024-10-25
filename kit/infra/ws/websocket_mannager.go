package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	// 用于升级 HTTP 连接
	websocket.Upgrader
	// 用于存储连接
	connections map[string]*websocket.Conn
}

// NewWebSocketManager 创建一个 WebSocket 管理器
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// 允许所有请求来源
				return true
			},
		},
		connections: make(map[string]*websocket.Conn),
	}
}

// 开始处理 WebSocket 连接
func (m *WebSocketManager) RegisterWebSocket(w http.ResponseWriter, r *http.Request, name string) *websocket.Conn {
	conn, err := m.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return nil
	}

	// 将新的连接加入客户端列表
	m.connections[name] = conn

	return conn
}

// 处理接收到的消息
func (m *WebSocketManager) OnReceiveMsg(name string, receive Receive) {
	conn, ok := m.connections[name]
	if !ok {
		return
	}
	for {
		//读取客户端发送的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			delete(m.connections, conn.RemoteAddr().String())
			break
		}

		//处理消息
		tx, err := receive(msg)
		if err != nil {
			log.Println("Failed to handle message:", err)
			_ = conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
			continue
		}

		if tx == nil {
			continue
		}

		//将消息广播给所有连接的客户端
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, tx)
			if err != nil {
				log.Println("Failed to send message:", err)
				client.Close()
			}
		}
	}

}

// 给指定的客户端发送消息
func (m *WebSocketManager) SendMsgToClient(name string, msg []byte) {
	conn, ok := m.connections[name]
	if !ok {
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("Failed to send message:", err)
		conn.Close()
		delete(m.connections, name)
	}
}

// 广播消息
func (m *WebSocketManager) BroadcastMsg(msg []byte) {
	for _, conn := range m.connections {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		log.Println("Failed to send message:", err)
		conn.Close()
		delete(clients, conn)
	}
}
