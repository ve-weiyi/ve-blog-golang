package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/middleware/visitx"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/syslogservice"
	"github.com/ve-weiyi/blog-cloud/infra/http/headerx"
)

type VisitLogMiddleware struct {
	enforcer visitx.Enforcer
	sr       syslogservice.SyslogService
}

func NewVisitLogMiddleware(enforcer visitx.Enforcer, sr syslogservice.SyslogService) *VisitLogMiddleware {
	return &VisitLogMiddleware{enforcer: enforcer, sr: sr}
}

func (m *VisitLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		pageName, ok := m.enforcer.IsVisitPage(r.URL.Path, r.Method)
		if !ok {
			return
		}

		_, err := m.sr.CreateVisitLog(r.Context(), &syslogservice.CreateVisitLogRequest{
			UserId:   r.Header.Get(headerx.HeaderUid),
			DeviceId: r.Header.Get(headerx.HeaderXDeviceId),
			PageName: pageName,
		})
		if err != nil {
			logx.Errorf("VisitLog CreateVisitLog err: %v", err)
		}
	}
}
