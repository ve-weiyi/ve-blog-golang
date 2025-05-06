package stompws

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// Broker 管理所有客户端连接和消息广播
type Broker struct {
	clients      sync.Map
	broadcast    chan *BroadcastMessage
	register     chan *Client
	unregister   chan *Client
	topics       sync.Map
	patterns     sync.Map
	topicConfigs sync.Map
	shutdown     chan struct{}
	logger       logx.Logger
	stats        *StatsCollector
}

func NewBroker(logger logx.Logger) *Broker {
	if logger == nil {
		logger = logx.WithContext(context.Background())
	}
	return &Broker{
		broadcast:  make(chan *BroadcastMessage, 1024),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		shutdown:   make(chan struct{}),
		logger:     logger,
		stats:      NewStatsCollector(),
	}
}

func (b *Broker) DefineTopic(topic string, handler *TopicHandler) error {
	if handler == nil {
		return errors.New("handler cannot be nil")
	}
	b.topicConfigs.Store(topic, handler)
	b.logger.Info("Topic defined", logx.Field("topic", topic))
	return nil
}

func (b *Broker) Run(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			b.logger.Info("Broker shutting down")
			close(b.shutdown)
			return
		case client := <-b.register:
			b.handleClientRegister(client)
		case client := <-b.unregister:
			b.handleClientUnregister(client)
		case bMsg := <-b.broadcast:
			b.handleBroadcast(bMsg)
		}
	}
}

func (b *Broker) handleClientRegister(client *Client) {
	b.clients.Store(client, true)
	b.stats.IncrementConnections()

	b.logger.Info("Client connected",
		logx.Field("clientID", client.ClientID),
		logx.Field("sessionID", client.sessionID))
}

func (b *Broker) handleClientUnregister(client *Client) {
	if _, loaded := b.clients.LoadAndDelete(client); loaded {
		client.Disconnect()
		b.stats.DecrementConnections()

		b.logger.Info("Client disconnected",
			logx.Field("clientID", client.ClientID),
			logx.Field("sessionID", client.sessionID))
	}
}

func (b *Broker) handleBroadcast(bMsg *BroadcastMessage) {
	handler, err := b.getTopicHandler(bMsg.Topic)
	if err != nil {
		b.stats.IncrementMessagesDropped()
		b.logger.Errorw("No handler for topic, message dropped",
			logx.Field("topic", bMsg.Topic))

		if defaultHandler, err := b.getDefaultHandler(); err == nil && defaultHandler.OnMessageDrop != nil {
			if err := defaultHandler.OnMessageDrop(bMsg.Context, bMsg.Topic, bMsg.Message, bMsg.Sender); err != nil {
				b.logger.Error("Error in default OnMessageDrop handler",
					logx.Field("topic", bMsg.Topic),
					logx.Field("err", err))
			}
		}
		return
	}

	if handler.OnMessagePublish != nil {
		if err := handler.OnMessagePublish(bMsg.Context, bMsg.Topic, bMsg.Message, bMsg.Sender); err != nil {
			b.stats.IncrementErrors()
			b.logger.Error("Error in OnMessagePublish handler",
				logx.Field("topic", bMsg.Topic),
				logx.Field("err", err))

			if errors.Is(err, ErrPublishDenied) && handler.OnMessageDrop != nil {
				if err := handler.OnMessageDrop(bMsg.Context, bMsg.Topic, bMsg.Message, bMsg.Sender); err != nil {
					b.logger.Error("Error in OnMessageDrop handler",
						logx.Field("topic", bMsg.Topic),
						logx.Field("err", err))
				}
			}
			return
		}
	}

	stompFrame, messageID := b.buildStompMessage(bMsg.Topic, bMsg.Message)
	matchedClients := b.findMatchingClients(bMsg.Topic)

	if len(matchedClients) == 0 {
		b.stats.IncrementMessagesDropped()
		b.logger.Errorw("No subscribers for topic, message dropped",
			logx.Field("topic", bMsg.Topic))

		if handler.OnMessageDrop != nil {
			if err := handler.OnMessageDrop(bMsg.Context, bMsg.Topic, bMsg.Message, bMsg.Sender); err != nil {
				b.logger.Error("Error in OnMessageDrop handler",
					logx.Field("topic", bMsg.Topic),
					logx.Field("err", err))
			}
		}
		return
	}

	b.deliverMessage(bMsg, stompFrame, messageID, handler, matchedClients)
}

