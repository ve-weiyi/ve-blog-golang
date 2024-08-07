package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type ArticleService struct {
	svcCtx *svctx.ServiceContext
}

func NewArticleService(svcCtx *svctx.ServiceContext) *ArticleService {
	return &ArticleService{
		svcCtx: svcCtx,
	}
}

// 创建Article记录
func (l *ArticleService) SaveArticle(reqCtx *request.Context, req *dto.ArticleDetailsDTOReq) (data *entity.Article, err error) {
	// 创建文章
	article := &entity.Article{
		Id:             req.Id,
		UserId:         reqCtx.Uid,
		ArticleCover:   req.ArticleCover,
		ArticleTitle:   req.ArticleTitle,
		ArticleContent: req.ArticleContent,
		Type:           req.Type,
		OriginalUrl:    req.OriginalUrl,
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
		article.CategoryId = category.Id
	}

	// 创建文章或保存文章
	if article.Id == 0 {
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
	_, _ = l.svcCtx.ArticleTagRepository.Delete(reqCtx, "article_id = ?", article.Id)
	// 创建不存在的标签
	tags, _ := l.svcCtx.TagRepository.InsertBatchTagNotExist(reqCtx, req.TagNameList)
	// 创建文章标签映射
	for _, tag := range tags {
		at := &entity.ArticleTag{
			ArticleId: article.Id,
			TagId:     tag.Id,
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
func (l *ArticleService) FindArticle(reqCtx *request.Context, req *request.IdReq) (data *dto.ArticleBack, err error) {
	// 查询id对应文章
	article, err := l.svcCtx.ArticleRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryId)

	// 查询文章标签
	tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.Id)

	resp := &dto.ArticleBack{}
	resp.ArticleDTO = convertArticle(article)
	resp.CategoryName = getCategoryName(category)
	resp.TagNameList = getTagNameList(tags)
	return resp, nil
}

// 分页获取Article记录
func (l *ArticleService) FindArticleList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.ArticleBack, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询文章列表
	articles, err := l.svcCtx.ArticleRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = l.svcCtx.ArticleRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// id
	var articleIds []int64
	var categoryIds []int64
	for _, article := range articles {
		articleIds = append(articleIds, article.Id)
		categoryIds = append(categoryIds, article.CategoryId)
	}

	// 查询所有文章分类
	category, _ := l.svcCtx.CategoryRepository.FindALL(reqCtx, "id in (?)", categoryIds)
	var cmp = make(map[int64]*entity.Category)
	for _, item := range category {
		cmp[item.Id] = item
	}

	// 查询所有文章标签映射
	amp, _ := l.svcCtx.TagRepository.FindArticleTagMap(reqCtx, articleIds)

	for _, article := range articles {

		articleDTO := &dto.ArticleBack{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.CategoryName = getCategoryName(cmp[article.CategoryId])
		articleDTO.TagNameList = getTagNameList(amp[article.Id])
		list = append(list, articleDTO)
	}
	return list, total, err
}

// 更新Article删除状态
func (l *ArticleService) UpdateArticleDelete(reqCtx *request.Context, req *dto.ArticleDeleteReq) (rows int, err error) {
	return l.svcCtx.ArticleRepository.UpdateArticleDelete(reqCtx, req.Id, req.IsDelete)
}

// 更新Article记录
func (l *ArticleService) UpdateArticleTop(reqCtx *request.Context, req *dto.ArticleTopReq) (rows int, err error) {
	return l.svcCtx.ArticleRepository.UpdateArticleTop(reqCtx, req.Id, req.IsTop)
}

// 文章归类 category
func (l *ArticleService) FindArticleClassifyCategory(reqCtx *request.Context, req *dto.ArticleClassifyReq) (data *dto.ArticleClassifyResp, err error) {
	data = &dto.ArticleClassifyResp{}
	// 查询文章列表
	var articles []*entity.Article

	category, err := l.svcCtx.CategoryRepository.First(reqCtx, "category_name = ?", req.ClassifyName)
	if err != nil {
		return nil, err
	}

	articles, err = l.svcCtx.ArticleRepository.FindArticleListByCategoryId(reqCtx, category.Id)
	if err != nil {
		return nil, err
	}

	data.ConditionName = category.CategoryName

	var list []*dto.ArticleHome
	for _, article := range articles {
		// 查询文章分类
		ctg, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryId)
		// 查询文章标签
		tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.Id)

		articleDTO := &dto.ArticleHome{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.ArticleCategory = convertCategory(ctg)
		articleDTO.ArticleTagList = convertTagList(tags)
		list = append(list, articleDTO)
	}

	data.ArticleList = list
	return data, err
}

// 文章归类 tag
func (l *ArticleService) FindArticleClassifyTag(reqCtx *request.Context, req *dto.ArticleClassifyReq) (data *dto.ArticleClassifyResp, err error) {
	data = &dto.ArticleClassifyResp{}
	// 查询文章列表
	var articles []*entity.Article

	tag, err := l.svcCtx.TagRepository.First(reqCtx, "tag_name = ?", req.ClassifyName)
	if err != nil {
		return nil, err
	}

	articles, err = l.svcCtx.ArticleRepository.FindArticleListByTagId(reqCtx, tag.Id)
	if err != nil {
		return nil, err
	}

	data.ConditionName = tag.TagName

	var list []*dto.ArticleHome
	for _, article := range articles {
		// 查询文章分类
		ctg, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryId)
		// 查询文章标签
		tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.Id)

		articleDTO := &dto.ArticleHome{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.ArticleCategory = convertCategory(ctg)
		articleDTO.ArticleTagList = convertTagList(tags)
		list = append(list, articleDTO)
	}

	data.ArticleList = list
	return data, err
}

