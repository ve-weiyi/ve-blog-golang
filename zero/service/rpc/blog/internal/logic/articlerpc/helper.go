package articlerpclogic

import (
	"context"
	"fmt"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type ArticleHelperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleHelperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleHelperLogic {
	return &ArticleHelperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func convertArticleIn(in *articlerpc.ArticleNewReq) (out *model.Article) {
	out = &model.Article{
		Id:             in.Id,
		UserId:         in.UserId,
		CategoryId:     0,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		ArticleType:    in.ArticleType,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          0,
		IsDelete:       0,
		Status:         in.Status,
		LikeCount:      0,
	}

	return out
}

func convertCategoryIn(in *articlerpc.CategoryNewReq) (out *model.Category) {
	out = &model.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
	}

	return out
}

func convertTagIn(in *articlerpc.TagNewReq) (out *model.Tag) {
	out = &model.Tag{
		Id:      in.Id,
		TagName: in.TagName,
	}

	return out
}

func convertArticlePreviewOut(entity *model.Article) (out *articlerpc.ArticlePreview) {
	out = &articlerpc.ArticlePreview{
		Id:           entity.Id,
		ArticleCover: entity.ArticleCover,
		ArticleTitle: entity.ArticleTitle,
		CreatedAt:    entity.CreatedAt.Unix(),
	}
	return out
}

func convertArticleOut(entity *model.Article, acm map[int64]*model.Category, atm map[int64][]*model.Tag) (out *articlerpc.ArticleDetails) {
	out = &articlerpc.ArticleDetails{
		Id:             entity.Id,
		UserId:         entity.UserId,
		CategoryId:     entity.CategoryId,
		ArticleCover:   entity.ArticleCover,
		ArticleTitle:   entity.ArticleTitle,
		ArticleContent: entity.ArticleContent,
		ArticleType:    entity.ArticleType,
		OriginalUrl:    entity.OriginalUrl,
		IsTop:          entity.IsTop,
		IsDelete:       entity.IsDelete,
		Status:         entity.Status,
		CreatedAt:      entity.CreatedAt.Unix(),
		UpdatedAt:      entity.UpdatedAt.Unix(),
		LikeCount:      entity.LikeCount,
		Category:       nil,
		TagList:        nil,
	}

	if v, ok := acm[entity.Id]; ok {
		out.Category = &articlerpc.ArticleCategory{
			Id:           v.Id,
			CategoryName: v.CategoryName,
		}
	}

	if v, ok := atm[entity.Id]; ok {
		list := make([]*articlerpc.ArticleTag, 0, len(v))
		for _, tag := range v {
			list = append(list, &articlerpc.ArticleTag{
				Id:      tag.Id,
				TagName: tag.TagName,
			})
		}
		out.TagList = list
	}

	return out
}

func convertCategoryOut(in *model.Category, acm map[int64]int) (out *articlerpc.CategoryDetails) {
	out = &articlerpc.CategoryDetails{
		Id:           in.Id,
		CategoryName: in.CategoryName,
		ArticleCount: 0,
		CreatedAt:    in.CreatedAt.Unix(),
		UpdatedAt:    in.UpdatedAt.Unix(),
	}

	if v, ok := acm[in.Id]; ok {
		out.ArticleCount = int64(v)
	}

	return out
}

func convertTagOut(in *model.Tag, acm map[int64]int) (out *articlerpc.TagDetails) {
	out = &articlerpc.TagDetails{
		Id:           in.Id,
		TagName:      in.TagName,
		ArticleCount: 0,
		CreatedAt:    in.CreatedAt.Unix(),
		UpdatedAt:    in.UpdatedAt.Unix(),
	}

	if v, ok := acm[in.Id]; ok {
		out.ArticleCount = int64(v)
	}

	return out
}

func (l *ArticleHelperLogic) findOrAddCategory(name string) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("category name is empty")
	}

	category, err := l.svcCtx.CategoryModel.FindOneByCategoryName(l.ctx, name)
	if err != nil {
		insert := &model.Category{
			CategoryName: name,
		}
		_, err := l.svcCtx.CategoryModel.Insert(l.ctx, insert)
		if err != nil {
			return 0, err
		}
		return insert.Id, nil
	}

	return category.Id, nil
}

func (l *ArticleHelperLogic) findOrAddTag(name string) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("tag name is empty")
	}

	tag, err := l.svcCtx.TagModel.FindOneByTagName(l.ctx, name)
	if err != nil {
		insert := &model.Tag{
			TagName: name,
		}
		_, err := l.svcCtx.TagModel.Insert(l.ctx, insert)
		if err != nil {
			return 0, err
		}
		return insert.Id, nil
	}

	return tag.Id, nil
}

