package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type UploadService struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadService(svcCtx *svctx.ServiceContext) *UploadService {
	return &UploadService{
		svcCtx: svcCtx,
	}
}

// 删除文件列表
func (s *UploadService) DeletesUploadFile(reqCtx *request.Context, in *dto.DeletesUploadFileReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 获取文件列表
func (s *UploadService) ListUploadFile(reqCtx *request.Context, in *dto.ListUploadFileReq) (out *dto.PageResp, err error) {
	// todo

	return
}

// 上传文件列表
func (s *UploadService) MultiUploadFile(reqCtx *request.Context, in *dto.MultiUploadFileReq) (out []*dto.FileInfoVO, err error) {
	// todo

	return
}

// 上传文件
func (s *UploadService) UploadFile(reqCtx *request.Context, in *dto.UploadFileReq) (out *dto.FileInfoVO, err error) {
	// todo

	return
}
