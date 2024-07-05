package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertUploadPbToModel(in *blog.UploadRecordReq) (out *model.UploadRecord) {
	out = &model.UploadRecord{
		Id:       in.Id,
		UserId:   in.UserId,
		Label:    in.Label,
		FileName: in.FileName,
		FileSize: in.FileSize,
		FileMd5:  in.FileMd5,
		FileUrl:  in.FileUrl,
		//CreatedAt: time.Unix(in.CreatedAt, 0),
		//UpdatedAt: time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func ConvertUploadModelToPb(in *model.UploadRecord) (out *blog.UploadRecordResp) {
	out = &blog.UploadRecordResp{
		Id:        in.Id,
		UserId:    in.UserId,
		Label:     in.Label,
		FileName:  in.FileName,
		FileSize:  in.FileSize,
		FileMd5:   in.FileMd5,
		FileUrl:   in.FileUrl,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