func (b *Broker) findMatchingClients(topic string) []*Client {
	var matchedClients []*Client

	// Check exact topic matches
	if clients, ok := b.topics.Load(topic); ok {
		for client := range clients.(map[*Client]bool) {
			matchedClients = append(matchedClients, client)
		}
	}

	// Check pattern matches
	b.patterns.Range(func(key, value interface{}) bool {
		client := key.(*Client)
		patterns := value.([]string)

		for _, pattern := range patterns {
			if topicMatchesPattern(topic, pattern) {
				matchedClients = append(matchedClients, client)
				break
			}
		}
		return true
	})

	return matchedClients
}

func (b *Broker) getTopicHandler(topic string) (*TopicHandler, error) {
	if handler, ok := b.topicConfigs.Load(topic); ok {
		return handler.(*TopicHandler), nil
	}

	var matchedHandler *TopicHandler
	b.topicConfigs.Range(func(key, value interface{}) bool {
		pattern := key.(string)
		if topicMatchesPattern(topic, pattern) {
			matchedHandler = value.(*TopicHandler)
			return false
		}
		return true
	})

	if matchedHandler != nil {
		return matchedHandler, nil
	}

	return nil, ErrHandlerNotDefined
}

func (b *Broker) getDefaultHandler() (*TopicHandler, error) {
	if handler, ok := b.topicConfigs.Load("**"); ok {
		return handler.(*TopicHandler), nil
	}
	return nil, ErrHandlerNotDefined
}

func (b *Broker) isTopicAllowed(pattern string) bool {
	if _, ok := b.topicConfigs.Load("**"); ok {
		return true
	}

	found := false
	b.topicConfigs.Range(func(key, _ interface{}) bool {
		topic := key.(string)
		if topic == pattern || topicMatchesPattern(pattern, topic) {
			found = true
			return false
		}
		return true
	})

	return found
}

func (b *Broker) buildStompMessage(destination, body string) (string, string) {
	messageID := fmt.Sprintf("msg-%d", time.Now().UnixNano())
	frame := fmt.Sprintf("MESSAGE\ndestination:%s\nmessage-id:%s\ncontent-length:%d\n\n%s\x00",
		destination, messageID, len(body), body)
	return frame, messageID
}

func (b *Broker) deliverMessage(bMsg *BroadcastMessage, frame, messageID string, handler *TopicHandler, clients []*Client) {
	for _, client := range clients {
		if err := client.Send([]byte(frame)); err != nil {
			b.stats.IncrementErrors()
			b.logger.Error("Failed to deliver message",
				logx.Field("clientID", client.ClientID),
				logx.Field("topic", bMsg.Topic),
				logx.Field("err", err))
			continue
		}

		b.stats.IncrementMessagesSent()
		if handler.OnMessageDeliver != nil {
			if err := handler.OnMessageDeliver(bMsg.Context, bMsg.Topic, bMsg.Message, client); err != nil {
				b.stats.IncrementErrors()
				b.logger.Error("Error in OnMessageDeliver handler",
					logx.Field("clientID", client.ClientID),
					logx.Field("topic", bMsg.Topic),
					logx.Field("err", err))
			}
		}

		b.logger.Debug("Message delivered",
			logx.Field("clientID", client.ClientID),
			logx.Field("sessionID", client.sessionID),
			logx.Field("topic", bMsg.Topic),
			logx.Field("messageID", messageID))
	}

	b.logger.Info("Message published",
		logx.Field("topic", bMsg.Topic),
		logx.Field("messageID", messageID))
}

