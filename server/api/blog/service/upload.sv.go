package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type UploadService struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadService(svcCtx *svctx.ServiceContext) *UploadService {
	return &UploadService{
		svcCtx: svcCtx,
	}
}

// 上传文件
func (s *UploadService) UploadFile(reqCtx *request.Context, in *dto.UploadFileReq) (out *dto.UploadFileResp, err error) {
	// todo

	return
}
