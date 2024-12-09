package ws

import (
	"log"
	"time"

	"github.com/gorilla/websocket"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
)

const (
	// MaxMessageSize 消息大小
	MaxMessageSize = 8192
	// PingPeriod 每次ping的间隔时长
	PingPeriod = 30 * time.Second
	// PongPeriod 每次pong的间隔时长，可以是PingPeriod的一倍|两倍
	PongPeriod = 54 * time.Second
	// WriteWait client的写入等待时长
	WriteWait = 5 * time.Second
	// ReadWait client的读取等待时长
	ReadWait = 60 * time.Second
)

// 用户管理
type SocketServer struct {
	Register   chan *Client // 注册消息
	Unregister chan *Client // 注销消息

	RegisterHandler   func(*Client) error
	UnregisterHandler func(*Client) error
	ReceiveHandler    func(*Client, []byte) error
}

func NewSocketServer() *SocketServer {
	s := SocketServer{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}

	go s.Start()
	return &s
}

func (s *SocketServer) Start() {
	log.Println("<---监听管道通信--->")
	for {
		select {

		// 建立连接
		case conn := <-s.Register: // 建立连接
			glog.Infof("建立新连接: %v", conn.ClientId)
			err := s.RegisterHandler(conn)
			if err != nil {
				return
			}
		// 断开连接
		case conn := <-s.Unregister: // 断开连接
			glog.Infof("关闭连接:%v", conn.ClientId)
			err := s.UnregisterHandler(conn)
			if err != nil {
				return
			}
		}
	}
}

func (s *SocketServer) ReadChannel(c *Client) {
	defer func() { // 避免忘记关闭，所以要加上close
		glog.Info("关闭连接")
		s.Unregister <- c
	}()

	for {

		//sendMsg := new(model.ReceiveMsg)
		//err := c.WsConn.ReadJSON(&sendMsg) // 读取json格式，如果不是json格式，会报错
		messageType, message, err := c.WsConn.ReadMessage()
		if err != nil {
			glog.Info("读取数据错误", err)
			return
		}

		switch messageType {
		case websocket.TextMessage:
			glog.Info("接收到消息：", string(message))
			err := s.ReceiveHandler(c, message)
			if err != nil {
				return
			}

		default:
			glog.Info("数据格式不正确", messageType, message, err)
		}
	}
}
