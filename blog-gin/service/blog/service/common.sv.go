package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
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
