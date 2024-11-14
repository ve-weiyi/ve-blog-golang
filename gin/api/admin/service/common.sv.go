package service

import (
	"github.com/ve-weiyi/ve-blog-golang/gin/api/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

type CommonService struct {
	svcCtx *svctx.ServiceContext
}

func NewCommonService(svcCtx *svctx.ServiceContext) *CommonService {
	return &CommonService{
		svcCtx: svcCtx,
	}
}

// ping
func (s *CommonService) Ping(reqCtx *request.Context, in *dto.PingReq) (out *dto.PingResp, err error) {
	// todo

	return
}
