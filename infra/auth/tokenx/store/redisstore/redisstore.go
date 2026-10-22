// Package redisstore 提供 tokenx.SessionStore 的 Redis 实现。
//
// 所有 key 共享 {tokenx} hash tag；session 独立 TTL，token/family/rotation 各自建索引。
package redisstore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
)

const (
	keyTTLSlack = 5 * time.Minute
	keyTag      = "{tokenx}"
)

// Store 是 Redis 会话存储。
type Store struct {
	client redis.UniversalClient
	base   string
}

// New 创建 Redis 会话存储。prefix 例如 "blog:admin:token"。
func New(client redis.UniversalClient, prefix string) (*Store, error) {
	if client == nil {
		return nil, errors.New("redisstore: redis client is required")
	}
	prefix = strings.TrimSuffix(strings.TrimSpace(prefix), ":")
	if prefix == "" {
		return nil, errors.New("redisstore: key prefix is required")
	}
	return &Store{client: client, base: prefix + ":" + keyTag + ":"}, nil
}

func (s *Store) Create(ctx context.Context, session tokenx.Session, mode tokenx.LoginMode) error {
	if err := validateSession(session); err != nil {
		return err
	}
	ttl, err := ttlSeconds(session)
	if err != nil {
		return err
	}
	payload, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("redisstore: marshal session: %w", err)
	}

	modeArg := "multi"
	if mode == tokenx.SinglePoint {
		modeArg = "single"
	}

	_, err = createScript.Run(ctx, s.client,
		[]string{
			s.sessionKey(session.ID),
			s.sessionTokensKey(session.ID),
			s.sessionFamilyKey(session.ID),
			s.sessionDeviceKey(session.ID),
			s.userSessionsKey(session.UserID),
			s.userDeviceKey(session.UserID, session.DeviceID),
		},
		session.ID, session.UserID, session.DeviceID, session.FamilyID,
		string(payload), ttl, modeArg, session.AccessLookupHash, session.RefreshLookupHash, s.base,
	).Result()
	if err != nil {
		return fmt.Errorf("redisstore: create session: %w", err)
	}
	return nil
}

func (s *Store) GetBySession(ctx context.Context, userID, sessionID string) (tokenx.Session, error) {
	data, err := s.client.Get(ctx, s.sessionKey(sessionID)).Result()
	if err == redis.Nil {
		return tokenx.Session{}, tokenx.ErrSessionNotFound
	}
	if err != nil {
		return tokenx.Session{}, fmt.Errorf("redisstore: get session: %w", err)
	}

	var session tokenx.Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		return tokenx.Session{}, fmt.Errorf("redisstore: unmarshal session: %w", err)
	}
	if session.UserID != userID {
		return tokenx.Session{}, tokenx.ErrSessionNotFound
	}
	return session, nil
}

func (s *Store) GetByTokenHash(ctx context.Context, tokenHash string) (tokenx.Session, error) {
	value, err := s.client.Get(ctx, s.tokenKey(tokenHash)).Result()
	if err == redis.Nil {
		return tokenx.Session{}, tokenx.ErrSessionNotFound
	}
	if err != nil {
		return tokenx.Session{}, fmt.Errorf("redisstore: get token index: %w", err)
	}

	parts := strings.SplitN(value, "\t", 3)
	if len(parts) != 3 {
		return tokenx.Session{}, fmt.Errorf("redisstore: invalid token index")
	}
	return s.GetBySession(ctx, parts[0], parts[1])
}

