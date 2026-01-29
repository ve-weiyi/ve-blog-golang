package articlerpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/infra/queryx"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
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

	var opts []queryx.Option
	if in.Paginate != nil {
		opts = append(opts, queryx.WithPage(int(in.Paginate.Page)))
		opts = append(opts, queryx.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, queryx.WithSorts(in.Paginate.Sorts...))
	}
	if in.CategoryName != "" {
		opts = append(opts, queryx.WithCondition("category_name like ?", "%"+in.CategoryName+"%"))
	}

	page, size, sorts, conditions, params := queryx.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TCategoryModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	list, err := helper.convertCategory(records)
	if err != nil {
		return nil, err
	}

	return &articlerpc.FindCategoryListResp{
		List: list,
		Pagination: &articlerpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}
