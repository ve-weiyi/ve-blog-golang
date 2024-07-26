package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除页面
func NewDeletePageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePageListLogic {
	return &DeletePageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePageListLogic) DeletePageList(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.PageRpc.DeletePageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}