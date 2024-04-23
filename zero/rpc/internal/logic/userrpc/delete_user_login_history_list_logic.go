package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLoginHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLoginHistoryListLogic {
	return &DeleteUserLoginHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除登录历史
func (l *DeleteUserLoginHistoryListLogic) DeleteUserLoginHistoryList(in *account.IdsReq) (*account.BatchResp, error) {
	result, err := l.svcCtx.UserLoginHistoryModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &account.BatchResp{
		SuccessCount: result,
	}, nil
}
