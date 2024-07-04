package middlewarex

import (
	"net/http"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/client/authrpc"
)

type SsoMiddleware struct {
	AuthRpc authrpc.AuthRpc
}

func NewSsoMiddleware(auth authrpc.AuthRpc) *SsoMiddleware {
	return &SsoMiddleware{
		AuthRpc: auth,
	}
}

// 单点登录（single sign on）
func (m *SsoMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("SsoMiddleware Handle")
		uid := r.Header.Get(constant.HeaderUid)
		ts := r.Header.Get(constant.HeaderTimestamp)

		at, err := m.AuthRpc.GetLogoutAt(r.Context(), &authrpc.GetLogoutAtReq{
			UserId: cast.ToInt64(uid),
		})
		if err != nil {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage(err.Error()))
			return
		}

		if at.LogoutAt > cast.ToInt64(ts) {
			responsex.Response(r, w, nil, apierr.ErrorUnauthorized.WrapMessage("user already logout or login in other place"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
