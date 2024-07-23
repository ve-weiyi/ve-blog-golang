package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取相册列表
func NewGetAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlbumListLogic {
	return &GetAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAlbumListLogic) GetAlbumList(req *types.PageQuery) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
