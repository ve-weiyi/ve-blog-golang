package tokenx

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
)

const (
	defaultAccessTTL          = 120 * time.Minute
	defaultRefreshIdleTTL     = 7 * 24 * time.Hour
	defaultRefreshAbsoluteTTL = 30 * 24 * time.Hour
)

type service struct {
	accessCodec  Codec
	refreshCodec Codec
	store        SessionStore
	loginMode    LoginMode
	accessTTL    time.Duration
	idleTTL      time.Duration
	absoluteTTL  time.Duration
	clock        Clock
}

// New 创建令牌服务。
func New(cfg Config) (Service, error) {
	if cfg.AccessCodec == nil {
		return nil, errors.New("tokenx: AccessCodec is required")
	}
	if cfg.RefreshCodec == nil {
		cfg.RefreshCodec = cfg.AccessCodec
	}
	if cfg.AccessCodec.Name() == "" || cfg.RefreshCodec.Name() == "" {
		return nil, errors.New("tokenx: codec name is required")
	}
	if cfg.Store == nil {
		return nil, errors.New("tokenx: Store is required")
	}

	accessTTL := cfg.AccessTTL
	if accessTTL <= 0 {
		accessTTL = defaultAccessTTL
	}
	idleTTL := cfg.RefreshIdleTTL
	if idleTTL <= 0 {
		idleTTL = defaultRefreshIdleTTL
	}
	absoluteTTL := cfg.RefreshAbsoluteTTL
	if absoluteTTL <= 0 {
		absoluteTTL = defaultRefreshAbsoluteTTL
	}

	clock := cfg.Clock
	if clock == nil {
		clock = SystemClock{}
	}

	return &service{
		accessCodec:  cfg.AccessCodec,
		refreshCodec: cfg.RefreshCodec,
		store:        cfg.Store,
		loginMode:    cfg.LoginMode,
		accessTTL:    accessTTL,
		idleTTL:      idleTTL,
		absoluteTTL:  absoluteTTL,
		clock:        clock,
	}, nil
}

// Issue 签发令牌对。
func (s *service) Issue(ctx context.Context, req IssueRequest) (*TokenPair, error) {
	if req.UserID == "" {
		return nil, errors.New("tokenx: user ID is required")
	}
	req.DeviceID = NormalizeDeviceID(req.DeviceID)

	now := s.clock.Now()
	familyID, err := newID()
	if err != nil {
		return nil, err
	}
	absoluteExpiresAt := now.Add(s.absoluteTTL)

	session, pair, err := s.build(ctx, req, familyID, now, absoluteExpiresAt)
	if err != nil {
		return nil, err
	}
	if err := s.store.Create(ctx, session, s.loginMode); err != nil {
		return nil, err
	}
	return pair, nil
}

