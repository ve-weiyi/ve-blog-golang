package category

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"
)

type FindCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文章分类列表
func NewFindCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCategoryListLogic {
	return &FindCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCategoryListLogic) FindCategoryList(req *types.FindCategoryListReq) (resp *types.PageResp, err error) {
	in := &articlerpc.FindCategoryListReq{
		Query: &articlerpc.PageLimit{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		CategoryName: req.CategoryName,
	}

	out, err := l.svcCtx.ArticleRpc.FindCategoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.CategoryBackDTO
	for _, item := range out.List {
		list = append(list, ConvertCategoryTypes(item))
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = out.Total
	resp.List = list
	return
}

func ConvertCategoryTypes(in *articlerpc.CategoryDetails) (out *types.CategoryBackDTO) {
	return &types.CategoryBackDTO{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		ArticleCount: in.ArticleCount,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
