package pagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPageCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageCountLogic {
	return &FindPageCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询页面数量
func (l *FindPageCountLogic) FindPageCount(in *blog.PageQuery) (*blog.CountResp, error) {
	_, _, _, conditions, params := convert.ParsePageQuery(in)

	count, err := l.svcCtx.PageModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	return &blog.CountResp{
		Count: count,
	}, nil
}
