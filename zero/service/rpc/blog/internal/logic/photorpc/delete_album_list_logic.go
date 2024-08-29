package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAlbumListLogic {
	return &DeleteAlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除相册
func (l *DeleteAlbumListLogic) DeleteAlbumList(in *photorpc.IdsReq) (*photorpc.BatchResp, error) {
	rows, err := l.svcCtx.AlbumModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &photorpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
