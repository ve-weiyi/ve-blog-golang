package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPhotoAlbumDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumDetailsLogic {
	return &FindPhotoAlbumDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumDetailsLogic) FindPhotoAlbumDetails(req *types.IdReq) (resp *types.PhotoAlbumDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
