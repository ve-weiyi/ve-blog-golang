package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/messagerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除留言
func NewDeletesRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesRemarkLogic {
	return &DeletesRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesRemarkLogic) DeletesRemark(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &messagerpc.IdsReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.MessageRpc.DeletesRemark(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
