package pagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageListLogic {
	return &FindPageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取页面列表
func (l *FindPageListLogic) FindPageList(in *blog.PageQuery) (*blog.PagePageResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.PageModel.FindList(l.ctx, limit, offset, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.Page
	for _, v := range result {
		list = append(list, convert.ConvertPageModelToPb(v))
	}

	return &blog.PagePageResp{
		List: list,
	}, nil
}
