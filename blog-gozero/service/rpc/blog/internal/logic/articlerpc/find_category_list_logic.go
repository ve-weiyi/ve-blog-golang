package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

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
func (l *FindCategoryListLogic) FindCategoryList(in *articlerpc.FindCategoryListReq) (*articlerpc.FindCategoryListResp, error) {
	helper := NewArticleHelperLogic(l.ctx, l.svcCtx)

	opts := []query.Option{
		query.WithPage(int(in.Paginate.Page)),
		query.WithSize(int(in.Paginate.PageSize)),
		query.WithSorts(in.Paginate.Sorts...),
	}
	if in.CategoryName != "" {
		opts = append(opts, query.WithCondition("category_name like ?", "%"+in.CategoryName+"%"))
	}

	page, size, sorts, conditions, params := query.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TCategoryModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	list, err := helper.convertCategoryDetailsResp(records)
	if err != nil {
		return nil, err
	}

	return &articlerpc.FindCategoryListResp{
		List: list,
		Pagination: &articlerpc.PageResp{
			Page:     in.Paginate.Page,
			PageSize: in.Paginate.PageSize,
			Total:    total,
		},
	}, nil
}
