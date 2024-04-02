package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPhotoAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumListLogic {
	return &FindPhotoAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumListLogic) FindPhotoAlbumList(req *types.PageQuery) (resp []types.PhotoAlbum, err error) {
	// todo: add your logic here and delete this line

	return
}
