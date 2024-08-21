package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCategoryLogic {
	return &AddCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建文章分类
func (l *AddCategoryLogic) AddCategory(in *blog.CategoryNew) (*blog.CategoryDetails, error) {
	entity := convertCategoryIn(in)
	_, err := l.svcCtx.CategoryModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &blog.CategoryDetails{}, nil
}