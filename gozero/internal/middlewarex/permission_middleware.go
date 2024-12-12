package middlewarex

import (
	"context"
	"net/http"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
)

type PermissionHolder interface {
	Enforce(ctx context.Context, user string, resource string, action string) error
}

type PermissionMiddleware struct {
	holder PermissionHolder
}

func NewPermissionMiddleware(holder PermissionHolder) *PermissionMiddleware {
	return &PermissionMiddleware{
		holder: holder,
	}
}

// 权限拦截
func (m *PermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("PermissionMiddleware Handle path: %v", r.URL.Path)
		var uid string
		uid = r.Header.Get(restx.HeaderUid)
		// 请求头缺少参数
		if uid == "" {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, "无效请求,缺少用户信息"))
			return
		}

		// 验证用户是否有权限访问资源
		err := m.holder.Enforce(r.Context(), uid, r.URL.Path, r.Method)
		if err != nil {
			responsex.Response(r, w, nil, apierr.NewApiError(apierr.CodeUserNotPermission, err.Error()))
			return
		}

		// 调用下一层的处理
		next.ServeHTTP(w, r)
	}
}
