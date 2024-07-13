package categoryrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

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
func (l *AddCategoryLogic) AddCategory(in *blog.Category) (*blog.Category, error) {
	entity := convert.ConvertCategoryPbToModel(in)

	_, err := l.svcCtx.CategoryModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertCategoryModelToPb(entity), nil
}
