package logrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOperationLogCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogCountLogic {
	return &FindOperationLogCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询操作记录数量
func (l *FindOperationLogCountLogic) FindOperationLogCount(in *blog.PageQuery) (*blog.CountResp, error) {
	_, _, _, conditions, params := convert.ParsePageQuery(in)

	count, err := l.svcCtx.OperationLogModel.FindCount(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	return &blog.CountResp{
		Count: count,
	}, nil
}
