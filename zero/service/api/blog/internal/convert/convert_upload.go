package convert

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
)

func ConvertUploadTypes(in *blog.UploadRecordResp) (out *types.UploadFileResp) {

	out = &types.UploadFileResp{
		Id:        in.Id,
		UserId:    in.UserId,
		Label:     in.Label,
		FileName:  in.FileName,
		FileSize:  in.FileSize,
		FileMd5:   in.FileMd5,
		FileUrl:   in.FileUrl,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	return
}
