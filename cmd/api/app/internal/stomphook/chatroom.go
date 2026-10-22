package stomphook

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-stomp/stomp/v3/frame"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/chatservice"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
	"github.com/ve-weiyi/stompws/server"
	"github.com/ve-weiyi/stompws/server/client"
	"github.com/ve-weiyi/vkit/adapter/iplocx"
	"github.com/ve-weiyi/vkit/x/jsonv"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/userservice"

	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/messaging/stompwsx"
)

// rpcTimeout 限定回调内远程调用的最长等待时间，避免连接协程被悬挂的 RPC 占住。
const rpcTimeout = 3 * time.Second

type ChatRoomEventHook struct {
	server.DefaultEventHook
	connectTime sync.Map // key: clientId, value: connectTime
	onlineUser  sync.Map // key: clientId, value: *userservice.User

	UserService userservice.UserService
	ChatService chatservice.ChatService
}

func NewChatRoomEventHook(userSvc userservice.UserService, chatSvc chatservice.ChatService) *ChatRoomEventHook {
	return &ChatRoomEventHook{
		UserService: userSvc,
		ChatService: chatSvc,
	}
}

func (h *ChatRoomEventHook) OnConnect(server *server.StompHubServer, c *client.Conn) {
	id, login, _ := c.GetSessionInfo()
	ip := c.GetIpAddress()

	h.connectTime.Store(id, time.Now())

	if login != "" {
		ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()

		out, err := h.UserService.GetUser(ctx, &userservice.GetUserRequest{
			UserId: login,
		})
		if err == nil && out.User != nil {
			h.onlineUser.Store(id, out.User)
		}
	}

	log.Printf("✅ User connected: (client: %s, user: %s, ip: %s) ", id, login, ip)
}

func (h *ChatRoomEventHook) OnDisconnect(server *server.StompHubServer, c *client.Conn) {
	id, login, _ := c.GetSessionInfo()
	ip := c.GetIpAddress()

	var duration time.Duration
	if value, exists := h.connectTime.LoadAndDelete(id); exists {
		connectTime := value.(time.Time)
		duration = time.Since(connectTime)
	}

	h.onlineUser.Delete(id)

	log.Printf("❌ User disconnected: (client: %s, user: %s, ip: %s), online: %v", id, login, ip, duration.Round(time.Second))
}

func (h *ChatRoomEventHook) OnSubscribe(server *server.StompHubServer, c *client.Conn, destination string, subscriptionId string) {
	_, login, _ := c.GetSessionInfo()
	ip := c.GetIpAddress()
	log.Printf("📢 User %s subscribed to %s", login, destination)
	count := stompwsx.OnlineUserCount(server)

	// 1. 私发问候消息
	greeting := frame.New(frame.MESSAGE, frame.Destination, destination, frame.MessageId, "0", frame.Subscription, subscriptionId)
	greeting.Body = []byte(jsonv.AnyToJsonNE(MessageEvent{
		Type: MessageTypeGreeting,
		Data: jsonv.AnyToJsonNE(GreetingMessageEvent{
			Content:   fmt.Sprintf("👋 welcome %s to the chat channel", login),
			IpAddress: ip,
			IpSource:  iplocx.GetIpSourceByBaidu(ip),
		}),
		TimeStamp: time.Now().UnixMilli(),
	}))
	c.SendFrame(greeting)

	// 2. 广播在线人数
	online := frame.New(frame.MESSAGE, frame.Destination, destination, frame.MessageId, "0", frame.Subscription, subscriptionId)
	online.Body = []byte(jsonv.AnyToJsonNE(MessageEvent{
		Type: MessageTypeOnline,
		Data: jsonv.AnyToJsonNE(OnlineMessageEvent{
			Online: true,
			Count:  count,
			Tips:   fmt.Sprintf("👋 %s joined the chat channel", login),
		}),
		TimeStamp: time.Now().UnixMilli(),
	}))
	server.RouteMessage(nil, online)

	// 3. 私发历史消息
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	out, err := h.ChatService.ListChats(ctx, &chatservice.ListChatsRequest{
		After:  time.Now().Add(-365 * 24 * time.Hour).UnixMilli(),
		Before: time.Now().UnixMilli(),
		Limit:  0,
	})
	if err != nil {
		return
	}
	list := make([]*ChatMessageEvent, 0)
	for _, msg := range out.List {
		list = append(list, &ChatMessageEvent{
			Id:        msg.Id,
			UserId:    msg.UserId,
			DeviceId:  msg.DeviceId,
			Nickname:  msg.Nickname,
			Avatar:    msg.Avatar,
			IpAddress: msg.IpAddress,
			IpSource:  msg.IpSource,
			Type:      msg.Type,
			Content:   msg.Content,
			Status:    msg.Status,
			CreatedAt: msg.CreatedAt,
			UpdatedAt: msg.UpdatedAt,
		})
	}

	history := frame.New(frame.MESSAGE, frame.Destination, destination, frame.MessageId, "0", frame.Subscription, subscriptionId)
	history.Body = []byte(jsonv.AnyToJsonNE(MessageEvent{
		Type: MessageTypeHistory,
		Data: jsonv.AnyToJsonNE(HistoryMessageEvent{
			List:  list,
			Page:  out.ListResult.Page,
			Size:  out.ListResult.PageSize,
			Total: out.ListResult.Total,
		}),
		TimeStamp: time.Now().UnixMilli(),
	}))
	c.SendFrame(history)
}