func (b *Broker) subscribe(ctx context.Context, client *Client, pattern string) error {
	if _, ok := client.topics[pattern]; ok {
		return nil
	}

	if !b.isTopicAllowed(pattern) {
		b.logger.Errorw("Subscription denied for undefined topic",
			logx.Field("clientID", client.ClientID),
			logx.Field("pattern", pattern))
		return ErrTopicNotAllowed
	}

	handler, err := b.getTopicHandler(pattern)
	if err != nil {
		return err
	}

	if handler.OnMessageSubscribe != nil {
		if err := handler.OnMessageSubscribe(ctx, pattern, client); err != nil {
			b.logger.Errorw("Subscription denied by handler",
				logx.Field("clientID", client.ClientID),
				logx.Field("pattern", pattern),
				logx.Field("err", err))
			return ErrSubscriptionDenied
		}
	}

	client.topics[pattern] = true

	var patterns []string
	if p, ok := b.patterns.Load(client); ok {
		patterns = p.([]string)
	}
	patterns = append(patterns, pattern)
	b.patterns.Store(client, patterns)

	b.stats.IncrementSubscriptions()

	if !strings.Contains(pattern, WildcardSingleLevel) && !strings.Contains(pattern, WildcardMultiLevel) {
		var clients map[*Client]bool
		if c, ok := b.topics.Load(pattern); ok {
			clients = c.(map[*Client]bool)
		} else {
			clients = make(map[*Client]bool)
		}
		clients[client] = true
		b.topics.Store(pattern, clients)
	}

	b.logger.Info("Client subscribed",
		logx.Field("clientID", client.ClientID),
		logx.Field("sessionID", client.sessionID),
		logx.Field("pattern", pattern))

	return nil
}

func (b *Broker) unsubscribe(ctx context.Context, client *Client, pattern string) error {
	if _, ok := client.topics[pattern]; !ok {
		return nil
	}

	handler, err := b.getTopicHandler(pattern)
	if err == nil && handler.OnMessageUnsubscribe != nil {
		if err := handler.OnMessageUnsubscribe(ctx, pattern, client); err != nil {
			b.logger.Error("Error in OnMessageUnsubscribe handler",
				logx.Field("clientID", client.ClientID),
				logx.Field("pattern", pattern),
				logx.Field("err", err))
			return err
		}
	}

	delete(client.topics, pattern)

	if patterns, ok := b.patterns.Load(client); ok {
		updatedPatterns := make([]string, 0, len(patterns.([]string))-1)
		for _, p := range patterns.([]string) {
			if p != pattern {
				updatedPatterns = append(updatedPatterns, p)
			}
		}
		if len(updatedPatterns) == 0 {
			b.patterns.Delete(client)
		} else {
			b.patterns.Store(client, updatedPatterns)
		}
	}

	if clients, ok := b.topics.Load(pattern); ok {
		delete(clients.(map[*Client]bool), client)
		if len(clients.(map[*Client]bool)) == 0 {
			b.topics.Delete(pattern)
		}
	}

	b.stats.DecrementSubscriptions()

	b.logger.Info("Client unsubscribed",
		logx.Field("clientID", client.ClientID),
		logx.Field("sessionID", client.sessionID),
		logx.Field("pattern", pattern))

	return nil
}

func (b *Broker) handleConnect(ctx context.Context, client *Client, frame *StompFrame) {
	// 客户端建议10秒发一次，期望5秒收一次
	respFrame := StompFrame{
		Command: "CONNECTED",
		Headers: map[string]string{
			"version":    "1.2",
			"session":    client.sessionID,
			"heart-beat": frame.Headers["heart-beat"], // 服务器心跳设置：发送间隔,期望接收间隔
		},
		Body: "",
	}

	if err := client.Send([]byte(respFrame.Serialize())); err != nil {
		b.logger.Error("Failed to send CONNECTED frame",
			logx.Field("clientID", client.ClientID),
			logx.Field("err", err))
	}
}

func (b *Broker) handlePing(ctx context.Context, client *Client, frame *StompFrame) {
	// 客户端建议10秒发一次，期望5秒收一次
	if err := client.Send([]byte("\n")); err != nil {
		b.logger.Error("Failed to send PING frame",
			logx.Field("clientID", client.ClientID),
			logx.Field("err", err))
	}
}

