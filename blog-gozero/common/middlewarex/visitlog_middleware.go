package middlewarex

import (
	"context"
	"net/http"

	"github.com/go-openapi/spec"
	"github.com/mssola/useragent"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
)

type VisitLogMiddleware struct {
	sp *spec.Swagger

	sr syslogrpc.SyslogRpc
}

func NewVisitLogMiddleware(sp *spec.Swagger, sr syslogrpc.SyslogRpc) *VisitLogMiddleware {
	return &VisitLogMiddleware{
		sp: sp,
		sr: sr,
	}
}

// 记录操作记录
func (m *VisitLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("VisitLogMiddleware Handle path: %v", r.URL.Path)

		// 调用下一层的处理
		next.ServeHTTP(w, r)

		ip := restx.GetClientIP(r)
		is, err := ipx.GetIpSourceByBaidu(ip)
		if err != nil {
			logx.Errorf("VisitLogMiddleware Handle GetIpInfoByBaidu err: %v", err)
		}

		// 分割字符串，提取 IP 部分
		os := useragent.New(r.UserAgent()).OS()
		browser, _ := useragent.New(r.UserAgent()).Browser()

		op := &syslogrpc.VisitLogNewReq{
			UserId:     r.Header.Get(restx.HeaderUid),
			TerminalId: r.Header.Get(restx.HeaderTerminal),
			IpAddress:  ip,
			IpSource:   is,
			Os:         os,
			Browser:    browser,
			Page:       module,
		}

		_, err = m.sr.AddVisitLog(context.Background(), op)
		if err != nil {
			logx.Errorf("VisitLogMiddleware Handle AddVisitLog err: %v", err)
		}
	}
}

func (m *VisitLogMiddleware) getApi(path string, method string) *spec.Operation {
	sp := m.sp
	for k, v := range sp.Paths.Paths {
		if k == path {
			switch method {
			case http.MethodGet:
				return v.Get
			case http.MethodPost:
				return v.Post
			case http.MethodPut:
				return v.Put
			case http.MethodDelete:
				return v.Delete
			case http.MethodPatch:
				return v.Patch
			case http.MethodOptions:
				return v.Options
			case http.MethodHead:
				return v.Head
			}
		}
	}

	return nil
}