func (h *ChatRoomEventHook) OnUnsubscribe(server *server.StompHubServer, c *client.Conn, destination string, subscriptionId string) {
	_, login, _ := c.GetSessionInfo()
	log.Printf("📤 User %s unsubscribed from %s", login, destination)
	count := stompwsx.OnlineUserCount(server)

	online := frame.New(frame.MESSAGE, frame.Destination, destination, frame.MessageId, "0", frame.Subscription, subscriptionId)
	online.Body = []byte(jsonv.AnyToJsonNE(MessageEvent{
		Type: MessageTypeOnline,
		Data: jsonv.AnyToJsonNE(OnlineMessageEvent{
			Online: true,
			Count:  count,
			Tips:   fmt.Sprintf("👋 %s left the chat channel", login),
		}),
		TimeStamp: time.Now().UnixMilli(),
	}))
	server.RouteMessage(nil, online)
}

func (h *ChatRoomEventHook) OnSend(server *server.StompHubServer, c *client.Conn, message *frame.Frame) bool {
	clientId, login, _ := c.GetSessionInfo()
	ip := c.GetIpAddress()

	destination := message.Header.Get(frame.Destination)
	body := string(message.Body)

	log.Printf("💬 Message from %s to %s: %s", login, destination, body)
	if len(body) > 1000 {
		log.Printf("⚠️ Long message from %s (%d chars)", login, len(body))
	}

	var event MessageEvent
	jsonv.JsonToAny(body, &event)

	switch event.Type {
	case MessageTypeSend:
		var send SendMessageEvent
		jsonv.JsonToAny(event.Data, &send)

		ipSource := iplocx.GetIpSourceByBaidu(ip)
		var userId, nickname, avatar string
		if userInfo, ok := h.onlineUser.Load(clientId); ok {
			if user, ok := userInfo.(*userservice.User); ok {
				userId = user.UserId
				nickname = user.Nickname
				avatar = user.Avatar
			}
		}

		now := time.Now().UnixMilli()
		ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()

		out, err := h.ChatService.CreateChat(ctx, &chatservice.CreateChatRequest{
			UserId:    userId,
			DeviceId:  clientId,
			IpAddress: ip,
			IpSource:  ipSource,
			Nickname:  nickname,
			Avatar:    avatar,
			Type:      send.Type,
			Content:   send.Content,
			Status:    enums.ChatStatusNormal,
		})
		if err != nil {
			return false
		}

		msgFrame := frame.New(frame.MESSAGE, frame.Destination, destination, frame.MessageId, "1")
		msgFrame.Body = []byte(jsonv.AnyToJsonNE(MessageEvent{
			Type: MessageTypeMessage,
			Data: jsonv.AnyToJsonNE(ChatMessageEvent{
				Id:        out.Id,
				UserId:    userId,
				DeviceId:  clientId,
				Nickname:  nickname,
				Avatar:    avatar,
				IpAddress: ip,
				IpSource:  ipSource,
				Type:      send.Type,
				Content:   send.Content,
				Status:    enums.ChatStatusNormal,
				CreatedAt: now,
				UpdatedAt: now,
			}),
			TimeStamp: now,
		}))
		server.RouteMessage(c, msgFrame)

	case MessageTypeEdit:
		var edit EditMessageEvent
		jsonv.JsonToAny(event.Data, &edit)

		var userId string
		if userInfo, ok := h.onlineUser.Load(clientId); ok {
			if user, ok := userInfo.(*userservice.User); ok {
				userId = user.UserId
			}
		}

		now := time.Now().UnixMilli()
		rpcCtx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
		defer cancel()

		ctx := metadata.AppendToOutgoingContext(rpcCtx,
			headerx.HeaderUid, userId,
			headerx.HeaderXDeviceId, clientId,
		)
		_, err := h.ChatService.UpdateChat(ctx, &chatservice.UpdateChatRequest{
			Id:      edit.Id,
			Type:    edit.Type,
			Content: edit.Content,
			Status:  edit.Status,
		})
		if err != nil {
			return false
		}

		msgFrame := frame.New(frame.MESSAGE, frame.Destination, destination, frame.MessageId, "1")
		msgFrame.Body = []byte(jsonv.AnyToJsonNE(MessageEvent{
			Type: MessageTypeEdit,
			Data: jsonv.AnyToJsonNE(EditMessageEvent{
				Id:        edit.Id,
				Type:      edit.Type,
				Content:   edit.Content,
				Status:    edit.Status,
				UpdatedAt: now,
			}),
			TimeStamp: now,
		}))
		server.RouteMessage(c, msgFrame)

	default:
		return false
	}

	return true
}

func (h *ChatRoomEventHook) OnAck(c *client.Conn, messageId string) {
	_, login, _ := c.GetSessionInfo()
	log.Printf("✅ Message %s acknowledged by %s", messageId, login)
}

func (h *ChatRoomEventHook) OnNack(c *client.Conn, messageId string) {
	_, login, _ := c.GetSessionInfo()
	log.Printf("❌ Message %s rejected by %s", messageId, login)
}

func (h *ChatRoomEventHook) OnError(c *client.Conn, err error) {
	_, login, _ := c.GetSessionInfo()
	log.Printf("🚨 Error for user %s: %v", login, err)
}
