package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoAlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePhotoAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoAlbumListLogic {
	return &DeletePhotoAlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除相册
func (l *DeletePhotoAlbumListLogic) DeletePhotoAlbumList(in *blog.IdsReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.PhotoAlbumModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
