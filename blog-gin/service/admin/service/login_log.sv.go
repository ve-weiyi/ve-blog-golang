package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type LoginLogService struct {
	svcCtx *svctx.ServiceContext
}

func NewLoginLogService(svcCtx *svctx.ServiceContext) *LoginLogService {
	return &LoginLogService{
		svcCtx: svcCtx,
	}
}

// 删除登录日志
func (s *LoginLogService) DeletesLoginLog(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 查询登录日志
func (s *LoginLogService) FindLoginLogList(reqCtx *request.Context, in *dto.LoginLogQuery) (out *dto.PageResp, err error) {
	// todo

	return
}
