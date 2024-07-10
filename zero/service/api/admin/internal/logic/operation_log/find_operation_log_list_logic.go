package operation_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取操作记录列表
func NewFindOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogListLogic {
	return &FindOperationLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOperationLogListLogic) FindOperationLogList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.LogRpc.FindOperationLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.LogRpc.FindOperationLogCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.OperationLog
	for _, v := range out.List {
		list = append(list, convert.ConvertOperationLogTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
