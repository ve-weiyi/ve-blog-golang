package category

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文章分类
func NewFindCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCategoryLogic {
	return &FindCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCategoryLogic) FindCategory(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.Category, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.CategoryRpc.FindCategory(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertCategoryTypes(out), nil
}
