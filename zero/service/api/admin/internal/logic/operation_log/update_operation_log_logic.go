package operation_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOperationLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新操作记录
func NewUpdateOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOperationLogLogic {
	return &UpdateOperationLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOperationLogLogic) UpdateOperationLog(req *types.OperationLog) (resp *types.OperationLog, err error) {
	in := ConvertOperationLogPb(req)

	api, err := l.svcCtx.LogRpc.UpdateOperationLog(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return ConvertOperationLogTypes(api), nil
}
