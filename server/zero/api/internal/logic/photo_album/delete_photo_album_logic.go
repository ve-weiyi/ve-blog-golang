package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoAlbumLogic {
	return &DeletePhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePhotoAlbumLogic) DeletePhotoAlbum(req *types.IdReq) error {
	// todo: add your logic here and delete this line

	return nil
}
