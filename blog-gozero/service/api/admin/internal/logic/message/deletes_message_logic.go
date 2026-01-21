package message

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除留言
func NewDeletesMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesMessageLogic {
	return &DeletesMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesMessageLogic) DeletesMessage(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &newsrpc.DeletesMessageReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.NewsRpc.DeletesMessage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
