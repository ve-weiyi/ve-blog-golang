package pagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePageListLogic {
	return &DeletePageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除页面
func (l *DeletePageListLogic) DeletePageList(in *blog.IdsReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.PageModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
