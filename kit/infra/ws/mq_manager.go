package ws

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mq"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
)

const (
	onlineKey = "chat:online"
)

type MqClientManager struct {
	Clients map[string]*Client
	// 用于存储在线用户
	Redis *redis.Redis

	p mq.MessagePublisher
	s mq.MessageSubscriber

	glog.Logger
}

func (m *MqClientManager) RegisterClient(client *Client) (err error) {
	// 将客户端 ID 加入 Redis Set
	m.Redis.ZaddCtx(context.Background(), onlineKey, time.Now().Unix(), client.ClientId)

	if _, ok := m.Clients[client.ClientId]; ok {
		return
	}

	m.Clients[client.ClientId] = client

	return
}

func (m *MqClientManager) UnRegisterClient(client *Client) (err error) {
	// 从 Redis Set 中移除客户端 ID
	m.Redis.ZremCtx(context.Background(), onlineKey, client.ClientId)

	if _, ok := m.Clients[client.ClientId]; ok {
		delete(m.Clients, client.ClientId)
		_ = client.WsConn.Close()
		client = nil
	}

	return
}

func (m *MqClientManager) GetOnlineCount() int {
	count, err := m.Redis.ZcardCtx(context.Background(), onlineKey)
	if err != nil {
		return 0
	}

	return count
}

func (m *MqClientManager) PushBroadcast(msg any) (err error) {
	err = m.p.PublishMessage(context.Background(), []byte(jsonconv.AnyToJsonNE(msg)))
	if err != nil {
		return err
	}
	return nil
}

func (m *MqClientManager) PushMessage(clientId string, msg any) (err error) {
	err = m.p.PublishMessage(context.Background(), []byte(jsonconv.AnyToJsonNE(msg)))
	if err != nil {
		return err
	}
	return nil
}

// 检测心跳任务
func (m *MqClientManager) OnlineTask() {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		// 定期收到信号
		case <-ticker.C:
			m.ClearInactiveClient()
		}
	}
}

func (m *MqClientManager) ClearInactiveClient() {

	redisClient := m.Redis
	expiredTime := time.Now().Add(-2 * time.Minute).Unix()

	count, _ := redisClient.ZcardCtx(context.Background(), onlineKey)
	if count > 0 {
		// 移除所有过期元素
		removedCount, err := redisClient.ZremrangebyscoreCtx(context.Background(), onlineKey, 0, expiredTime)
		if err != nil {
			m.Errorf("Failed to remove expired elements: %v", err)
		}
		if removedCount > 0 {
			m.Infof("Removed %d expired elements", removedCount)
		}
	}

	m.Infof("Removed %d expired elements", count)
}

var _ ClientManager = (*MqClientManager)(nil)
