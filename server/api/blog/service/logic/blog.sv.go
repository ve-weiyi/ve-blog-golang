package logic

import "github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"

type BlogService struct {
	svcCtx *svc.ServiceContext
}

func NewBlogService(svcCtx *svc.ServiceContext) *BlogService {
	return &BlogService{
		svcCtx: svcCtx,
	}
}
