package logic

import (
	"mime/multipart"
	"path"

	"github.com/spf13/cast"

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
func (s *UploadService) CreateUpload(reqCtx *request.Context, label string, file *multipart.FileHeader) (data *entity.UploadRecord, err error) {
	s.svcCtx.Log.Println("上传文件")
	url, err := s.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(reqCtx.UID), label), file)
	if err != nil {
		return nil, err
	}

	up := &entity.UploadRecord{
		UserID:   reqCtx.UID,
		Label:    label,
		FileName: file.Filename,
		FileSize: int(file.Size),
		FileMd5:  crypto.MD5V([]byte(file.Filename)),
		FileURL:  url,
	}

	return s.svcCtx.UploadRecordRepository.CreateUploadRecord(reqCtx, up)
}
