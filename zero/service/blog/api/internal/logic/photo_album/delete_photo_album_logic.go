package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除相册
func NewDeletePhotoAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoAlbumLogic {
	return &DeletePhotoAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePhotoAlbumLogic) DeletePhotoAlbum(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
