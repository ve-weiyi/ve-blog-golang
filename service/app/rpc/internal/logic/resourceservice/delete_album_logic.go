package resourceservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type DeleteAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAlbumLogic {
	return &DeleteAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除相册
func (l *DeleteAlbumLogic) DeleteAlbum(in *resourcerpc.DeleteAlbumRequest) (*resourcerpc.DeleteAlbumResponse, error) {
	rows, err := l.svcCtx.TAlbumModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &resourcerpc.DeleteAlbumResponse{SuccessCount: rows}, nil
}
