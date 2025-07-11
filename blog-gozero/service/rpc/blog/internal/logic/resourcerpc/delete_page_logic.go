package resourcerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePageLogic {
	return &DeletePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除页面
func (l *DeletePageLogic) DeletePage(in *resourcerpc.IdsReq) (*resourcerpc.BatchResp, error) {
	rows, err := l.svcCtx.TPageModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &resourcerpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
