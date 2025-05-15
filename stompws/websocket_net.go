package stompws

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketConn struct {
	Conn   *websocket.Conn
	reader io.Reader

	readMu  sync.Mutex // 保护读取操作和reader状态
	writeMu sync.Mutex // 保护写入操作
}

func NewWebSocketConn(conn *websocket.Conn) net.Conn {
	return &WebSocketConn{Conn: conn}
}

func (c *WebSocketConn) Read(b []byte) (int, error) {
	c.readMu.Lock()
	defer c.readMu.Unlock()

	// 直接读取下一条消息
	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			return 0, io.EOF
		}
		return 0, err
	}

	// 将整条消息拷贝到提供的缓冲区
	n := copy(b, message)

	// 如果提供的缓冲区太小，丢弃剩余数据
	if n < len(message) {
		return n, io.ErrShortBuffer
	}

	// stomp心跳帧，不需要打印
	if string(message) != "\n" {
		fmt.Println("<---------start----------")
		fmt.Println(string(message))
		fmt.Println("<----------end-----------")
	}
	return n, nil
}

func (c *WebSocketConn) Write(b []byte) (int, error) {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()

	err := c.Conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return 0, err
	}

	// stomp心跳帧，不需要打印
	if string(b) != "\n" {
		fmt.Println("----------start--------->")
		fmt.Println(string(b))
		fmt.Println("-----------end---------->")
	}
	return len(b), nil
}

func (c *WebSocketConn) Close() error {
	// 直接委托给底层连接，gorilla/websocket会处理关闭握手
	return c.Conn.Close()
}

func (c *WebSocketConn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

func (c *WebSocketConn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *WebSocketConn) SetDeadline(t time.Time) error {
	if err := c.SetReadDeadline(t); err != nil {
		return err
	}
	return c.SetWriteDeadline(t)
}

func (c *WebSocketConn) SetReadDeadline(t time.Time) error {
	return c.Conn.SetReadDeadline(t)
}

func (c *WebSocketConn) SetWriteDeadline(t time.Time) error {
	return c.Conn.SetWriteDeadline(t)
}
