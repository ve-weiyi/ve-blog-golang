package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type BlogService struct {
	svcCtx *svc.ServiceContext
}

func NewBlogService(svcCtx *svc.ServiceContext) *BlogService {
	return &BlogService{
		svcCtx: svcCtx,
	}
}

func (s *BlogService) GetAdminHomeInfo(reqCtx *request.Context, data interface{}) (resp *response.BlogBackInfoDTO, err error) {
	page := &request.PageQuery{}
	// 查询消息数量
	_, msgCount, err := s.svcCtx.RemarkRepository.FindRemarkList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	// 查询用户数量
	_, userCount, err := s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	// 查询文章数量
	articles, articleCount, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	// 查询分类数量
	categories, _, err := s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	// 查询标签数量
	tags, _, err := s.svcCtx.TagRepository.FindTagList(reqCtx, page)
	if err != nil {
		return nil, err
	}

	uniqueViews, _, err := s.svcCtx.UniqueViewRepository.FindUniqueViewList(reqCtx, page)
	if err != nil {
		return nil, err
	}
	resp = &response.BlogBackInfoDTO{
		ViewsCount:            10,
		MessageCount:          msgCount,
		UserCount:             userCount,
		ArticleCount:          articleCount,
		CategoryDTOList:       convertCategoryList(categories),
		TagDTOList:            convertTagList(tags),
		ArticleStatisticsList: convertArticleStatisticsList(articles),
		UniqueViewDTOList:     convertUniqueViewList(uniqueViews),
		ArticleRankDTOList:    convertArticleRankList(articles),
	}

	return resp, err
}
