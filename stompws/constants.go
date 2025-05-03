package stompws

import (
	"errors"
	"time"
)

// 常量定义
const (
	StompCommandConnected   = "CONNECTED"
	StompCommandMessage     = "MESSAGE"
	StompCommandReceipt     = "RECEIPT"
	StompCommandError       = "ERROR"
	WildcardSingleLevel     = "*"         // 单级通配符
	WildcardMultiLevel      = "**"        // 多级通配符
	Separator               = "/"         // 主题层级分隔符
	DefaultMessageQueueSize = 256         // 默认消息队列大小
	MaxMessageSize          = 1024 * 1024 // 1MB 最大消息大小

	WriteWait  = 10 * time.Second
	PongWait   = 60 * time.Second
	PingPeriod = (PongWait * 9) / 10
)

var (
	ErrTopicNotAllowed    = errors.New("topic not allowed")
	ErrSubscriptionDenied = errors.New("subscription denied")
	ErrPublishDenied      = errors.New("publish denied")
	ErrHandlerNotDefined  = errors.New("handler not defined")
	ErrInvalidFrame       = errors.New("invalid STOMP frame")
	ErrMessageTooLarge    = errors.New("message too large")
	ErrClientDisconnected = errors.New("client disconnected")
	ErrSendTimeout        = errors.New("send timeout")
	ErrServerShuttingDown = errors.New("server shutting down")
)
