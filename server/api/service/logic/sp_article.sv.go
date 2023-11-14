package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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
func (s *ArticleService) SaveArticle(reqCtx *request.Context, req *request.ArticleDetailsReq) (data *entity.Article, err error) {
	// 创建文章
	article := &entity.Article{
		ID:             req.ID,
		UserID:         reqCtx.UID,
		ArticleCover:   req.ArticleCover,
		ArticleTitle:   req.ArticleTitle,
		ArticleContent: req.ArticleContent,
		Type:           req.Type,
		OriginalURL:    req.OriginalURL,
		IsTop:          req.IsTop,
		IsDelete:       0,
		Status:         req.Status,
	}

	// 设置默认文章封面
	if article.ArticleCover == "" {
		article.ArticleCover = "https://ve-blog.oss-cn-shenzhen.aliyuncs.com/ve-blog/2021/08/31/163040.jpg"
	}

	// 查找分类是否存在
	category, _ := s.svcCtx.CategoryRepository.CreateCategoryNotExist(reqCtx, req.CategoryName)
	if category != nil {
		article.CategoryID = category.ID
	}

	// 创建文章
	_, err = s.svcCtx.ArticleRepository.CreateArticle(reqCtx, article)
	if err != nil {
		return nil, err
	}

	// 删除文章标签映射
	_, _ = s.svcCtx.ArticleTagRepository.DeleteArticleTag(reqCtx, sqlx.NewCondition("article_id = ?", article.ID))
	// 创建不存在的标签
	tags, _ := s.svcCtx.TagRepository.BatchCreateTagNotExist(reqCtx, req.TagNameList)
	// 创建文章标签映射
	for _, tag := range tags {
		at := &entity.ArticleTag{
			ArticleID: article.ID,
			TagID:     tag.ID,
		}
		_, _ = s.svcCtx.ArticleTagRepository.CreateArticleTag(reqCtx, at)
	}

	return article, nil
}

// 删除Article记录
func (s *ArticleService) DeleteArticle(reqCtx *request.Context, id int) (rows int, err error) {
	// 删除文章标签映射
	_, err = s.svcCtx.ArticleTagRepository.DeleteArticleTag(reqCtx, sqlx.NewCondition("article_id = ?", id))
	if err != nil {
		return 0, err
	}

	return s.svcCtx.ArticleRepository.DeleteArticleById(reqCtx, id)
}

// 根据id获取Article记录
func (s *ArticleService) FindArticle(reqCtx *request.Context, id int) (data *response.ArticleBack, err error) {
	// 查询id对应文章
	article, err := s.svcCtx.ArticleRepository.FindArticleById(reqCtx, id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, _ := s.svcCtx.CategoryRepository.FindCategoryById(reqCtx, article.CategoryID)

	// 查询文章标签
	tags, _ := s.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

	resp := &response.ArticleBack{}
	resp.ArticleDTO = convertArticle(article)
	resp.CategoryName = getCategoryName(category)
	resp.TagNameList = getTagNameList(tags)
	return resp, nil
}

// 分页获取Article记录
func (s *ArticleService) FindArticleList(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticleBack, total int64, err error) {
	// 查询文章列表
	articles, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}

	// id
	var articleIds []int
	var categoryIds []int
	for _, article := range articles {
		articleIds = append(articleIds, article.ID)
		categoryIds = append(categoryIds, article.CategoryID)
	}

	// 查询所有文章分类
	category, _ := s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, nil, nil, sqlx.NewCondition("id in ?", categoryIds))
	var cmp = make(map[int]*entity.Category)
	for _, item := range category {
		cmp[item.ID] = item
	}

	// 查询所有文章标签映射
	amp, _ := s.svcCtx.TagRepository.FindArticleTagMap(reqCtx, articleIds)

	for _, article := range articles {

		articleVO := &response.ArticleBack{}
		articleVO.ArticleDTO = convertArticle(article)
		articleVO.CategoryName = getCategoryName(cmp[article.CategoryID])
		articleVO.TagNameList = getTagNameList(amp[article.ID])
		list = append(list, articleVO)
	}
	return list, total, err
}

// 更新Article删除状态
func (s *ArticleService) UpdateArticleDelete(reqCtx *request.Context, req *request.ArticleDeleteReq) (rows int, err error) {
	return s.svcCtx.ArticleRepository.UpdateArticleDelete(reqCtx, req.ID, req.IsDelete)
}

// 更新Article记录
func (s *ArticleService) UpdateArticleTop(reqCtx *request.Context, req *request.ArticleTopReq) (rows int, err error) {
	return s.svcCtx.ArticleRepository.UpdateArticleTop(reqCtx, req.ID, req.IsTop)
}

