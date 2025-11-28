package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/blog/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CommonLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewCommonLogic(svcCtx *svctx.ServiceContext) *CommonLogic {
	return &CommonLogic{
		svcCtx: svcCtx,
	}
}

// ping
func (s *CommonLogic) Ping(reqCtx *request.Context, in *types.PingReq) (out *types.PingResp, err error) {
	// todo

	return
}
