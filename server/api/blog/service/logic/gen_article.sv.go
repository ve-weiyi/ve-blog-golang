package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
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

// 删除Article记录
func (s *ArticleService) DeleteArticle(reqCtx *request.Context, article *entity.Article) (rows int64, err error) {
	return s.svcCtx.ArticleRepository.DeleteArticle(reqCtx, article)
}

// 更新Article记录
func (s *ArticleService) UpdateArticle(reqCtx *request.Context, article *entity.Article) (data *entity.Article, err error) {
	return s.svcCtx.ArticleRepository.UpdateArticle(reqCtx, article)
}

// 查询Article记录
func (s *ArticleService) FindArticle(reqCtx *request.Context, article *entity.Article) (data *entity.Article, err error) {
	return s.svcCtx.ArticleRepository.GetArticle(reqCtx, article.ID)
}

// 批量删除Article记录
func (s *ArticleService) DeleteArticleByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.ArticleRepository.DeleteArticleByIds(reqCtx, ids)
}

// 分页获取Article记录
func (s *ArticleService) FindArticleList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Article, total int64, err error) {
	return s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
}
