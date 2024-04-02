package operation_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogLogic {
	return &FindOperationLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOperationLogLogic) FindOperationLog(req *types.IdReq) (resp *types.OperationLog, err error) {
	// todo: add your logic here and delete this line

	return
}