func (l *ArticleHelperLogic) findArticleCountGroupCategory(list []*model.Category) (acm map[int64]int, err error) {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}

	// 查询每个 category_id 的文章数量
	var results []struct {
		CategoryID   int64 `gorm:"column:category_id"`
		ArticleCount int   `gorm:"column:article_count"`
	}

	err = l.svcCtx.Gorm.Model(&model.Article{}).
		Select("category_id, COUNT(*) as article_count").
		Where("category_id IN ?", ids).
		Group("category_id").
		Order("category_id").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]int)
	for _, result := range results {
		acm[result.CategoryID] = result.ArticleCount
	}

	return acm, nil
}

func (l *ArticleHelperLogic) findArticleCountGroupTag(list []*model.Tag) (acm map[int64]int, err error) {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	// 查询每个 tag_id 的文章数量
	var results []struct {
		TagID        int64 `gorm:"column:tag_id"`
		ArticleCount int   `gorm:"column:article_count"`
	}

	err = l.svcCtx.Gorm.Model(&model.ArticleTag{}).
		Select("tag_id, COUNT(*) as article_count").
		Where("tag_id IN ?", ids).
		Group("tag_id").
		Order("tag_id").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]int)
	for _, result := range results {
		acm[result.TagID] = result.ArticleCount
	}

	return acm, nil
}

func (l *ArticleHelperLogic) findCategoryGroupArticle(list []*model.Article) (acm map[int64]*model.Category, err error) {
	var categoryIds []int64
	for _, v := range list {
		categoryIds = append(categoryIds, v.CategoryId)
	}

	cs, err := l.svcCtx.CategoryModel.FindALL(l.ctx, "id IN ?", categoryIds)
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]*model.Category)
	for _, v := range list {
		for _, category := range cs {
			if category.Id == v.CategoryId {
				acm[v.Id] = category
			}
		}
	}

	return acm, nil
}

func (l *ArticleHelperLogic) findTagGroupArticle(list []*model.Article) (atm map[int64][]*model.Tag, err error) {
	var articleIds []int64
	for _, v := range list {
		articleIds = append(articleIds, v.Id)
	}

	ats, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "article_id in (?)", articleIds)
	if err != nil {
		return nil, err
	}

	var tagIds []int64
	for _, v := range ats {
		tagIds = append(tagIds, v.TagId)
	}

	ts, err := l.svcCtx.TagModel.FindALL(l.ctx, "id in (?)", tagIds)
	if err != nil {
		return nil, err
	}

	atm = make(map[int64][]*model.Tag)
	for _, v := range ats {
		for _, tag := range ts {
			if tag.Id == v.TagId {
				atm[v.ArticleId] = append(atm[v.ArticleId], tag)
			}
		}
	}

	return atm, nil
}

func (l *ArticleHelperLogic) convertArticleQuery(in *articlerpc.FindArticleListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")

	if sorts == "" {
		sorts = "id desc"
	}

	if in.Status != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "status = ?"
		params = append(params, in.Status)
	}

	if in.IsTop != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "is_top = ?"
		params = append(params, in.IsTop)
	}

	if in.IsDelete != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "is_delete = ?"
		params = append(params, in.IsDelete)
	}

	if in.ArticleType != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "article_type = ?"
		params = append(params, in.ArticleType)
	}

	if in.ArticleTitle != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "article_title like ?"
		params = append(params, "%"+in.ArticleTitle+"%")
	}

	if in.CategoryName != "" {
		category, err := l.svcCtx.CategoryModel.FindOneByCategoryName(l.ctx, in.CategoryName)
		if err != nil {
			return
		}

		if conditions != "" {
			conditions += " and "
		}
		conditions += "category_id = ?"
		params = append(params, category.Id)
	}

	if in.TagName != "" {
		tag, err := l.svcCtx.TagModel.FindOneByTagName(l.ctx, in.TagName)
		if err != nil {
			return
		}
		ats, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "tag_id = ?", tag.Id)
		if err != nil {
			return
		}

		var articleIds []int64
		for _, v := range ats {
			articleIds = append(articleIds, v.ArticleId)
		}
		if conditions != "" {
			conditions += " and "
		}
		conditions += "id in (?)"
		params = append(params, articleIds)
	}

	return page, size, sorts, conditions, params
}
