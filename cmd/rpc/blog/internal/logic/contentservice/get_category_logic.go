package contentservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type GetCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询分类详情
func (l *GetCategoryLogic) GetCategory(in *contentrpc.GetCategoryRequest) (*contentrpc.GetCategoryResponse, error) {
	entity, err := l.svcCtx.TCategoryModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &contentrpc.GetCategoryResponse{
		Category: convertCategoryOut(entity),
	}, nil
}
