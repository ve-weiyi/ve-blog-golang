// Package stompwsx 装本仓对 stompws 的封装：与在线状态有关的客户端推送。
//
// 它被 app 与 admin 两个网关共用 —— 两边的这段逻辑本就一字不差，各留一份只会各自漂移。
// 在线状态本身不在这里：存储与成员格式见 infra/storex/presencex，装配见各服务的 svc。
//
// 在线状态从回调收到的 hub 上取（`hub.Presence()`），不作为构造参数传进来：
// hub 本就持有它，重复注入只多出一种「hub 用一个、钩子用另一个」的错配，
// 而且回调的第一参就是 hub —— 它正是这个体系里的上下文。
package stompwsx

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/go-stomp/stomp/v3/frame"

	"github.com/ve-weiyi/stompws/server"
	"github.com/ve-weiyi/stompws/server/client"
)

// presenceTimeout 限定在线人数查询的最长等待时间；查询失败按 0 上报。
const presenceTimeout = time.Second

// OnlineUserCount 取当前在线人数（唯一用户数）。查不到就按 0 报并记一条 ——
// 在线人数只是提示信息，不该因为它查询失败就打断问候或退订的应答。
func OnlineUserCount(hub *server.StompHubServer) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), presenceTimeout)
	defer cancel()

	stats, err := hub.Presence().PresenceStats(ctx)
	if err != nil {
		log.Printf("📢 查询在线人数失败: %v", err)
		return 0
	}
	return int64(stats.NumUsers)
}

// OnlineBroadcastHook 在有人上线/下线时把在线人数播给 /topic/system/online。
// 只在人数确实变了才播，免得每次连接都刷一遍。
type OnlineBroadcastHook struct {
	server.DefaultEventHook

	mu        sync.Mutex
	lastCount int
}

func NewOnlineBroadcastHook() *OnlineBroadcastHook {
	return &OnlineBroadcastHook{}
}

func (h *OnlineBroadcastHook) OnConnect(hub *server.StompHubServer, _ *client.Conn) {
	h.broadcast(hub)
}

func (h *OnlineBroadcastHook) OnDisconnect(hub *server.StompHubServer, _ *client.Conn) {
	h.broadcast(hub)
}

func (h *OnlineBroadcastHook) broadcast(hub *server.StompHubServer) {
	ctx, cancel := context.WithTimeout(context.Background(), presenceTimeout)
	defer cancel()

	stats, err := hub.Presence().PresenceStats(ctx)
	if err != nil {
		log.Printf("📢 播报在线人数失败: %v", err)
		return
	}

	h.mu.Lock()
	if stats.NumUsers == h.lastCount {
		h.mu.Unlock()
		return
	}
	h.lastCount = stats.NumUsers
	h.mu.Unlock()

	msg := frame.New(frame.MESSAGE, frame.Destination, "/topic/system/online", frame.MessageId, "0")
	msg.Body = []byte(fmt.Sprintf(`{"type":"online_count","count":%d}`, stats.NumUsers))
	hub.RouteMessage(nil, msg)
}

// OnlineCatchupHook 在新订阅者订阅在线主题时，把当前人数补发给它 ——
// 广播只播「变化」，刚订阅的人得先拿到当前值。
type OnlineCatchupHook struct {
	server.DefaultEventHook
}

func NewOnlineCatchupHook() *OnlineCatchupHook {
	return &OnlineCatchupHook{}
}

func (h *OnlineCatchupHook) OnSubscribe(hub *server.StompHubServer, c *client.Conn, destination string, subscriptionId string) {
	if !strings.HasPrefix(destination, "/topic/system/online") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), presenceTimeout)
	defer cancel()

	stats, err := hub.Presence().PresenceStats(ctx)
	if err != nil {
		log.Printf("📢 补发在线人数失败: %v", err)
		return
	}

	msg := frame.New(frame.MESSAGE, frame.Destination, "/topic/system/online", frame.MessageId, "0", frame.Subscription, subscriptionId)
	msg.Body = []byte(fmt.Sprintf(`{"type":"online_count","count":%d}`, stats.NumUsers))
	c.SendFrame(msg)
}
