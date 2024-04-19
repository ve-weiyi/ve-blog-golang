package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanApiListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanApiListLogic {
	return &CleanApiListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清空接口列表
func (l *CleanApiListLogic) CleanApiList(in *account.EmptyReq) (*account.BatchResult, error) {
	row, err := l.svcCtx.RoleModel.DeleteBatch(l.ctx, "1 = 1")
	if err != nil {
		return nil, err
	}

	return &account.BatchResult{
		SuccessCount: row,
	}, nil
}
