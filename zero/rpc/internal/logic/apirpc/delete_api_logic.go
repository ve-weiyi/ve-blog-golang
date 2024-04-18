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

// 删除接口
func (l *DeleteApiLogic) DeleteApi(in *account.IdReq) (*account.BatchResult, error) {
	result, err := l.svcCtx.ApiModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	result2, err := l.svcCtx.ApiModel.BatchDelete(l.ctx, "parent_id = ? ", in.Id)
	if err != nil {
		return nil, err
	}

	return &account.BatchResult{
		SuccessCount: result + result2,
	}, nil
}
