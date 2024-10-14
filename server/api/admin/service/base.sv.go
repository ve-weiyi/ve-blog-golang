package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type BaseService struct {
	svcCtx *svctx.ServiceContext
}

func NewBaseService(svcCtx *svctx.ServiceContext) *BaseService {
	return &BaseService{
		svcCtx: svcCtx,
	}
}

// ping
func (s *BaseService) Ping(reqCtx *request.Context, in *dto.PingReq) (out *dto.PingResp, err error) {
	// todo

	return
}
