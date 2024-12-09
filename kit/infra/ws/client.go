package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// 用户类
type Client struct {
	ClientId  string
	UserId    string
	DeviceId  string
	IpAddress string
	IpSource  string

	ConnectTime int64
	ActiveTime  int64
	WsConn      *websocket.Conn
	Close       chan error

	//RegisterHandler   func(*Client) error
	//UnregisterHandler func(*Client) error
	//ReceiveHandler    func(*Client, *ReceiveMsg) error
}

//type Handler func(msg []byte) (err error)

func (c *Client) ReadMessage(handler func(msg []byte) (err error)) {
	closeText := "关闭连接"
	// 关闭连接（后端主动关闭）
	defer func() { // 避免忘记关闭，所以要加上close
		closeHandler := c.WsConn.CloseHandler()
		err := closeHandler(websocket.CloseInternalServerErr, closeText)
		if err != nil {
			return
		}
	}()

	for {

		//sendMsg := new(model.ReceiveMsg)
		//err := c.WsConn.ReadJSON(&sendMsg) // 读取json格式，如果不是json格式，会报错
		messageType, message, err := c.WsConn.ReadMessage()
		if err != nil {
			closeText = fmt.Sprintf("读取数据错误: %v", err)
			return
		}

		switch messageType {
		case websocket.TextMessage:
			err = handler(message)
			if err != nil {
				closeText = fmt.Sprintf("处理数据错误: %v", err)
				return
			}

		default:
			closeText = fmt.Sprintf("数据格式不正确:%v", messageType)
			return
		}
	}
}
