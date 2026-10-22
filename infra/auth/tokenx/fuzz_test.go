package tokenx_test

import (
	"context"
	"strings"
	"testing"

	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/hmaccodec"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/jwtcodec"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/codec/opaque"
	"github.com/ve-weiyi/blog-cloud/infra/auth/tokenx/store/memstore"
)

// FuzzDecode 保证畸形输入只返回错误，不 panic。
func FuzzDecode(f *testing.F) {
	jwtCodec, err := jwtcodec.New(testSecret)
	if err != nil {
		f.Fatal(err)
	}
	hmacCodec, err := hmaccodec.New(testSecret)
	if err != nil {
		f.Fatal(err)
	}
	codecs := []tokenx.Codec{opaque.New(), jwtCodec, hmacCodec}

	svc, err := tokenx.New(tokenx.Config{
		AccessCodec: opaque.New(),
		Store:       memstore.New(),
	})
	if err != nil {
		f.Fatal(err)
	}

	for _, seed := range []string{"", "at_", "rt_", "v1.x.y", "a.b.c", "..", "at_" + strings.Repeat("A", 43)} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, raw string) {
		for _, codec := range codecs {
			_, _ = codec.Decode(context.Background(), raw)
		}
		_, _ = svc.Validate(context.Background(), raw)
	})
}
