package operation_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOperationLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOperationLogLogic {
	return &UpdateOperationLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOperationLogLogic) UpdateOperationLog(req *types.OperationLog) (resp *types.OperationLog, err error) {
	// todo: add your logic here and delete this line

	return
}
