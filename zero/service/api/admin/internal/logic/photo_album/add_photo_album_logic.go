package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建相册
func NewAddPhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhotoAlbumLogic {
	return &AddPhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPhotoAlbumLogic) AddPhotoAlbum(req *types.PhotoAlbum) (resp *types.PhotoAlbum, err error) {
	// todo: add your logic here and delete this line

	return
}
