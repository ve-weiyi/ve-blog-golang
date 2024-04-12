package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLoginHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLoginHistoryListLogic {
	return &FindUserLoginHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户登录历史
func (l *FindUserLoginHistoryListLogic) FindUserLoginHistoryList(in *account.PageQuery) (*account.LoginHistoryPageResp, error) {
	// todo: add your logic here and delete this line

	return &account.LoginHistoryPageResp{}, nil
}