// 文章归类
func (s *ArticleService) FindArticleSeries(reqCtx *request.Context, req *request.ArticleCondition) (data *response.ArticleConditionDTO, err error) {
	data = &response.ArticleConditionDTO{}
	// 查询文章列表
	var articles []*entity.Article

	if req.CategoryID != 0 {
		category, err := s.svcCtx.CategoryRepository.FindCategoryById(reqCtx, req.CategoryID)
		if err != nil {
			return nil, err
		}
		articles, err = s.svcCtx.ArticleRepository.FindArticleListByCategoryId(reqCtx, category.ID)
		data.ConditionName = category.CategoryName
	} else if req.TagID != 0 {
		tag, err := s.svcCtx.TagRepository.FindTagById(reqCtx, req.TagID)
		if err != nil {
			return nil, err
		}
		articles, err = s.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, tag.ID)
		data.ConditionName = tag.TagName
	}

	var list []*response.ArticleHome
	for _, article := range articles {
		//查询文章分类
		category, _ := s.svcCtx.CategoryRepository.FindCategoryById(reqCtx, article.CategoryID)
		// 查询文章标签
		tags, _ := s.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

		articleVO := &response.ArticleHome{}
		articleVO.ArticleDTO = convertArticle(article)
		articleVO.ArticleCategory = convertCategory(category)
		articleVO.ArticleTagList = convertTagList(tags)
		list = append(list, articleVO)
	}

	data.ArticleDTOList = list
	return data, err
}

// 文章时间轴
func (s *ArticleService) FindArticleArchives(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticlePreviewDTO, total int64, err error) {
	// 查找最新数据
	page.Sorts = []*sqlx.Sort{
		{Field: "id", Order: "desc"},
	}
	page.Conditions = append(page.Conditions, sqlx.NewCondition("status = ?", entity.ArticleStatusPublic))
	newestArticle, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}

	total, err = s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return convertArticlePreviewList(newestArticle), total, err
}

// 文章推荐
func (s *ArticleService) FindArticleDetails(reqCtx *request.Context, id int) (data *response.ArticlePageDetails, err error) {
	// 查询id对应文章
	article, err := s.svcCtx.ArticleRepository.FindArticleById(reqCtx, id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, _ := s.svcCtx.CategoryRepository.FindCategoryById(reqCtx, article.CategoryID)

	// 查询文章标签
	tags, _ := s.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

	// 查询推荐文章
	rmArticle, err := s.svcCtx.ArticleRepository.FindRecommendArticle(reqCtx, article.CategoryID)
	if err != nil {
		return nil, err
	}
	// 查询最新文章
	page := &request.PageQuery{
		PageLimit: sqlx.PageLimit{
			Page:     0,
			PageSize: 5,
		},
		Sorts: []*sqlx.Sort{
			{Field: "id", Order: "desc"},
		},
	}
	newestArticle, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, err
	}
	// 查询上一篇文章
	lastArticle, err := s.svcCtx.ArticleRepository.FindLastArticle(reqCtx, id)
	if err != nil {
		return nil, err
	}
	// 查询下一篇文章
	nextArticle, err := s.svcCtx.ArticleRepository.FindNextArticle(reqCtx, id)
	if err != nil {
		return nil, err
	}

	resp := &response.ArticlePageDetails{}
	resp.ArticleDTO = convertArticle(article)
	resp.ArticleCategory = convertCategory(category)
	resp.ArticleTagList = convertTagList(tags)
	resp.RecommendArticleList = convertArticlePreviewList(rmArticle)
	resp.NewestArticleList = convertArticlePreviewList(newestArticle)
	resp.LastArticle = convertArticlePreview(lastArticle)
	resp.NextArticle = convertArticlePreview(nextArticle)
	return resp, nil
}

// 分页获取Article记录
func (s *ArticleService) FindArticleHomeList(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticleHome, total int64, err error) {
	page.Sorts = append(page.Sorts, &sqlx.Sort{Field: "is_top", Order: "desc"})
	page.Conditions = append(page.Conditions, sqlx.NewCondition("status = ?", entity.ArticleStatusPublic))
	// 查询文章列表
	articles, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}

	// id
	var articleIds []int
	var categoryIds []int
	for _, article := range articles {
		articleIds = append(articleIds, article.ID)
		categoryIds = append(categoryIds, article.CategoryID)
	}

	// 查询所有文章分类
	category, _ := s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, nil, nil, sqlx.NewCondition("id in ?", categoryIds))
	var cmp = make(map[int]*entity.Category)
	for _, item := range category {
		cmp[item.ID] = item
	}

	// 查询所有文章标签映射
	amp, _ := s.svcCtx.TagRepository.FindArticleTagMap(reqCtx, articleIds)

	for _, article := range articles {

		articleVO := &response.ArticleHome{}
		articleVO.ArticleDTO = convertArticle(article)
		articleVO.ArticleCategory = convertCategory(cmp[article.CategoryID])
		articleVO.ArticleTagList = convertTagList(amp[article.ID])
		list = append(list, articleVO)
	}
	return list, total, err
}

func getCategoryName(category *entity.Category) string {
	if category == nil {
		return ""
	}
	return category.CategoryName
}

func getTagNameList(tags []*entity.Tag) (list []string) {
	for _, tag := range tags {
		list = append(list, tag.TagName)
	}

	if len(list) == 0 {
		return []string{}
	}

	return list
}
