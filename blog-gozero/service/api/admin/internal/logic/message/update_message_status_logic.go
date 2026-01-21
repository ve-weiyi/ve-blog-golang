package message

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMessageStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新留言
func NewUpdateMessageStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageStatusLogic {
	return &UpdateMessageStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMessageStatusLogic) UpdateMessageStatus(req *types.UpdateMessageStatusReq) (resp *types.BatchResp, err error) {
	in := &newsrpc.UpdateMessageStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	}

	out, err := l.svcCtx.NewsRpc.UpdateMessageStatus(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
