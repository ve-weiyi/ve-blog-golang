// Package identityx 定义用户与游客标识的生成与形态判定。
//
// 标识形态是对外契约，因此 NewGuestID 直接按形态生成，不依赖通用随机工具。
package identityx

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/google/uuid"
)

// Kind 是身份标识的主体类型。
type Kind uint8

const (
	KindUnknown Kind = iota
	KindUser
	KindGuest
)

const (
	userIDLength  = 36
	guestIDLength = 32
)

// NewID 生成通用的 UUIDv7 标识：36 位、小写、带横线。
func NewID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// NewUserID 生成 UUIDv7 形式的用户标识。
func NewUserID() (string, error) {
	return NewID()
}

// NewGuestID 生成 32 位小写十六进制的游客标识，代表一个客户端实例。
func NewGuestID() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

// KindOf 依据标识的规范形态判断主体类型：
// 用户为 36 位小写 UUIDv7，游客为 32 位小写十六进制，其余为未知。
//
// 判定只是提示，不是保证：历史用户标识可能呈游客形态。
// 消费方在查无结果时必须回查另一类，不能据此丢弃记录。
func KindOf(id string) Kind {
	if len(id) == userIDLength {
		if u, err := uuid.Parse(id); err == nil && u.String() == id && u.Version() == uuid.Version(7) {
			return KindUser
		}
	}
	if isLowerHex32(id) {
		return KindGuest
	}
	return KindUnknown
}

func isLowerHex32(s string) bool {
	if len(s) != guestIDLength {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}
