package operation_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOperationLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建操作记录
func NewCreateOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOperationLogLogic {
	return &CreateOperationLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOperationLogLogic) CreateOperationLog(req *types.OperationLog) (resp *types.OperationLog, err error) {
	in := convert.ConvertOperationLogPb(req)
	out, err := l.svcCtx.LogRpc.CreateOperationLog(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convert.ConvertOperationLogTypes(out)
	return resp, nil
}