func (s *Store) Rotate(ctx context.Context, old tokenx.Session, next tokenx.Session, opts tokenx.RotateOptions) error {
	if err := validateSession(next); err != nil {
		return err
	}
	current, err := s.GetBySession(ctx, old.UserID, old.ID)
	if err != nil {
		return err
	}
	if current.UserID != next.UserID || current.DeviceID != next.DeviceID {
		return errors.New("redisstore: rotate must keep the same user and device")
	}
	next.FamilyID = current.FamilyID

	ttl, err := ttlSeconds(next)
	if err != nil {
		return err
	}
	payload, err := json.Marshal(next)
	if err != nil {
		return fmt.Errorf("redisstore: marshal session: %w", err)
	}

	rotationTTL := secondsFor(opts.RotationTTL)
	graceTTL := secondsFor(opts.GraceTTL)

	result, err := rotateScript.Run(ctx, s.client,
		[]string{
			s.sessionKey(next.ID),
			s.sessionTokensKey(next.ID),
			s.sessionFamilyKey(next.ID),
			s.sessionDeviceKey(next.ID),
			s.userSessionsKey(next.UserID),
			s.userDeviceKey(next.UserID, next.DeviceID),
		},
		old.ID, next.ID, next.UserID, next.DeviceID, next.FamilyID, old.RefreshTokenHash,
		string(payload), ttl, next.AccessLookupHash, next.RefreshLookupHash,
		rotationTTL, graceTTL, string(opts.GracePayload), s.base,
	).Int()
	if err != nil {
		return fmt.Errorf("redisstore: rotate session: %w", err)
	}
	if result == 0 {
		return tokenx.ErrSessionNotFound
	}
	return nil
}

func (s *Store) LookupRotation(ctx context.Context, userID, oldRefreshHash string) (tokenx.Rotation, bool, error) {
	value, err := s.client.Get(ctx, s.rotationKey(oldRefreshHash)).Result()
	if err == redis.Nil {
		return tokenx.Rotation{}, false, nil
	}
	if err != nil {
		return tokenx.Rotation{}, false, fmt.Errorf("redisstore: get rotation: %w", err)
	}

	parts := strings.SplitN(value, "\t", 2)
	if len(parts) != 2 {
		return tokenx.Rotation{}, false, fmt.Errorf("redisstore: invalid rotation record")
	}
	if parts[0] != userID {
		return tokenx.Rotation{}, false, nil
	}

	grace, err := s.client.Get(ctx, s.rotationGraceKey(oldRefreshHash)).Bytes()
	if err != nil && err != redis.Nil {
		return tokenx.Rotation{}, false, fmt.Errorf("redisstore: get rotation grace: %w", err)
	}

	return tokenx.Rotation{
		UserID:       parts[0],
		FamilyID:     parts[1],
		GracePayload: grace,
	}, true, nil
}

func (s *Store) RevokeFamily(ctx context.Context, userID, familyID string) error {
	value, err := s.client.Get(ctx, s.familyKey(familyID)).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return fmt.Errorf("redisstore: get family: %w", err)
	}
	parts := strings.SplitN(value, "\t", 2)
	if len(parts) != 2 || parts[0] != userID {
		return nil
	}

	_, err = deleteScript.Run(ctx, s.client, nil, userID, parts[1], s.base).Result()
	if err != nil {
		return fmt.Errorf("redisstore: revoke family: %w", err)
	}
	return nil
}

func (s *Store) Delete(ctx context.Context, userID, sessionID string) error {
	// 先确认会话归属，避免用不匹配的 userID 删除他人会话。
	if _, err := s.GetBySession(ctx, userID, sessionID); err != nil {
		if errors.Is(err, tokenx.ErrSessionNotFound) {
			return nil
		}
		return err
	}
	_, err := deleteScript.Run(ctx, s.client, nil, userID, sessionID, s.base).Result()
	if err != nil {
		return fmt.Errorf("redisstore: delete session: %w", err)
	}
	return nil
}

func (s *Store) DeleteByDevice(ctx context.Context, userID, deviceID string) error {
	_, err := deleteDeviceScript.Run(ctx, s.client, nil, userID, deviceID, s.base).Result()
	if err != nil {
		return fmt.Errorf("redisstore: delete device session: %w", err)
	}
	return nil
}

func (s *Store) DeleteByUser(ctx context.Context, userID string) error {
	_, err := deleteUserScript.Run(ctx, s.client, nil, userID, s.base).Result()
	if err != nil {
		return fmt.Errorf("redisstore: delete user sessions: %w", err)
	}
	return nil
}

