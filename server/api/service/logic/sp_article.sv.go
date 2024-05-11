package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
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
func (l *ArticleService) SaveArticle(reqCtx *request.Context, req *request.ArticleDetailsDTOReq) (data *entity.Article, err error) {
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
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
	}

	// 设置默认文章封面
	if article.ArticleCover == "" {
		article.ArticleCover = "https://ve-blog.oss-cn-shenzhen.aliyuncs.com/ve-blog/2021/08/31/163040.jpg"
	}

	// 查找分类是否存在
	category, _ := l.svcCtx.CategoryRepository.CreateCategoryNotExist(reqCtx, req.CategoryName)
	if category != nil {
		article.CategoryID = category.ID
	}

	// 创建文章或保存文章
	if article.ID == 0 {
		_, err = l.svcCtx.ArticleRepository.Create(reqCtx, article)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = l.svcCtx.ArticleRepository.Update(reqCtx, article)
		if err != nil {
			return nil, err
		}
	}

	// 删除文章标签映射
	_, _ = l.svcCtx.ArticleTagRepository.Delete(reqCtx, "article_id = ?", article.ID)
	// 创建不存在的标签
	tags, _ := l.svcCtx.TagRepository.InsertBatchTagNotExist(reqCtx, req.TagNameList)
	// 创建文章标签映射
	for _, tag := range tags {
		at := &entity.ArticleTag{
			ArticleID: article.ID,
			TagID:     tag.ID,
		}
		_, _ = l.svcCtx.ArticleTagRepository.Create(reqCtx, at)
	}

	return article, nil
}

// 删除Article记录
func (l *ArticleService) DeleteArticle(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	// 删除文章标签映射
	_, err = l.svcCtx.ArticleTagRepository.Delete(reqCtx, "article_id = ?", req.Id)
	if err != nil {
		return 0, err
	}

	return l.svcCtx.ArticleRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 根据id获取Article记录
func (l *ArticleService) FindArticle(reqCtx *request.Context, req *request.IdReq) (data *response.ArticleBack, err error) {
	// 查询id对应文章
	article, err := l.svcCtx.ArticleRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryID)

	// 查询文章标签
	tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

	resp := &response.ArticleBack{}
	resp.ArticleDTO = convertArticle(article)
	resp.CategoryName = getCategoryName(category)
	resp.TagNameList = getTagNameList(tags)
	return resp, nil
}

// 分页获取Article记录
func (l *ArticleService) FindArticleList(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticleBack, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询文章列表
	articles, err := l.svcCtx.ArticleRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = l.svcCtx.ArticleRepository.Count(reqCtx, cond, args...)
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
	category, _ := l.svcCtx.CategoryRepository.FindALL(reqCtx, "id in ?", categoryIds)
	var cmp = make(map[int]*entity.Category)
	for _, item := range category {
		cmp[item.ID] = item
	}

	// 查询所有文章标签映射
	amp, _ := l.svcCtx.TagRepository.FindArticleTagMap(reqCtx, articleIds)

	for _, article := range articles {

		articleDTO := &response.ArticleBack{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.CategoryName = getCategoryName(cmp[article.CategoryID])
		articleDTO.TagNameList = getTagNameList(amp[article.ID])
		list = append(list, articleDTO)
	}
	return list, total, err
}

// 更新Article删除状态
func (l *ArticleService) UpdateArticleDelete(reqCtx *request.Context, req *request.ArticleDeleteReq) (rows int, err error) {
	return l.svcCtx.ArticleRepository.UpdateArticleDelete(reqCtx, req.ID, req.IsDelete)
}

// 更新Article记录
func (l *ArticleService) UpdateArticleTop(reqCtx *request.Context, req *request.ArticleTopReq) (rows int, err error) {
	return l.svcCtx.ArticleRepository.UpdateArticleTop(reqCtx, req.ID, req.IsTop)
}

// 文章归类
func (l *ArticleService) FindArticleSeries(reqCtx *request.Context, req *request.ArticleConditionReq) (data *response.ArticleConditionDTO, err error) {
	data = &response.ArticleConditionDTO{}
	// 查询文章列表
	var articles []*entity.Article

	if req.CategoryID != 0 {
		category, err := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", req.CategoryID)
		if err != nil {
			return nil, err
		}
		articles, err = l.svcCtx.ArticleRepository.FindArticleListByCategoryId(reqCtx, category.ID)
		data.ConditionName = category.CategoryName
	} else if req.TagID != 0 {
		tag, err := l.svcCtx.TagRepository.First(reqCtx, "id = ?", req.TagID)
		if err != nil {
			return nil, err
		}
		articles, err = l.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, tag.ID)
		data.ConditionName = tag.TagName
	}

	var list []*response.ArticleHome
	for _, article := range articles {
		//查询文章分类
		category, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryID)
		// 查询文章标签
		tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

		articleDTO := &response.ArticleHome{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.ArticleCategory = convertCategory(category)
		articleDTO.ArticleTagList = convertTagList(tags)
		list = append(list, articleDTO)
	}

	data.ArticleDTOList = list
	return data, err
}

