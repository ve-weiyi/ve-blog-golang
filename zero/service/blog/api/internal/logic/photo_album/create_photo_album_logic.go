package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建相册
func NewCreatePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePhotoAlbumLogic {
	return &CreatePhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePhotoAlbumLogic) CreatePhotoAlbum(req *types.PhotoAlbum) (resp *types.PhotoAlbum, err error) {
	in := convert.ConvertPhotoAlbumPb(req)
	out, err := l.svcCtx.PhotoRpc.CreatePhotoAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convert.ConvertPhotoAlbumTypes(out)
	return resp, nil
}
