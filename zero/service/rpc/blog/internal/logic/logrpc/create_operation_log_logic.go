package logrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOperationLogLogic {
	return &CreateOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建操作记录
func (l *CreateOperationLogLogic) CreateOperationLog(in *blog.OperationLog) (*blog.OperationLog, error) {
	entity := convert.ConvertOperationLogPbToModel(in)

	_, err := l.svcCtx.OperationLogModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertOperationLogModelToPb(entity), nil
}
