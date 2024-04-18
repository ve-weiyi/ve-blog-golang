package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiListLogic {
	return &DeleteApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除接口
func (l *DeleteApiListLogic) DeleteApiList(in *account.IdsReq) (*account.BatchResult, error) {
	result, err := l.svcCtx.ApiModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &account.BatchResult{
		SuccessCount: result,
	}, nil
}
