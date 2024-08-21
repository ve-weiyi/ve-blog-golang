package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文章分类
func (l *UpdateCategoryLogic) UpdateCategory(in *blog.CategoryNew) (*blog.CategoryDetails, error) {
	entity := convertCategoryIn(in)
	_, err := l.svcCtx.CategoryModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &blog.CategoryDetails{}, nil
}