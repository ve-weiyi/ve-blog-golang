package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type BlogService struct {
	svcCtx *svc.ServiceContext
}

func NewBlogService(svcCtx *svc.ServiceContext) *BlogService {
	return &BlogService{
		svcCtx: svcCtx,
	}
}

func (s *BlogService) FindChatRecordList(reqCtx *request.Context, page *request.PageInfo) (list []*response.MenuTree, total int64, err error) {

	return list, int64(len(list)), nil
}
