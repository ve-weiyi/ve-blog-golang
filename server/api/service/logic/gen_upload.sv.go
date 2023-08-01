package logic

import (
	"mime/multipart"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/crypto"
)

type UploadService struct {
	svcCtx *svc.ServiceContext
}

func NewUploadService(svcCtx *svc.ServiceContext) *UploadService {
	return &UploadService{
		svcCtx: svcCtx,
	}
}

// 上传文件
func (s *UploadService) CreateUpload(reqCtx *request.Context, label string, file *multipart.FileHeader) (data *entity.Upload, err error) {
	s.svcCtx.Log.Println("上传文件")
	url, err := s.svcCtx.Uploader.UploadFile(label, file)
	if err != nil {
		return nil, err
	}

	up := &entity.Upload{
		UserID:   reqCtx.UID,
		Label:    label,
		FileName: file.Filename,
		FileSize: int(file.Size),
		FileMd5:  crypto.MD5V([]byte(file.Filename)),
		FileUrl:  url,
	}

	return s.svcCtx.UploadRepository.CreateUpload(reqCtx, up)
}
