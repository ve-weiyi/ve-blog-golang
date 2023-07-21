package tmpl

const AppRouter = `
package router

import (

)

type AppRouter struct {
	svcCtx *svc.RouterContext //持有的controller引用
}

func NewRouter(svcCtx *svc.RouterContext) *AppRouter {
	return &AppRouter{
		svcCtx: svcCtx,
	}
}
`
const RouterContext = `
package svc

import (

)

// 注册需要用到的api
type RouterContext struct {
	*controller.AppController
}

func NewRouterContext(cfg *config.Config) *RouterContext {
	ctx := svc.NewControllerContext(cfg)
	ctl := controller.NewController(ctx)
	if ctl == nil {
		panic("ctl cannot be null")
	}

	return &RouterContext{
		AppController: ctl,
	}
}

`

const Router = `
package logic

import (
	"github.com/gin-gonic/gin"

	{{range .ImportPkgPaths}}{{.}} ` + "\n" + `{{end}}
)

type {{.StructName}}Router struct {
	svcCtx *svc.RouterContext
}

func New{{.StructName}}Router(svcCtx *svc.RouterContext) *{{.StructName}}Router {
	return &{{.StructName}}Router{
		svcCtx: svcCtx,
	}
}

// 初始化 {{.StructName}} 路由信息
// publicRouter 公开路由，不登录就可以访问
// loginRouter  登录路由，登录后才可以访问
func (s *{{.StructName}}Router) Init{{.StructName}}GenRouter(publicRouter *gin.RouterGroup, loginRouter *gin.RouterGroup) {

	var handler = s.svcCtx.AppController.{{.StructName}}Controller
	{
		publicRouter.POST("{{.JsonName}}", handler.Create{{.StructName}})             // 新建{{.StructName}}
		publicRouter.PUT("{{.JsonName}}", handler.Update{{.StructName}})              // 更新{{.StructName}}
		publicRouter.DELETE("{{.JsonName}}/:id", handler.Delete{{.StructName}})           // 删除{{.StructName}}
		publicRouter.GET("{{.JsonName}}/:id", handler.Find{{.StructName}})				  // 查询{{.StructName}}

		publicRouter.DELETE("{{.JsonName}}/batch_delete", handler.Delete{{.StructName}}ByIds)	// 批量删除{{.StructName}}列表
		publicRouter.POST("{{.JsonName}}/list", handler.Find{{.StructName}}List)  				// 分页查询{{.StructName}}列表
	}
}
`

const CommonRouter = `
package logic

import (

)

type {{.StructName}}Router struct {
	svcCtx *svc.RouterContext
}

func New{{.StructName}}Router(svcCtx *svc.RouterContext) *{{.StructName}}Router {
	return &{{.StructName}}Router{
		svcCtx: svcCtx,
	}
}
`
