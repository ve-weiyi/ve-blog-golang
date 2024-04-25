package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询相册
func NewFindPhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumLogic {
	return &FindPhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumLogic) FindPhotoAlbum(req *types.IdReq) (resp *types.PhotoAlbum, err error) {
	// todo: add your logic here and delete this line

	return
}
