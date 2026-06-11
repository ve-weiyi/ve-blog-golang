package category

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/articleservice"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建分类
func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryReq) (resp *types.CategoryVO, err error) {
	out, err := l.svcCtx.ArticleService.CreateCategory(l.ctx, &articleservice.CreateCategoryRequest{
		CategoryName: req.CategoryName,
	})
	if err != nil {
		return nil, err
	}

	return &types.CategoryVO{
		Id:           out.Id,
		CategoryName: req.CategoryName,
		ArticleCount: 0,
	}, nil
}
