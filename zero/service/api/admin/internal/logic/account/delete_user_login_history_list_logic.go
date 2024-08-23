package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLoginHistoryListLogic {
	return &DeleteUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLoginHistoryListLogic) DeleteUserLoginHistoryList(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &accountrpc.IdsReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.AccountRpc.DeleteUserLoginHistoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
