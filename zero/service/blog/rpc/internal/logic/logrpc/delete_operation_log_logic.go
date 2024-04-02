package logrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOperationLogLogic {
	return &DeleteOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除操作记录
func (l *DeleteOperationLogLogic) DeleteOperationLog(in *blog.IdReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.OperationLogModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
