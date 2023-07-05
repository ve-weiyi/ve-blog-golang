package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type ArticleTagService struct {
	svcCtx *svc.ServiceContext
}

func NewArticleTagService(svcCtx *svc.ServiceContext) *ArticleTagService {
	return &ArticleTagService{
		svcCtx: svcCtx,
	}
}

// 创建ArticleTag记录
func (s *ArticleTagService) CreateArticleTag(reqCtx *request.Context, articleTag *entity.ArticleTag) (data *entity.ArticleTag, err error) {
	return s.svcCtx.ArticleTagRepository.CreateArticleTag(articleTag)
}

// 删除ArticleTag记录
func (s *ArticleTagService) DeleteArticleTag(reqCtx *request.Context, articleTag *entity.ArticleTag) (rows int64, err error) {
	return s.svcCtx.ArticleTagRepository.DeleteArticleTag(articleTag)
}

// 更新ArticleTag记录
func (s *ArticleTagService) UpdateArticleTag(reqCtx *request.Context, articleTag *entity.ArticleTag) (data *entity.ArticleTag, err error) {
	return s.svcCtx.ArticleTagRepository.UpdateArticleTag(articleTag)
}

// 根据id获取ArticleTag记录
func (s *ArticleTagService) FindArticleTag(reqCtx *request.Context, id int) (data *entity.ArticleTag, err error) {
	return s.svcCtx.ArticleTagRepository.FindArticleTag(id)
}

// 批量删除ArticleTag记录
func (s *ArticleTagService) DeleteArticleTagByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.ArticleTagRepository.DeleteArticleTagByIds(ids)
}

// 分页获取ArticleTag记录
func (s *ArticleTagService) GetArticleTagList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.ArticleTag, total int64, err error) {
	return s.svcCtx.ArticleTagRepository.GetArticleTagList(page)
}
