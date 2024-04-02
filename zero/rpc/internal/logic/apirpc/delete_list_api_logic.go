package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteListApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteListApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteListApiLogic {
	return &DeleteListApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteListApiLogic) DeleteListApi(in *account.IdsReq) (*account.BatchResult, error) {
	// todo: add your logic here and delete this line

	return &account.BatchResult{}, nil
}
