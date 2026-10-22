// Package memstore 提供 tokenx.SessionStore 的内存实现，供测试与本地开发使用。
package memstore

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
)

// Store 是基于内存的会话存储。
type Store struct {
	mu        sync.Mutex
	clock     tokenx.Clock
	sessions  map[string]map[string]tokenx.Session // userID -> sessionID -> session
	devices   map[string]map[string]string         // userID -> deviceID -> sessionID
	tokens    map[string]tokenRef                  // lookup hash -> session ref
	families  map[string]familyRef                 // familyID -> current session
	rotations map[string]rotationRecord            // old refresh hash -> rotation
}

type tokenRef struct {
	userID    string
	sessionID string
	kind      tokenx.TokenKind
}

type familyRef struct {
	userID    string
	sessionID string
}

type rotationRecord struct {
	userID            string
	familyID          string
	gracePayload      []byte
	graceExpiresAt    time.Time
	rotationExpiresAt time.Time
}

// New 创建内存会话存储。
func New() *Store {
	return NewWithClock(tokenx.SystemClock{})
}

// NewWithClock 创建使用指定时间源的内存会话存储。
func NewWithClock(clock tokenx.Clock) *Store {
	if clock == nil {
		clock = tokenx.SystemClock{}
	}
	return &Store{
		clock:     clock,
		sessions:  make(map[string]map[string]tokenx.Session),
		devices:   make(map[string]map[string]string),
		tokens:    make(map[string]tokenRef),
		families:  make(map[string]familyRef),
		rotations: make(map[string]rotationRecord),
	}
}

func (s *Store) Create(_ context.Context, session tokenx.Session, mode tokenx.LoginMode) error {
	if err := validateSession(session); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if mode == tokenx.SinglePoint {
		for _, old := range s.sessions[session.UserID] {
			s.deleteSessionLocked(old)
		}
		delete(s.sessions, session.UserID)
		delete(s.devices, session.UserID)
	} else if oldID, ok := s.devices[session.UserID][session.DeviceID]; ok {
		if old, ok := s.sessions[session.UserID][oldID]; ok {
			s.deleteSessionLocked(old)
		}
	}

	s.addSessionLocked(session)
	return nil
}

func (s *Store) GetBySession(_ context.Context, userID, sessionID string) (tokenx.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, ok := s.sessions[userID][sessionID]
	if !ok {
		return tokenx.Session{}, tokenx.ErrSessionNotFound
	}
	return cloneSession(session), nil
}

func (s *Store) GetByTokenHash(_ context.Context, tokenHash string) (tokenx.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ref, ok := s.tokens[tokenHash]
	if !ok {
		return tokenx.Session{}, tokenx.ErrSessionNotFound
	}
	session, ok := s.sessions[ref.userID][ref.sessionID]
	if !ok {
		return tokenx.Session{}, tokenx.ErrSessionNotFound
	}
	return cloneSession(session), nil
}

func (s *Store) Rotate(_ context.Context, old tokenx.Session, next tokenx.Session, opts tokenx.RotateOptions) error {
	if err := validateSession(next); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.sessions[old.UserID][old.ID]; !ok {
		return tokenx.ErrSessionNotFound
	}
	if old.UserID != next.UserID || old.DeviceID != next.DeviceID {
		return errors.New("memstore: rotate must keep the same user and device")
	}
	next.FamilyID = old.FamilyID

	s.deleteSessionLocked(old)
	s.addSessionLocked(next)

	now := s.clock.Now()
	record := rotationRecord{
		userID:   old.UserID,
		familyID: old.FamilyID,
	}
	record.rotationExpiresAt = now.Add(opts.RotationTTL)
	if len(opts.GracePayload) > 0 && opts.GraceTTL > 0 {
		record.gracePayload = append([]byte(nil), opts.GracePayload...)
		record.graceExpiresAt = now.Add(opts.GraceTTL)
	}
	s.rotations[old.RefreshTokenHash] = record
	return nil
}

