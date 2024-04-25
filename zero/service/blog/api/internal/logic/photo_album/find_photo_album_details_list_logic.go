package photo_album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoAlbumDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取相册详情列表
func NewFindPhotoAlbumDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoAlbumDetailsListLogic {
	return &FindPhotoAlbumDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoAlbumDetailsListLogic) FindPhotoAlbumDetailsList(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