// 文章时间轴
func (l *ArticleService) FindArticleArchives(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticlePreviewDTO, total int64, err error) {
	// 查找最新数据
	newestArticle, err := l.svcCtx.ArticleRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, "id desc", "status = ?", entity.ArticleStatusPublic)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.ArticleRepository.Count(reqCtx, "status = ?", entity.ArticleStatusPublic)
	if err != nil {
		return nil, 0, err
	}
	return convertArticlePreviewList(newestArticle), total, err
}

// 文章推荐
func (l *ArticleService) FindArticleDetails(reqCtx *request.Context, req *request.IdReq) (data *response.ArticlePageDetailsDTO, err error) {
	// 查询id对应文章
	article, err := l.svcCtx.ArticleRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryID)

	// 查询文章标签
	tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.ID)

	// 查询推荐文章
	rmArticle, err := l.svcCtx.ArticleRepository.FindRecommendArticle(reqCtx, article.CategoryID)
	if err != nil {
		return nil, err
	}
	// 查询最新文章
	newestArticle, err := l.svcCtx.ArticleRepository.FindList(reqCtx, 0, 5, "id desc", "")
	if err != nil {
		return nil, err
	}
	// 查询上一篇文章
	lastArticle, err := l.svcCtx.ArticleRepository.FindLastArticle(reqCtx, req.Id)
	if err != nil {
		return nil, err
	}
	// 查询下一篇文章
	nextArticle, err := l.svcCtx.ArticleRepository.FindNextArticle(reqCtx, req.Id)
	if err != nil {
		return nil, err
	}

	resp := &response.ArticlePageDetailsDTO{}
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
func (l *ArticleService) FindArticleHomeList(reqCtx *request.Context, page *request.PageQuery) (list []*response.ArticleHome, total int64, err error) {
	page.Sorts = append(page.Sorts, &request.PageSort{Field: "is_top", Order: "desc"})
	page.Conditions = append(page.Conditions, &request.PageCondition{Field: "status", Operator: "=", Value: entity.ArticleStatusPublic})

	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询文章列表
	articles, err := l.svcCtx.ArticleRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = l.svcCtx.ArticleRepository.Count(reqCtx, cond, args...)
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
	category, _ := l.svcCtx.CategoryRepository.FindALL(reqCtx, "id in ?", categoryIds)
	var cmp = make(map[int]*entity.Category)
	for _, item := range category {
		cmp[item.ID] = item
	}

	// 查询所有文章标签映射
	amp, _ := l.svcCtx.TagRepository.FindArticleTagMap(reqCtx, articleIds)

	for _, article := range articles {

		articleDTO := &response.ArticleHome{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.ArticleCategory = convertCategory(cmp[article.CategoryID])
		articleDTO.ArticleTagList = convertTagList(amp[article.ID])
		list = append(list, articleDTO)
	}
	return list, total, err
}

func (l *ArticleService) LikeArticle(reqCtx *request.Context, req *request.IdReq) (data interface{}, err error) {
	return l.svcCtx.ArticleRepository.LikeArticle(reqCtx, reqCtx.UID, req.Id)
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