func (s *Store) LookupRotation(_ context.Context, userID, oldRefreshHash string) (tokenx.Rotation, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	record, ok := s.rotations[oldRefreshHash]
	if !ok || record.userID != userID {
		return tokenx.Rotation{}, false, nil
	}
	now := s.clock.Now()
	if !record.rotationExpiresAt.IsZero() && !now.Before(record.rotationExpiresAt) {
		return tokenx.Rotation{}, false, nil
	}

	rotation := tokenx.Rotation{
		UserID:   record.userID,
		FamilyID: record.familyID,
	}
	if len(record.gracePayload) > 0 && now.Before(record.graceExpiresAt) {
		rotation.GracePayload = append([]byte(nil), record.gracePayload...)
	}
	return rotation, true, nil
}

func (s *Store) RevokeFamily(_ context.Context, userID, familyID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ref, ok := s.families[familyID]
	if !ok || ref.userID != userID {
		return nil
	}
	if session, ok := s.sessions[userID][ref.sessionID]; ok {
		s.deleteSessionLocked(session)
	}
	return nil
}

func (s *Store) Delete(_ context.Context, userID, sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	session, ok := s.sessions[userID][sessionID]
	if !ok {
		return nil
	}
	s.deleteSessionLocked(session)
	return nil
}

func (s *Store) DeleteByDevice(_ context.Context, userID, deviceID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	sessionID, ok := s.devices[userID][deviceID]
	if !ok {
		return nil
	}
	if session, ok := s.sessions[userID][sessionID]; ok {
		s.deleteSessionLocked(session)
	}
	return nil
}

func (s *Store) DeleteByUser(_ context.Context, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, session := range s.sessions[userID] {
		s.deleteSessionLocked(session)
	}
	delete(s.sessions, userID)
	delete(s.devices, userID)
	for hash, record := range s.rotations {
		if record.userID == userID {
			delete(s.rotations, hash)
		}
	}
	return nil
}

func (s *Store) ListByUser(_ context.Context, userID string) ([]tokenx.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	sessions := make([]tokenx.Session, 0, len(s.sessions[userID]))
	for _, session := range s.sessions[userID] {
		sessions = append(sessions, cloneSession(session))
	}
	return sessions, nil
}

func (s *Store) addSessionLocked(session tokenx.Session) {
	if s.sessions[session.UserID] == nil {
		s.sessions[session.UserID] = make(map[string]tokenx.Session)
	}
	if s.devices[session.UserID] == nil {
		s.devices[session.UserID] = make(map[string]string)
	}

	s.sessions[session.UserID][session.ID] = cloneSession(session)
	s.devices[session.UserID][session.DeviceID] = session.ID
	s.families[session.FamilyID] = familyRef{userID: session.UserID, sessionID: session.ID}
	if session.AccessLookupHash != "" {
		s.tokens[session.AccessLookupHash] = tokenRef{userID: session.UserID, sessionID: session.ID, kind: tokenx.KindAccess}
	}
	if session.RefreshLookupHash != "" {
		s.tokens[session.RefreshLookupHash] = tokenRef{userID: session.UserID, sessionID: session.ID, kind: tokenx.KindRefresh}
	}
}

func (s *Store) deleteSessionLocked(session tokenx.Session) {
	delete(s.sessions[session.UserID], session.ID)
	if s.devices[session.UserID][session.DeviceID] == session.ID {
		delete(s.devices[session.UserID], session.DeviceID)
	}
	if s.families[session.FamilyID].sessionID == session.ID {
		delete(s.families, session.FamilyID)
	}
	if session.AccessLookupHash != "" {
		delete(s.tokens, session.AccessLookupHash)
	}
	if session.RefreshLookupHash != "" {
		delete(s.tokens, session.RefreshLookupHash)
	}
}

func validateSession(session tokenx.Session) error {
	if session.ID == "" || session.UserID == "" || session.DeviceID == "" || session.FamilyID == "" {
		return errors.New("memstore: session, user, device and family id are required")
	}
	if session.AccessCodec == "" || session.RefreshCodec == "" {
		return errors.New("memstore: codec names are required")
	}
	return nil
}

func cloneSession(session tokenx.Session) tokenx.Session {
	return session
}
