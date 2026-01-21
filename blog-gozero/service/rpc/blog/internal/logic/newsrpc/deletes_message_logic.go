package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesMessageLogic {
	return &DeletesMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除留言
func (l *DeletesMessageLogic) DeletesMessage(in *newsrpc.DeletesMessageReq) (*newsrpc.DeletesMessageResp, error) {
	rows, err := l.svcCtx.TMessageModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &newsrpc.DeletesMessageResp{
		SuccessCount: rows,
	}, nil
}