func (s *Store) ListByUser(ctx context.Context, userID string) ([]tokenx.Session, error) {
	sessionIDs, err := s.client.SMembers(ctx, s.userSessionsKey(userID)).Result()
	if err != nil {
		return nil, fmt.Errorf("redisstore: list session ids: %w", err)
	}
	if len(sessionIDs) == 0 {
		return nil, nil
	}

	keys := make([]string, len(sessionIDs))
	for i, sid := range sessionIDs {
		keys[i] = s.sessionKey(sid)
	}
	values, err := s.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("redisstore: list sessions: %w", err)
	}

	sessions := make([]tokenx.Session, 0, len(values))
	for _, value := range values {
		raw, ok := value.(string)
		if !ok {
			continue
		}
		var session tokenx.Session
		if err := json.Unmarshal([]byte(raw), &session); err != nil {
			return nil, fmt.Errorf("redisstore: unmarshal session: %w", err)
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (s *Store) sessionKey(sessionID string) string { return s.base + "session:" + sessionID }
func (s *Store) sessionTokensKey(sessionID string) string {
	return s.base + "session:" + sessionID + ":tokens"
}
func (s *Store) sessionFamilyKey(sessionID string) string {
	return s.base + "session:" + sessionID + ":family"
}
func (s *Store) sessionDeviceKey(sessionID string) string {
	return s.base + "session:" + sessionID + ":device"
}
func (s *Store) userSessionsKey(userID string) string {
	return s.base + "user:" + userID + ":sessions"
}
func (s *Store) userDeviceKey(userID, deviceID string) string {
	return s.base + "user:" + userID + ":device:" + deviceID
}
func (s *Store) tokenKey(hash string) string      { return s.base + "token:" + hash }
func (s *Store) familyKey(familyID string) string { return s.base + "family:" + familyID }
func (s *Store) rotationKey(hash string) string   { return s.base + "rotation:" + hash }
func (s *Store) rotationGraceKey(hash string) string {
	return s.base + "rotation-grace:" + hash
}

func ttlSeconds(session tokenx.Session) (int64, error) {
	ttl := session.RefreshExpiresAt.Sub(session.CreatedAt) + keyTTLSlack
	if ttl <= 0 {
		return 0, tokenx.ErrSessionExpired
	}
	return int64(math.Ceil(ttl.Seconds())), nil
}

func secondsFor(d time.Duration) int64 {
	seconds := int64(math.Ceil(d.Seconds()))
	if seconds < 1 {
		return 0
	}
	return seconds
}

func validateSession(session tokenx.Session) error {
	if session.ID == "" || session.UserID == "" || session.DeviceID == "" || session.FamilyID == "" {
		return errors.New("redisstore: session, user, device and family id are required")
	}
	if session.AccessCodec == "" || session.RefreshCodec == "" {
		return errors.New("redisstore: codec names are required")
	}
	return nil
}

const luaHelpers = `
local function delSession(base, uid, sid)
  local s = base .. "session:" .. sid
  local t = s .. ":tokens"
  local f = s .. ":family"
  local d = s .. ":device"
  local hashes = redis.call('SMEMBERS', t)
  for _, h in ipairs(hashes) do
    redis.call('DEL', base .. "token:" .. h)
  end
  local fid = redis.call('GET', f)
  if fid and redis.call('GET', base .. "family:" .. fid) == uid .. "\t" .. sid then
    redis.call('DEL', base .. "family:" .. fid)
  end
  local did = redis.call('GET', d)
  if did and redis.call('GET', base .. "user:" .. uid .. ":device:" .. did) == sid then
    redis.call('DEL', base .. "user:" .. uid .. ":device:" .. did)
  end
  redis.call('DEL', s, t, f, d)
  redis.call('SREM', base .. "user:" .. uid .. ":sessions", sid)
end
`

var createScript = redis.NewScript(luaHelpers + `
local base = ARGV[10]
local sid = ARGV[1]
local uid = ARGV[2]
local did = ARGV[3]
local fid = ARGV[4]
local data = ARGV[5]
local ttl = tonumber(ARGV[6])
local mode = ARGV[7]
local accessHash = ARGV[8]
local refreshHash = ARGV[9]

if mode == 'single' then
  local olds = redis.call('SMEMBERS', KEYS[5])
  for _, old in ipairs(olds) do
    delSession(base, uid, old)
  end
  redis.call('DEL', KEYS[5])
else
  local old = redis.call('GET', KEYS[6])
  if old then
    delSession(base, uid, old)
  end
end

redis.call('SET', KEYS[1], data, 'EX', ttl)
redis.call('SET', KEYS[3], fid, 'EX', ttl)
redis.call('SET', KEYS[4], did, 'EX', ttl)
redis.call('SET', base .. "family:" .. fid, uid .. "\t" .. sid, 'EX', ttl)
redis.call('SET', base .. "user:" .. uid .. ":device:" .. did, sid, 'EX', ttl)
redis.call('SADD', KEYS[5], sid)
redis.call('EXPIRE', KEYS[5], ttl)
redis.call('DEL', KEYS[2])
if accessHash ~= '' then
  redis.call('SADD', KEYS[2], accessHash)
  redis.call('SET', base .. "token:" .. accessHash, uid .. "\t" .. sid .. "\taccess", 'EX', ttl)
end
if refreshHash ~= '' then
  redis.call('SADD', KEYS[2], refreshHash)
  redis.call('SET', base .. "token:" .. refreshHash, uid .. "\t" .. sid .. "\trefresh", 'EX', ttl)
end
redis.call('EXPIRE', KEYS[2], ttl)
return 1
`)

var rotateScript = redis.NewScript(luaHelpers + `
local base = ARGV[14]
local oldSid = ARGV[1]
local newSid = ARGV[2]
local uid = ARGV[3]
local did = ARGV[4]
local fid = ARGV[5]
local oldRefreshHash = ARGV[6]
local data = ARGV[7]
local ttl = tonumber(ARGV[8])
local accessHash = ARGV[9]
local refreshHash = ARGV[10]
local rotationTTL = tonumber(ARGV[11])
local graceTTL = tonumber(ARGV[12])
local gracePayload = ARGV[13]

if redis.call('EXISTS', base .. "session:" .. oldSid) == 0 then
  return 0
end
delSession(base, uid, oldSid)

redis.call('SET', KEYS[1], data, 'EX', ttl)
redis.call('SET', KEYS[3], fid, 'EX', ttl)
redis.call('SET', KEYS[4], did, 'EX', ttl)
redis.call('SET', base .. "family:" .. fid, uid .. "\t" .. newSid, 'EX', ttl)
redis.call('SET', base .. "user:" .. uid .. ":device:" .. did, newSid, 'EX', ttl)
redis.call('SADD', KEYS[5], newSid)
redis.call('EXPIRE', KEYS[5], ttl)
redis.call('DEL', KEYS[2])
if accessHash ~= '' then
  redis.call('SADD', KEYS[2], accessHash)
  redis.call('SET', base .. "token:" .. accessHash, uid .. "\t" .. newSid .. "\taccess", 'EX', ttl)
end
if refreshHash ~= '' then
  redis.call('SADD', KEYS[2], refreshHash)
  redis.call('SET', base .. "token:" .. refreshHash, uid .. "\t" .. newSid .. "\trefresh", 'EX', ttl)
end
redis.call('EXPIRE', KEYS[2], ttl)

if oldRefreshHash ~= '' and rotationTTL > 0 then
  redis.call('SET', base .. "rotation:" .. oldRefreshHash, uid .. "\t" .. fid, 'EX', rotationTTL)
end
if oldRefreshHash ~= '' and graceTTL > 0 and gracePayload ~= '' then
  redis.call('SET', base .. "rotation-grace:" .. oldRefreshHash, gracePayload, 'EX', graceTTL)
end
return 1
`)

var deleteScript = redis.NewScript(luaHelpers + `
local uid = ARGV[1]
local sid = ARGV[2]
local base = ARGV[3]
delSession(base, uid, sid)
return 1
`)

var deleteDeviceScript = redis.NewScript(luaHelpers + `
local uid = ARGV[1]
local did = ARGV[2]
local base = ARGV[3]
local sid = redis.call('GET', base .. "user:" .. uid .. ":device:" .. did)
if sid then
  delSession(base, uid, sid)
end
return 1
`)

var deleteUserScript = redis.NewScript(luaHelpers + `
local uid = ARGV[1]
local base = ARGV[2]
local olds = redis.call('SMEMBERS', base .. "user:" .. uid .. ":sessions")
for _, sid in ipairs(olds) do
  delSession(base, uid, sid)
end
redis.call('DEL', base .. "user:" .. uid .. ":sessions")
return 1
`)
