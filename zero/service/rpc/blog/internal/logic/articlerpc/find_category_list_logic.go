package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCategoryListLogic {
	return &FindCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章分类数量
func (l *FindCategoryListLogic) FindCategoryList(in *blog.FindCategoryListReq) (*blog.FindCategoryListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)
	sorts = in.Sorts
	if in.CategoryName != "" {
		conditions += "category_name like ?"
		params = append(params, "%"+in.CategoryName+"%")
	}

	records, err := l.svcCtx.CategoryModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.CategoryModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	acm, err := findArticleCountGroupCategory(l.ctx, l.svcCtx, records)
	if err != nil {
		return nil, err
	}

	var list []*blog.CategoryDetails
	for _, v := range records {
		list = append(list, convertCategoryOut(v, acm))
	}

	return &blog.FindCategoryListResp{
		List:  list,
		Total: count,
	}, nil
}