// 文章时间轴
func (l *ArticleService) FindArticleArchives(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.ArticlePreviewDTO, total int64, err error) {
	p, s := page.PageClause()
	// 查找最新数据
	newestArticle, err := l.svcCtx.ArticleRepository.FindList(reqCtx, p, s, "id desc", "status = ?", entity.ArticleStatusPublic)
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
func (l *ArticleService) FindArticleRecommend(reqCtx *request.Context, req *request.IdReq) (data *dto.ArticlePageDetailsDTO, err error) {
	// 查询id对应文章
	article, err := l.svcCtx.ArticleRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	// 查询文章分类
	category, _ := l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", article.CategoryId)

	// 查询文章标签
	tags, _ := l.svcCtx.TagRepository.FindArticleTagList(reqCtx, article.Id)

	// 查询推荐文章
	rmArticle, err := l.svcCtx.ArticleRepository.FindRecommendArticle(reqCtx, article.CategoryId)
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

	resp := &dto.ArticlePageDetailsDTO{}
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
func (l *ArticleService) FindArticleHomeList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.ArticleHome, total int64, err error) {
	page.Sorts = append(page.Sorts, &dto.PageSort{Field: "is_top", Order: "desc"})
	page.Conditions = append(page.Conditions, &dto.PageCondition{Field: "status", Operator: "=", Value: entity.ArticleStatusPublic})

	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询文章列表
	articles, err := l.svcCtx.ArticleRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 查询文章总数
	total, err = l.svcCtx.ArticleRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// id
	var articleIds []int64
	var categoryIds []int64
	for _, article := range articles {
		articleIds = append(articleIds, article.Id)
		categoryIds = append(categoryIds, article.CategoryId)
	}

	// 查询所有文章分类
	category, _ := l.svcCtx.CategoryRepository.FindALL(reqCtx, "id in (?)", categoryIds)
	var cmp = make(map[int64]*entity.Category)
	for _, item := range category {
		cmp[item.Id] = item
	}

	// 查询所有文章标签映射
	amp, _ := l.svcCtx.TagRepository.FindArticleTagMap(reqCtx, articleIds)

	for _, article := range articles {

		articleDTO := &dto.ArticleHome{}
		articleDTO.ArticleDTO = convertArticle(article)
		articleDTO.ArticleCategory = convertCategory(cmp[article.CategoryId])
		articleDTO.ArticleTagList = convertTagList(amp[article.Id])
		list = append(list, articleDTO)
	}
	return list, total, err
}

func (l *ArticleService) LikeArticle(reqCtx *request.Context, req *request.IdReq) (data interface{}, err error) {
	return l.svcCtx.ArticleRepository.LikeArticle(reqCtx, reqCtx.Uid, req.Id)
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
