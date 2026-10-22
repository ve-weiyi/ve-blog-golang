// Package tokenx 提供令牌签发、校验、刷新与撤销能力。
//
// 令牌格式由 Codec 决定，会话状态由 SessionStore 保存；Service 只编排两者，
// 不依赖具体签名算法或存储实现。
package tokenx

import (
	"context"
	"errors"
	"time"
)

// Service 是令牌子系统的对外接口。
type Service interface {
	Issue(ctx context.Context, req IssueRequest) (*TokenPair, error)
	Validate(ctx context.Context, accessToken string) (*Claims, error)
	Refresh(ctx context.Context, req RefreshRequest) (*TokenPair, error)

	Revoke(ctx context.Context, userID, sessionID string) error
	RevokeDevice(ctx context.Context, userID, deviceID string) error
	RevokeAll(ctx context.Context, userID string) error

	ListSessions(ctx context.Context, userID string) ([]SessionInfo, error)
}

// IssueRequest 是一次令牌签发请求。
type IssueRequest struct {
	UserID   string
	DeviceID string
	Issuer   string
	ClientID string
}

// RefreshRequest 是一次刷新请求。
// UserID/DeviceID 必须来自当前请求，并与 token/session 中的身份一致。
type RefreshRequest struct {
	RefreshToken string
	UserID       string
	DeviceID     string
}

// TokenPair 是签发给客户端的令牌对。
type TokenPair struct {
	UserID   string
	DeviceID string

	AccessToken  string
	RefreshToken string
	SessionID    string

	AccessExpiresAt  time.Time
	AccessExpiresIn  int64
	RefreshExpiresAt time.Time
	RefreshExpiresIn int64
}

// SessionInfo 是面向调用方的脱敏会话视图。
type SessionInfo struct {
	ID                string
	UserID            string
	DeviceID          string
	ClientID          string
	CreatedAt         time.Time
	LastUsedAt        time.Time
	AccessExpiresAt   time.Time
	RefreshExpiresAt  time.Time
	AbsoluteExpiresAt time.Time
}

// Claims 是令牌携带的身份与生命周期信息。
type Claims struct {
	UserID    string
	DeviceID  string
	SessionID string
	Issuer    string
	TokenID   string // JWT jti；opaque 为空
	Kind      TokenKind
	IssuedAt  time.Time
	ExpiresAt time.Time
}

// TokenKind 区分 access 与 refresh 令牌。
type TokenKind string

const (
	KindAccess  TokenKind = "access"
	KindRefresh TokenKind = "refresh"
)

// LoginMode 决定新登录如何处理同一用户的既有会话。
type LoginMode int

const (
	MultiPoint LoginMode = iota
	SinglePoint
)

// DefaultDeviceID 是空 deviceID 归一化后的值。
const DefaultDeviceID = "default"

// ClockLeeway 是令牌与会话过期判断共用的时钟偏差容忍度。
const ClockLeeway = 30 * time.Second

// GraceWindow 是 refresh 轮换结果的宽限窗口，用于吸收并发刷新与网络重试。
const GraceWindow = 30 * time.Second

// NormalizeDeviceID 把空设备标识归一化为 DefaultDeviceID。
func NormalizeDeviceID(deviceID string) string {
	if deviceID == "" {
		return DefaultDeviceID
	}
	return deviceID
}

// Clock 是令牌服务使用的时间源，测试可注入假时钟。
type Clock interface {
	Now() time.Time
}

// SystemClock 使用系统时间。
type SystemClock struct{}

func (SystemClock) Now() time.Time { return time.Now() }

// Config 是 Service 的构造配置。
type Config struct {
	AccessCodec  Codec
	RefreshCodec Codec // 为空时复用 AccessCodec
	Store        SessionStore
	LoginMode    LoginMode

	AccessTTL          time.Duration
	RefreshIdleTTL     time.Duration
	RefreshAbsoluteTTL time.Duration

	Clock Clock
}

// 错误变量。调用方可用 errors.Is 判断。
var (
	ErrTokenEmpty      = errors.New("tokenx: token is empty")
	ErrTokenInvalid    = errors.New("tokenx: token is invalid")
	ErrTokenExpired    = errors.New("tokenx: token is expired")
	ErrSessionNotFound = errors.New("tokenx: session not found")
	ErrSessionExpired  = errors.New("tokenx: session is expired")
	ErrRefreshReused   = errors.New("tokenx: refresh token reused")
	ErrCodecMismatch   = errors.New("tokenx: token codec mismatch")
)
