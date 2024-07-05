package operation_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询操作记录
func NewFindOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogLogic {
	return &FindOperationLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOperationLogLogic) FindOperationLog(req *types.IdReq) (resp *types.OperationLog, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.LogRpc.FindOperationLog(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertOperationLogTypes(out), nil
}
