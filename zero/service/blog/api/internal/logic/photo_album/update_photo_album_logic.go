package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新相册
func NewUpdatePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoAlbumLogic {
	return &UpdatePhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePhotoAlbumLogic) UpdatePhotoAlbum(reqCtx *types.RestHeader, req *types.PhotoAlbum) (resp *types.PhotoAlbum, err error) {
	in := convert.ConvertPhotoAlbumPb(req)

	api, err := l.svcCtx.PhotoRpc.UpdatePhotoAlbum(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoAlbumTypes(api), nil
}
