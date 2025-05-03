package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/stompws"
)

func main() {
	logger := logx.WithContext(context.Background())
	broker := stompws.NewBroker(logger)

	// 定义主题处理函数
	broker.DefineTopic("/topic/chat", &stompws.TopicHandler{
		OnMessageSubscribe: func(ctx context.Context, topic string, client *stompws.Client) error {
			logger.Info("New subscription",
				logx.Field("topic", topic),
				logx.Field("clientID", client.ClientID))
			return nil
		},
		OnMessageUnsubscribe: func(ctx context.Context, topic string, client *stompws.Client) error {
			logger.Info("Subscription ended",
				logx.Field("topic", topic),
				logx.Field("clientID", client.ClientID))
			return nil
		},
		OnMessagePublish: func(ctx context.Context, topic, message string, client *stompws.Client) error {
			if strings.Contains(message, "forbidden") {
				return fmt.Errorf("message contains forbidden content")
			}
			logger.Debug("Message published",
				logx.Field("topic", topic),
				logx.Field("clientID", client.ClientID))
			return nil
		},
		OnMessageDeliver: func(ctx context.Context, topic, message string, client *stompws.Client) error {
			logger.Debug("Message delivered",
				logx.Field("topic", topic),
				logx.Field("clientID", client.ClientID))
			return nil
		},
		OnMessageDrop: func(ctx context.Context, topic, message string, client *stompws.Client) error {
			logger.Errorw("Message dropped",
				logx.Field("topic", topic),
				logx.Field("clientID", client.ClientID))
			return nil
		},
		OnMessageAck: func(ctx context.Context, topic, messageID string, client *stompws.Client) error {
			logger.Debug("Message acknowledged",
				logx.Field("topic", topic),
				logx.Field("messageID", messageID),
				logx.Field("clientID", client.ClientID))
			return nil
		},
		OnMessageNack: func(ctx context.Context, topic, messageID string, client *stompws.Client) error {
			logger.Errorw("Message not acknowledged",
				logx.Field("topic", topic),
				logx.Field("messageID", messageID),
				logx.Field("clientID", client.ClientID))
			return nil
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动broker
	go broker.Run(ctx)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		stompws.HandleWebSocket(broker, w, r)
	})

	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Starting STOMP server on :8080")
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server failed", err)
	}

	<-ctx.Done()
	logger.Info("Server shutdown complete")
}
