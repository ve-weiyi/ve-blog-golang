package logic

import "github.com/ve-weiyi/ve-admin-store/server/api/blog/router/svc"

type BlogRouter struct {
	svcCtx *svc.RouterContext
}

func NewBlogRouter(ctx *svc.RouterContext) *BlogRouter {
	return &BlogRouter{
		svcCtx: ctx,
	}
}
