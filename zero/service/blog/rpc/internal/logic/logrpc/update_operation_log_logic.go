package logrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOperationLogLogic {
	return &UpdateOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新操作记录
func (l *UpdateOperationLogLogic) UpdateOperationLog(in *blog.OperationLog) (*blog.OperationLog, error) {
	entity := convert.ConvertOperationLogPbToModel(in)

	result, err := l.svcCtx.OperationLogModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertOperationLogModelToPb(result), nil
}
