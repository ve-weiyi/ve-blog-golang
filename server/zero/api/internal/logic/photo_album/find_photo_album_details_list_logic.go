package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPhotoAlbumDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumDetailsListLogic {
	return &FindPhotoAlbumDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumDetailsListLogic) FindPhotoAlbumDetailsList(req *types.PageQuery) (resp []types.PhotoAlbumDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
