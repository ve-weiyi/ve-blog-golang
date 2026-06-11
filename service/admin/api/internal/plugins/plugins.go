package plugins

import (
	"net/http"

	"github.com/ve-weiyi/vkit/plugins/knife4j"
	"github.com/zeromicro/go-zero/rest"

	"github.com/ve-weiyi/ve-blog-golang/infra/staticx"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/docs"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
)

func RegisterPluginHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 注册knife4j服务
	knife4jPrefix := "/admin-api/v1/swagger"
	server.AddRoutes(staticx.PrefixRoutes(knife4jPrefix, http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		knife4j.NewKnife4jPlugin(docs.Docs).Handler(knife4jPrefix).ServeHTTP(w, r)
	}))
}