// Validate 校验 access token。
func (s *service) Validate(ctx context.Context, accessToken string) (*Claims, error) {
	if accessToken == "" {
		return nil, ErrTokenEmpty
	}

	decoded, err := s.accessCodec.Decode(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	if decoded.Kind != "" && decoded.Kind != KindAccess {
		return nil, ErrTokenInvalid
	}

	var session Session
	if decoded.Claims != nil {
		session, err = s.store.GetBySession(ctx, decoded.Claims.UserID, decoded.Claims.SessionID)
	} else {
		if decoded.LookupHash == "" {
			return nil, ErrTokenInvalid
		}
		session, err = s.store.GetByTokenHash(ctx, decoded.LookupHash)
	}
	if err != nil {
		return nil, err
	}
	if session.AccessCodec != s.accessCodec.Name() {
		return nil, ErrCodecMismatch
	}

	now := s.clock.Now()
	if !now.Before(session.AccessExpiresAt.Add(ClockLeeway)) {
		return nil, ErrTokenExpired
	}

	if decoded.Claims != nil {
		return s.validateClaims(session, decoded.Claims, KindAccess, now)
	}

	return &Claims{
		UserID:    session.UserID,
		DeviceID:  session.DeviceID,
		SessionID: session.ID,
		Issuer:    session.Issuer,
		Kind:      KindAccess,
		IssuedAt:  session.CreatedAt,
		ExpiresAt: session.AccessExpiresAt,
	}, nil
}

// Refresh 校验 refresh token 并轮换会话。
func (s *service) Refresh(ctx context.Context, req RefreshRequest) (*TokenPair, error) {
	if req.RefreshToken == "" {
		return nil, ErrTokenEmpty
	}
	if req.UserID == "" {
		return nil, errors.New("tokenx: user ID is required")
	}
	req.DeviceID = NormalizeDeviceID(req.DeviceID)

	decoded, err := s.refreshCodec.Decode(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if decoded.Kind != "" && decoded.Kind != KindRefresh {
		return nil, ErrTokenInvalid
	}

	var session Session
	if decoded.Claims != nil {
		session, err = s.store.GetBySession(ctx, decoded.Claims.UserID, decoded.Claims.SessionID)
	} else {
		if decoded.LookupHash == "" {
			return nil, ErrTokenInvalid
		}
		session, err = s.store.GetByTokenHash(ctx, decoded.LookupHash)
	}
	if err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return s.handleRotation(ctx, req.UserID, req.DeviceID, req.RefreshToken)
		}
		return nil, err
	}

	if session.UserID != req.UserID || session.DeviceID != req.DeviceID {
		return nil, ErrTokenInvalid
	}
	if session.RefreshCodec != s.refreshCodec.Name() {
		return nil, ErrCodecMismatch
	}

	now := s.clock.Now()
	if !now.Before(session.RefreshExpiresAt.Add(ClockLeeway)) {
		return nil, ErrTokenExpired
	}
	// absolute 是硬上限，不随 ClockLeeway 延长。
	if !now.Before(session.AbsoluteExpiresAt) {
		return nil, ErrTokenExpired
	}
	if now.After(session.LastUsedAt.Add(s.idleTTL + ClockLeeway)) {
		return nil, ErrTokenExpired
	}

	if decoded.Claims != nil {
		if _, err := s.validateClaims(session, decoded.Claims, KindRefresh, now); err != nil {
			return nil, err
		}
	}

	issueReq := IssueRequest{
		UserID:   session.UserID,
		DeviceID: session.DeviceID,
		Issuer:   session.Issuer,
		ClientID: session.ClientID,
	}
	next, pair, err := s.build(ctx, issueReq, session.FamilyID, now, session.AbsoluteExpiresAt)
	if err != nil {
		return nil, err
	}
	next.LastUsedAt = now

	payload, err := json.Marshal(pair)
	if err != nil {
		return nil, fmt.Errorf("tokenx: marshal grace pair: %w", err)
	}
	gracePayload, err := sealGrace(req.RefreshToken, payload)
	if err != nil {
		return nil, fmt.Errorf("tokenx: seal grace pair: %w", err)
	}
	// leeway 内刷新可能已过 refresh 过期点；记录至少覆盖 grace 窗口，
	// 否则 memstore 会留下永不过期的记录，Redis 会直接跳过写入。
	rotationTTL := session.RefreshExpiresAt.Sub(now)
	if rotationTTL < GraceWindow {
		rotationTTL = GraceWindow
	}
	opts := RotateOptions{
		RotationTTL:  rotationTTL,
		GracePayload: gracePayload,
		GraceTTL:     GraceWindow,
	}

	if err := s.store.Rotate(ctx, session, next, opts); err != nil {
		if errors.Is(err, ErrSessionNotFound) {
			return s.handleRotation(ctx, req.UserID, req.DeviceID, req.RefreshToken)
		}
		return nil, err
	}

	return pair, nil
}

// handleRotation 处理旧 refresh 查不到/并发轮换失败的分支。
func (s *service) handleRotation(ctx context.Context, userID, deviceID, refreshToken string) (*TokenPair, error) {
	if refreshToken == "" {
		return nil, ErrRefreshReused
	}

	rotation, found, err := s.store.LookupRotation(ctx, userID, HashToken(refreshToken))
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ErrRefreshReused
	}
	if rotation.UserID != userID {
		return nil, ErrTokenInvalid
	}

	if len(rotation.GracePayload) > 0 {
		plaintext, err := openGrace(refreshToken, rotation.GracePayload)
		if err != nil {
			return nil, ErrTokenInvalid
		}
		var pair TokenPair
		if err := json.Unmarshal(plaintext, &pair); err != nil {
			return nil, ErrTokenInvalid
		}
		if pair.UserID != userID || pair.DeviceID != deviceID {
			return nil, ErrTokenInvalid
		}
		return &pair, nil
	}

	if rotation.FamilyID != "" {
		_ = s.store.RevokeFamily(ctx, userID, rotation.FamilyID)
	}
	return nil, ErrRefreshReused
}

// Revoke 撤销单个会话。
func (s *service) Revoke(ctx context.Context, userID, sessionID string) error {
	if userID == "" || sessionID == "" {
		return errors.New("tokenx: user ID and session ID are required")
	}
	return s.store.Delete(ctx, userID, sessionID)
}

// RevokeDevice 撤销指定设备的会话。
func (s *service) RevokeDevice(ctx context.Context, userID, deviceID string) error {
	if userID == "" {
		return errors.New("tokenx: user ID is required")
	}
	return s.store.DeleteByDevice(ctx, userID, NormalizeDeviceID(deviceID))
}

// RevokeAll 撤销用户全部会话。
func (s *service) RevokeAll(ctx context.Context, userID string) error {
	if userID == "" {
		return errors.New("tokenx: user ID is required")
	}
	return s.store.DeleteByUser(ctx, userID)
}

