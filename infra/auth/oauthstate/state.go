package oauthstate

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/vkit/x/randomx"
)

const (
	keyPrefix = "blog:oauth:state:"
	stateTTL  = 10 * time.Minute
)

// ErrInvalid 表示 state 不存在、已过期或已被使用。
var ErrInvalid = errors.New("oauthstate: invalid state")

// State 是授权流程中由服务端保存的上下文。
type State struct {
	Platform string `json:"platform"`
	Intent   string `json:"intent"`
}

// Issue 生成一次性 state 并写入 Redis。
func Issue(ctx context.Context, rdb *redis.Client, platform, intent string) (string, error) {
	nonce := randomx.GenerateRandomUUID()
	payload, err := json.Marshal(State{Platform: platform, Intent: intent})
	if err != nil {
		return "", err
	}

	if err := rdb.Set(ctx, keyPrefix+nonce, payload, stateTTL).Err(); err != nil {
		return "", err
	}
	return nonce, nil
}

// Consume 读取并删除 state，保证一次性使用。
func Consume(ctx context.Context, rdb *redis.Client, nonce string) (*State, error) {
	if nonce == "" {
		return nil, ErrInvalid
	}

	payload, err := rdb.GetDel(ctx, keyPrefix+nonce).Result()
	if errors.Is(err, redis.Nil) {
		return nil, ErrInvalid
	}
	if err != nil {
		return nil, err
	}

	var state State
	if err := json.Unmarshal([]byte(payload), &state); err != nil {
		return nil, ErrInvalid
	}
	return &state, nil
}
