package logic

import (
	"github.com/ve-weiyi/go-sdk/utils/crypto"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
	"mime/multipart"
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
