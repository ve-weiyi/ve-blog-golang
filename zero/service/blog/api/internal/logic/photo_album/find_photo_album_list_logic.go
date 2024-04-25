package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取相册列表
func NewFindPhotoAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumListLogic {
	return &FindPhotoAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumListLogic) FindPhotoAlbumList(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
