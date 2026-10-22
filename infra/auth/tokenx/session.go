package tokenx

import "time"

// Session 是令牌服务的会话记录，由 SessionStore 持久化。
type Session struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	DeviceID string `json:"device_id"`
	FamilyID string `json:"family_id"`

	ClientID string `json:"client_id,omitempty"`
	Issuer   string `json:"issuer,omitempty"`

	AccessCodec  string `json:"access_codec"`
	RefreshCodec string `json:"refresh_codec"`

	AccessLookupHash  string `json:"access_lookup_hash,omitempty"`
	RefreshLookupHash string `json:"refresh_lookup_hash,omitempty"`
	RefreshTokenHash  string `json:"refresh_token_hash"`

	CreatedAt         time.Time `json:"created_at"`
	LastUsedAt        time.Time `json:"last_used_at"`
	AccessExpiresAt   time.Time `json:"access_expires_at"`
	RefreshExpiresAt  time.Time `json:"refresh_expires_at"`
	AbsoluteExpiresAt time.Time `json:"absolute_expires_at"`
}

// Info 返回脱敏后的会话视图。
func (s Session) Info() SessionInfo {
	return SessionInfo{
		ID:                s.ID,
		UserID:            s.UserID,
		DeviceID:          s.DeviceID,
		ClientID:          s.ClientID,
		CreatedAt:         s.CreatedAt,
		LastUsedAt:        s.LastUsedAt,
		AccessExpiresAt:   s.AccessExpiresAt,
		RefreshExpiresAt:  s.RefreshExpiresAt,
		AbsoluteExpiresAt: s.AbsoluteExpiresAt,
	}
}
