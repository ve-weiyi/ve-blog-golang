package middlewarex

import (
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/constantx"
)

// CtxMetadataHandel 将http header 放入 ctx 里面使用 metadata 保存.
func CtxMetadataHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//logx.Infof("CtxMetadataHandel")
		ctx := r.Context()
		md := metadata.MD{}
		// 从真实http 头获取header
		if len(r.Header) > 0 {
			for k, v := range r.Header {
				if len(k) == 0 {
					continue
				}
				var value string
				if len(v) > 0 {
					value = v[0]
				}
				keyLowercase := strings.ToLower(k)
				//logx.Infof("add k=%s, v=%+v", keyLowercase, value)
				for _, key := range constantx.HeaderFields {
					if key == keyLowercase {
						md.Set(keyLowercase, value)
					}
				}
			}
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}

func AppendToOutgoingContextInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		var pairs []string
		for k, v := range md {
			for _, value := range v {
				pairs = append(pairs, k, value)
			}
		}
		ctx = metadata.AppendToOutgoingContext(ctx, pairs...)
	}

	resp, err := handler(ctx, req)
	if err != nil {
		logx.Error("grpc server error", err)
	}
	return resp, err
}
