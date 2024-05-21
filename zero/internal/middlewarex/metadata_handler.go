package middlewarex

import (
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
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
				for _, key := range constant.HeaderFields {
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