func (b *Broker) handleSubscribe(ctx context.Context, client *Client, frame *StompFrame) {
	destination, ok := frame.Headers["destination"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing destination header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	if err := b.subscribe(ctx, client, destination); err != nil {
		b.stats.IncrementErrors()
		errorFrame := fmt.Sprintf("ERROR\nmessage:%s\n\n\x00", err.Error())
		client.Send([]byte(errorFrame))
		return
	}

	if receipt, ok := frame.Headers["receipt"]; ok {
		receiptFrame := "RECEIPT\nreceipt-id:" + receipt + "\n\n\x00"
		client.Send([]byte(receiptFrame))
	}
}

func (b *Broker) handleUnsubscribe(ctx context.Context, client *Client, frame *StompFrame) {
	destination, ok := frame.Headers["destination"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing destination header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	if err := b.unsubscribe(ctx, client, destination); err != nil {
		b.stats.IncrementErrors()
		errorFrame := fmt.Sprintf("ERROR\nmessage:%s\n\n\x00", err.Error())
		client.Send([]byte(errorFrame))
		return
	}

	if receipt, ok := frame.Headers["receipt"]; ok {
		receiptFrame := "RECEIPT\nreceipt-id:" + receipt + "\n\n\x00"
		client.Send([]byte(receiptFrame))
	}
}

func (b *Broker) handleSend(ctx context.Context, client *Client, frame *StompFrame) {
	destination, ok := frame.Headers["destination"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing destination header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	bMsg := &BroadcastMessage{
		Topic:   destination,
		Message: frame.Body,
		Sender:  client,
		Context: ctx,
	}

	select {
	case b.broadcast <- bMsg:
	case <-time.After(WriteWait):
		b.stats.IncrementErrors()
		b.logger.Error("Broadcast channel full",
			logx.Field("topic", destination))
	case <-b.shutdown:
	}
}

func (b *Broker) handleAck(ctx context.Context, client *Client, frame *StompFrame) {
	messageID, ok := frame.Headers["message-id"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing message-id header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	destination, ok := frame.Headers["destination"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing destination header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	handler, err := b.getTopicHandler(destination)
	if err != nil {
		b.logger.Errorw("No handler for topic",
			logx.Field("topic", destination))
		return
	}

	if handler.OnMessageAck != nil {
		if err := handler.OnMessageAck(ctx, destination, messageID, client); err != nil {
			b.stats.IncrementErrors()
			b.logger.Error("Error in OnMessageAck handler",
				logx.Field("clientID", client.ClientID),
				logx.Field("messageID", messageID),
				logx.Field("err", err))
		}
	}

	b.logger.Debug("Message acknowledged",
		logx.Field("clientID", client.ClientID),
		logx.Field("messageID", messageID),
		logx.Field("topic", destination))
}

func (b *Broker) handleNack(ctx context.Context, client *Client, frame *StompFrame) {
	messageID, ok := frame.Headers["message-id"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing message-id header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	destination, ok := frame.Headers["destination"]
	if !ok {
		b.stats.IncrementErrors()
		errorFrame := "ERROR\nmessage:Missing destination header\n\n\x00"
		client.Send([]byte(errorFrame))
		return
	}

	handler, err := b.getTopicHandler(destination)
	if err != nil {
		b.logger.Errorw("No handler for topic",
			logx.Field("topic", destination))
		return
	}

	if handler.OnMessageNack != nil {
		if err := handler.OnMessageNack(ctx, destination, messageID, client); err != nil {
			b.stats.IncrementErrors()
			b.logger.Error("Error in OnMessageNack handler",
				logx.Field("clientID", client.ClientID),
				logx.Field("messageID", messageID),
				logx.Field("err", err))
		}
	}

	b.logger.Debug("Message not acknowledged",
		logx.Field("clientID", client.ClientID),
		logx.Field("messageID", messageID),
		logx.Field("topic", destination))
}

func (b *Broker) handleDisconnect(ctx context.Context, client *Client, frame *StompFrame) {
	if receipt, ok := frame.Headers["receipt"]; ok {
		receiptFrame := "RECEIPT\nreceipt-id:" + receipt + "\n\n\x00"
		client.Send([]byte(receiptFrame))
	}
}

// topicMatchesPattern 检查主题是否匹配订阅模式
func topicMatchesPattern(topic, pattern string) bool {
	// 如果是精确匹配
	if pattern == topic {
		return true
	}

	// 分割主题和模式
	topicParts := strings.Split(topic, Separator)
	patternParts := strings.Split(pattern, Separator)

	// 处理多级通配符 **
	for i := 0; i < len(patternParts) && i < len(topicParts); i++ {
		if patternParts[i] == WildcardMultiLevel {
			return true
		}
		if patternParts[i] == WildcardSingleLevel {
			continue
		}
		if patternParts[i] != topicParts[i] {
			return false
		}
	}

	return len(patternParts) == len(topicParts)
}
