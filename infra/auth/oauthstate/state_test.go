package oauthstate

import (
	"context"
	"errors"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func TestIssueConsume(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	ctx := context.Background()
	nonce, err := Issue(ctx, rdb, "github", "bind_account")
	if err != nil {
		t.Fatalf("Issue: %v", err)
	}
	if nonce == "" {
		t.Fatal("nonce is empty")
	}

	state, err := Consume(ctx, rdb, nonce)
	if err != nil {
		t.Fatalf("Consume: %v", err)
	}
	if state.Platform != "github" || state.Intent != "bind_account" {
		t.Fatalf("state = %+v", state)
	}

	if _, err := Consume(ctx, rdb, nonce); !errors.Is(err, ErrInvalid) {
		t.Fatalf("second Consume err = %v, want ErrInvalid", err)
	}
}

func TestConsumeInvalid(t *testing.T) {
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	t.Cleanup(func() { _ = rdb.Close() })

	if _, err := Consume(context.Background(), rdb, "missing"); !errors.Is(err, ErrInvalid) {
		t.Fatalf("err = %v, want ErrInvalid", err)
	}
}
