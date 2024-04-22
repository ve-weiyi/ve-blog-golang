package user

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

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

func (l *DeleteUserLoginHistoryListLogic) DeleteUserLoginHistoryList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.UserRpc.DeleteUserLoginHistoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResult{
		SuccessCount: out.SuccessCount,
	}, nil
}
