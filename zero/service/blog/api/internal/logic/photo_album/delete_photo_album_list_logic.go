package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除相册
func NewDeletePhotoAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoAlbumListLogic {
	return &DeletePhotoAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePhotoAlbumListLogic) DeletePhotoAlbumList(reqCtx *types.RestHeader, req *types.IdsReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