// ListSessions 列出用户未过期的会话。
func (s *service) ListSessions(ctx context.Context, userID string) ([]SessionInfo, error) {
	if userID == "" {
		return nil, errors.New("tokenx: user ID is required")
	}
	sessions, err := s.store.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	now := s.clock.Now()
	infos := make([]SessionInfo, 0, len(sessions))
	for _, session := range sessions {
		if now.Before(session.RefreshExpiresAt.Add(ClockLeeway)) &&
			now.Before(session.AbsoluteExpiresAt.Add(ClockLeeway)) {
			infos = append(infos, session.Info())
		}
	}

	sort.Slice(infos, func(i, j int) bool {
		if infos[i].LastUsedAt.Equal(infos[j].LastUsedAt) {
			return infos[i].ID < infos[j].ID
		}
		return infos[i].LastUsedAt.After(infos[j].LastUsedAt)
	})
	return infos, nil
}

// build 生成会话与令牌对，不落库。
func (s *service) build(ctx context.Context, req IssueRequest, familyID string, now, absoluteExpiresAt time.Time) (Session, *TokenPair, error) {
	if !now.Before(absoluteExpiresAt) {
		return Session{}, nil, ErrTokenExpired
	}

	sessionID, err := newID()
	if err != nil {
		return Session{}, nil, err
	}

	// access 与 idle 都不得超过家族绝对期限。
	refreshExpiresAt := now.Add(s.idleTTL)
	if refreshExpiresAt.After(absoluteExpiresAt) {
		refreshExpiresAt = absoluteExpiresAt
	}
	accessExpiresAt := now.Add(s.accessTTL)
	if accessExpiresAt.After(absoluteExpiresAt) {
		accessExpiresAt = absoluteExpiresAt
	}

	accessClaims := Claims{
		UserID:    req.UserID,
		DeviceID:  req.DeviceID,
		SessionID: sessionID,
		Issuer:    req.Issuer,
		Kind:      KindAccess,
		IssuedAt:  now,
		ExpiresAt: accessExpiresAt,
	}
	refreshClaims := Claims{
		UserID:    req.UserID,
		DeviceID:  req.DeviceID,
		SessionID: sessionID,
		Issuer:    req.Issuer,
		Kind:      KindRefresh,
		IssuedAt:  now,
		ExpiresAt: refreshExpiresAt,
	}

	access, err := s.accessCodec.Encode(ctx, accessClaims)
	if err != nil {
		return Session{}, nil, err
	}
	refresh, err := s.refreshCodec.Encode(ctx, refreshClaims)
	if err != nil {
		return Session{}, nil, err
	}
	if access.Value == "" || refresh.Value == "" {
		return Session{}, nil, errors.New("tokenx: codec returned empty token")
	}

	session := Session{
		ID:                sessionID,
		UserID:            req.UserID,
		DeviceID:          req.DeviceID,
		FamilyID:          familyID,
		ClientID:          req.ClientID,
		Issuer:            req.Issuer,
		AccessCodec:       s.accessCodec.Name(),
		RefreshCodec:      s.refreshCodec.Name(),
		AccessLookupHash:  access.LookupHash,
		RefreshLookupHash: refresh.LookupHash,
		RefreshTokenHash:  HashToken(refresh.Value),
		CreatedAt:         now,
		LastUsedAt:        now,
		AccessExpiresAt:   accessClaims.ExpiresAt,
		RefreshExpiresAt:  refreshClaims.ExpiresAt,
		AbsoluteExpiresAt: absoluteExpiresAt,
	}

	pair := &TokenPair{
		UserID:           req.UserID,
		DeviceID:         req.DeviceID,
		AccessToken:      access.Value,
		RefreshToken:     refresh.Value,
		SessionID:        sessionID,
		AccessExpiresAt:  accessClaims.ExpiresAt,
		AccessExpiresIn:  int64(accessClaims.ExpiresAt.Sub(now) / time.Second),
		RefreshExpiresAt: refreshClaims.ExpiresAt,
		RefreshExpiresIn: int64(refreshClaims.ExpiresAt.Sub(now) / time.Second),
	}
	return session, pair, nil
}

func (s *service) validateClaims(session Session, claims *Claims, kind TokenKind, now time.Time) (*Claims, error) {
	if claims.Kind != kind {
		return nil, ErrTokenInvalid
	}
	if claims.UserID != session.UserID ||
		claims.DeviceID != session.DeviceID ||
		claims.SessionID != session.ID {
		return nil, ErrTokenInvalid
	}
	if claims.Issuer != session.Issuer {
		return nil, ErrTokenInvalid
	}
	if claims.ExpiresAt.IsZero() || !now.Before(claims.ExpiresAt.Add(ClockLeeway)) {
		return nil, ErrTokenExpired
	}
	return claims, nil
}

func newID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("tokenx: generate id: %w", err)
	}
	return id.String(), nil
}

// HashToken 返回 token 的稳定哈希，用于 refresh 轮换记录。
func HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}
