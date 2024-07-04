package logrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogLogic {
	return &FindOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询操作记录
func (l *FindOperationLogLogic) FindOperationLog(in *blog.IdReq) (*blog.OperationLog, error) {
	entity, err := l.svcCtx.OperationLogModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertOperationLogModelToPb(entity), nil
}
