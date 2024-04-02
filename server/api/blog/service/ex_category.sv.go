package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
)

// 分页获取Category记录
func (l *CategoryService) FindCategoryDetailsList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.CategoryDetailsDTO, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	categories, err := l.svcCtx.CategoryRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.CategoryRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// 查询分类下的文章数量
	for _, in := range categories {

		articles, err := l.svcCtx.ArticleRepository.FindArticleListByCategoryId(reqCtx, in.Id)
		if err != nil {
			return nil, 0, err
		}

		out := &dto.CategoryDetailsDTO{
			Id:           in.Id,
			CategoryName: in.CategoryName,
			ArticleCount: int64(len(articles)),
			CreatedAt:    in.CreatedAt,
			UpdatedAt:    in.UpdatedAt,
		}

		list = append(list, out)
	}

	return list, total, err
}
