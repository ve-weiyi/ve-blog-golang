package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
)

// 分页获取Tag记录
func (l *TagService) FindTagDetailsList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.TagDetailsDTO, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	categories, err := l.svcCtx.TagRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.TagRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 查询分类下的文章数量

	for _, in := range categories {

		articles, err := l.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, in.Id)
		if err != nil {
			return nil, 0, err
		}

		out := &dto.TagDetailsDTO{
			Id:           in.Id,
			TagName:      in.TagName,
			ArticleCount: int64(len(articles)),
			CreatedAt:    in.CreatedAt,
			UpdatedAt:    in.UpdatedAt,
		}

		list = append(list, out)
	}

	return list, total, err
}
