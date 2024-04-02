package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteApiLogic) DeleteApi(in *account.IdReq) (*account.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &account.EmptyResp{}, nil
}
