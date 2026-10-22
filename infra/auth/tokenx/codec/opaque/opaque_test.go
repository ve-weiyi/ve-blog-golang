package opaque_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/opaque"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {
	cases := []struct {
		kind   tokenx.TokenKind
		prefix string
	}{
		{tokenx.KindAccess, "at_"},
		{tokenx.KindRefresh, "rt_"},
	}

	for _, tc := range cases {
		t.Run(string(tc.kind), func(t *testing.T) {
			codec := opaque.New()
			encoded, err := codec.Encode(context.Background(), tokenx.Claims{Kind: tc.kind})
			if err != nil {
				t.Fatalf("Encode: %v", err)
			}
			if len(encoded.Value) != 46 {
				t.Fatalf("token length = %d, want 46", len(encoded.Value))
			}
			if !strings.HasPrefix(encoded.Value, tc.prefix) {
				t.Fatalf("token prefix = %q, want %q", encoded.Value[:3], tc.prefix)
			}
			if encoded.LookupHash == "" {
				t.Fatal("opaque token should return LookupHash")
			}

			decoded, err := codec.Decode(context.Background(), encoded.Value)
			if err != nil {
				t.Fatalf("Decode: %v", err)
			}
			if decoded.Kind != tc.kind {
				t.Fatalf("kind = %q, want %q", decoded.Kind, tc.kind)
			}
			if decoded.LookupHash != encoded.LookupHash {
				t.Fatal("LookupHash should be stable")
			}
			if decoded.Claims != nil {
				t.Fatal("opaque token should not carry claims")
			}
		})
	}
}

func TestEncodeProducesUniqueTokens(t *testing.T) {
	codec := opaque.New()
	claims := tokenx.Claims{Kind: tokenx.KindAccess}

	first, err := codec.Encode(context.Background(), claims)
	if err != nil {
		t.Fatal(err)
	}
	second, err := codec.Encode(context.Background(), claims)
	if err != nil {
		t.Fatal(err)
	}
	if first.Value == second.Value || first.LookupHash == second.LookupHash {
		t.Fatal("two issuances should produce different tokens")
	}
}

func TestEncodeRejectsInvalidKind(t *testing.T) {
	codec := opaque.New()
	if _, err := codec.Encode(context.Background(), tokenx.Claims{}); err == nil {
		t.Fatal("invalid kind should be rejected")
	}
}

func TestDecodeRejectsMalformed(t *testing.T) {
	codec := opaque.New()
	cases := []string{
		"",
		"at_",
		"at_" + strings.Repeat("a", 42),
		"xx_" + strings.Repeat("a", 43),
		"at_" + strings.Repeat("!", 43),
	}
	for _, raw := range cases {
		if _, err := codec.Decode(context.Background(), raw); !errors.Is(err, tokenx.ErrTokenInvalid) {
			t.Errorf("Decode(%q) = %v, want ErrTokenInvalid", raw, err)
		}
	}
}
