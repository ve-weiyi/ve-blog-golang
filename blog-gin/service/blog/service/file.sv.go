package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type FileService struct {
	svcCtx *svctx.ServiceContext
}

func NewFileService(svcCtx *svctx.ServiceContext) *FileService {
	return &FileService{
		svcCtx: svcCtx,
	}
}

// 上传文件列表
func (s *FileService) MultiUploadFile(reqCtx *request.Context, in *dto.MultiUploadFileReq) (out []*dto.FileBackVO, err error) {
	// todo

	return
}

// 上传文件
func (s *FileService) UploadFile(reqCtx *request.Context, in *dto.UploadFileReq) (out *dto.FileBackVO, err error) {
	// todo

	return
}
