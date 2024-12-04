package logic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/websiterpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 身份标识
func NewIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentityLogic {
	return &IdentityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IdentityLogic) Identity(req *types.EmptyReq) (resp *types.IdentityResp, err error) {
	id, err := l.svcCtx.WebsiteRpc.GetIdentity(l.ctx, &websiterpc.EmptyReq{})
	if err != nil {
		return nil, err
	}

	return &types.IdentityResp{
		TerminalId: id.TerminalId,
	}, nil
}
