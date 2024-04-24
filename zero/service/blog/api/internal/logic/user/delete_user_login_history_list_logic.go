package user

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

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
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.UserRpc.DeleteUserLoginHistoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
