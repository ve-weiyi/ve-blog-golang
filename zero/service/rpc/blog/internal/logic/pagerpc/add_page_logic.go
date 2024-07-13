package pagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPageLogic {
	return &AddPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建页面
func (l *AddPageLogic) AddPage(in *blog.Page) (*blog.Page, error) {
	entity := convert.ConvertPagePbToModel(in)

	_, err := l.svcCtx.PageModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPageModelToPb(entity), nil
}
