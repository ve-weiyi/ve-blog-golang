package tokenx

import (
	"context"
	"time"
)

// SessionStore 负责会话、token 索引、family 与轮换记录的持久化。
//
// 不变式：同一 user + device 同时只保留一个活跃会话。
// Create/Rotate 必须在实现内原子完成。
type SessionStore interface {
	// Create 保存新会话。
	//   - MultiPoint：删除同 user+device 的旧会话后保存；
	//   - SinglePoint：删除该 user 的全部旧会话后保存。
	Create(ctx context.Context, session Session, mode LoginMode) error

	// GetBySession 按 userID + sessionID 读取会话。
	GetBySession(ctx context.Context, userID, sessionID string) (Session, error)

	// GetByTokenHash 按不透明 token 的 LookupHash 读取会话。
	GetByTokenHash(ctx context.Context, tokenHash string) (Session, error)

	// Rotate 原子替换旧会话，并写入 rotation/grace 记录。
	Rotate(ctx context.Context, old Session, next Session, opts RotateOptions) error

	// LookupRotation 按旧 refresh token 哈希读取轮换记录。
	LookupRotation(ctx context.Context, userID, oldRefreshHash string) (Rotation, bool, error)

	// RevokeFamily 撤销整个 family 的当前会话。
	RevokeFamily(ctx context.Context, userID, familyID string) error

	Delete(ctx context.Context, userID, sessionID string) error
	DeleteByDevice(ctx context.Context, userID, deviceID string) error
	DeleteByUser(ctx context.Context, userID string) error

	ListByUser(ctx context.Context, userID string) ([]Session, error)
}

// Rotation 是一次 refresh 轮换留下的记录。
type Rotation struct {
	UserID       string
	FamilyID     string
	GracePayload []byte // GraceWindow 内有效；过期后为空
}

// RotateOptions 控制 Rotate 写入的辅助记录。
type RotateOptions struct {
	GracePayload []byte
	RotationTTL  time.Duration // 旧 refresh 剩余有效期
	GraceTTL     time.Duration // 固定为 GraceWindow
}
