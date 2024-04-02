package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoAlbumLogic {
	return &UpdatePhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePhotoAlbumLogic) UpdatePhotoAlbum(req *types.PhotoAlbum) (resp *types.PhotoAlbum, err error) {
	// todo: add your logic here and delete this line

	return
}
