package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
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

// 创建文件目录
func (s *FileService) AddFileFolder(reqCtx *request.Context, in *dto.FileFolderNewReq) (out *dto.FileBackVO, err error) {
	// todo

	return
}

// 删除文件列表
func (s *FileService) DeletesFile(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取文件列表
func (s *FileService) FindFileList(reqCtx *request.Context, in *dto.FileQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 获取文件列表
func (s *FileService) ListUploadFile(reqCtx *request.Context, in *dto.ListUploadFileReq) (out *dto.ListUploadFileResp, err error) {
	// todo

	return
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
