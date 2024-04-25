package logrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogListLogic {
	return &FindOperationLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取操作记录列表
func (l *FindOperationLogListLogic) FindOperationLogList(in *blog.PageQuery) (*blog.OperationLogPageResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.OperationLogModel.FindList(l.ctx, limit, offset, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	var list []*blog.OperationLog
	for _, v := range result {
		list = append(list, convert.ConvertOperationLogModelToPb(v))
	}

	return &blog.OperationLogPageResp{
		List: list,
	}, nil
}
