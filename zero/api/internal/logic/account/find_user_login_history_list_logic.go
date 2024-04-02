package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLoginHistoryListLogic {
	return &FindUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLoginHistoryListLogic) FindUserLoginHistoryList(req *types.PageQuery) (resp []types.LoginHistory, err error) {
	// todo: add your logic here and delete this line

	return
}
