package articleservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/model"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建分类
func (l *CreateCategoryLogic) CreateCategory(in *articlerpc.CreateCategoryRequest) (*articlerpc.CreateCategoryResponse, error) {
	entity := &model.TCategory{CategoryName: in.CategoryName}
	_, err := l.svcCtx.TCategoryModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &articlerpc.CreateCategoryResponse{Id: entity.Id}, nil
}
