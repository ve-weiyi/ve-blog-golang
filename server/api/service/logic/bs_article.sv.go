package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type ArticleService struct {
	svcCtx *svc.ServiceContext
}

func NewArticleService(svcCtx *svc.ServiceContext) *ArticleService {
	return &ArticleService{
		svcCtx: svcCtx,
	}
}

// 创建Article记录
func (s *ArticleService) CreateArticle(reqCtx *request.Context, article *entity.Article) (data *entity.Article, err error) {
	return s.svcCtx.ArticleRepository.CreateArticle(reqCtx, article)
}

// 更新Article记录
func (s *ArticleService) UpdateArticle(reqCtx *request.Context, article *entity.Article) (data *entity.Article, err error) {
	return s.svcCtx.ArticleRepository.UpdateArticle(reqCtx, article)
}

// 删除Article记录
func (s *ArticleService) DeleteArticle(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.ArticleRepository.DeleteArticleById(reqCtx, id)
}

// 查询Article记录
func (s *ArticleService) FindArticle(reqCtx *request.Context, id int) (data *entity.Article, err error) {
	return s.svcCtx.ArticleRepository.FindArticleById(reqCtx, id)
}

// 批量删除Article记录
func (s *ArticleService) DeleteArticleByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.ArticleRepository.DeleteArticleByIds(reqCtx, ids)
}

// 分页获取Article记录
func (s *ArticleService) FindArticleList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Article, total int64, err error) {
	list, err = s.svcCtx.ArticleRepository.FindArticleList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
