package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type UploadLogService struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadLogService(svcCtx *svctx.ServiceContext) *UploadLogService {
	return &UploadLogService{
		svcCtx: svcCtx,
	}
}

// 删除登录日志
func (s *UploadLogService) DeletesUploadLog(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 查询登录日志
func (s *UploadLogService) FindUploadLogList(reqCtx *request.Context, in *dto.UploadLogQuery) (out *dto.PageResp, err error) {
	// todo

	return
}
