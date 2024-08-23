package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteUserLoginHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除登录历史
func NewBatchDeleteUserLoginHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteUserLoginHistoryLogic {
	return &BatchDeleteUserLoginHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteUserLoginHistoryLogic) BatchDeleteUserLoginHistory(req *types.IdsReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
