package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 分页获取Tag记录
func (l *TagService) FindTagDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.TagDetailsDTO, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	categories, err := l.svcCtx.TagRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.TagRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 查询分类下的文章数量

	for _, in := range categories {

		articles, err := l.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, in.ID)
		if err != nil {
			return nil, 0, err
		}

		out := &response.TagDetailsDTO{
			ID:           in.ID,
			TagName:      in.TagName,
			ArticleCount: int64(len(articles)),
			CreatedAt:    in.CreatedAt,
			UpdatedAt:    in.UpdatedAt,
		}

		list = append(list, out)
	}

	return list, total, err
}
