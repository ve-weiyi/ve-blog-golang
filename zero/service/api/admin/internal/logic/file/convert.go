package file

import (
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/resourcerpc"
)

func ConvertFileFolderPb(in *types.FileFolderNewReq) (out *resourcerpc.FileFolderNewReq) {
	out = &resourcerpc.FileFolderNewReq{
		Id:         in.Id,
		UserId:     0,
		FilePath:   in.FilePath,
		FolderName: in.FolderName,
		FolderDesc: in.FolderDesc,
	}

	return
}

func ConvertFileFolderTypes(in *resourcerpc.FileFolderDetails) (out *types.FileFolderBackDTO) {
	out = &types.FileFolderBackDTO{
		Id:         in.Id,
		UserId:     in.UserId,
		FilePath:   in.FilePath,
		FolderName: in.FolderName,
		FolderDesc: in.FolderDesc,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}

	return
}

func ConvertFileUploadTypes(in *resourcerpc.FileUploadDetails) (out *types.FileUploadBackDTO) {
	out = &types.FileUploadBackDTO{
		Id:        in.Id,
		UserId:    in.UserId,
		FilePath:  in.FilePath,
		FileName:  in.FileName,
		FileType:  in.FileType,
		FileSize:  in.FileSize,
		FileMd5:   in.FileMd5,
		FileUrl:   in.FileUrl,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	return
}
