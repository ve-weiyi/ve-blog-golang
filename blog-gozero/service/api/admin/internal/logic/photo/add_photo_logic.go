package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建照片
func NewAddPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhotoLogic {
	return &AddPhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPhotoLogic) AddPhoto(req *types.PhotoNewReq) (resp *types.PhotoBackVO, err error) {
	in := ConvertPhotoPb(req)
	out, err := l.svcCtx.ResourceRpc.AddPhoto(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertPhotoTypes(out)
	return resp, nil
}

func ConvertPhotoPb(in *types.PhotoNewReq) (out *resourcerpc.PhotoNewReq) {
	out = &resourcerpc.PhotoNewReq{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
	}

	return
}

func ConvertPhotoTypes(in *resourcerpc.PhotoDetails) (out *types.PhotoBackVO) {
	out = &types.PhotoBackVO{
		Id:        in.Id,
		AlbumId:   in.AlbumId,
		PhotoName: in.PhotoName,
		PhotoDesc: in.PhotoDesc,
		PhotoSrc:  in.PhotoSrc,
		IsDelete:  in.IsDelete,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	return
}
