package articlerpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

func findOrAddCategory(ctx context.Context, svcCtx *svc.ServiceContext, name string) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("category name is empty")
	}

	category, err := svcCtx.CategoryModel.FindOneByCategoryName(ctx, name)
	if err != nil {
		insert := &model.Category{
			CategoryName: name,
		}
		_, err := svcCtx.CategoryModel.Insert(ctx, insert)
		if err != nil {
			return 0, err
		}
		return insert.Id, nil
	}

	return category.Id, nil
}

func findOrAddTag(ctx context.Context, svcCtx *svc.ServiceContext, name string) (int64, error) {
	if name == "" {
		return 0, fmt.Errorf("tag name is empty")
	}

	tag, err := svcCtx.TagModel.FindOneByTagName(ctx, name)
	if err != nil {
		insert := &model.Tag{
			TagName: name,
		}
		_, err := svcCtx.TagModel.Insert(ctx, insert)
		if err != nil {
			return 0, err
		}
		return insert.Id, nil
	}

	return tag.Id, nil
}

func convertArticleIn(in *blog.ArticleNew) (out *model.Article) {
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

func convertCategoryIn(in *blog.CategoryNew) (out *model.Category) {
	out = &model.Category{
		Id:           in.Id,
		CategoryName: in.CategoryName,
	}

	return out
}

func convertTagIn(in *blog.TagNew) (out *model.Tag) {
	out = &model.Tag{
		Id:      in.Id,
		TagName: in.TagName,
	}

	return out
}

func convertArticleOut(entity *model.Article, acm map[int64]*model.Category, atm map[int64][]*model.Tag) (out *blog.ArticleDetails) {
	out = &blog.ArticleDetails{
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

	if v, ok := acm[entity.CategoryId]; ok {
		out.Category = &blog.ArticleCategory{
			Id:           v.Id,
			CategoryName: v.CategoryName,
		}
	}

	if v, ok := atm[entity.Id]; ok {
		list := make([]*blog.ArticleTag, 0, len(v))
		for _, tag := range v {
			list = append(list, &blog.ArticleTag{
				Id:      tag.Id,
				TagName: tag.TagName,
			})
		}
		out.TagList = list
	}

	return out
}

func convertCategoryOut(in *model.Category, acm map[int64]int) (out *blog.CategoryDetails) {
	out = &blog.CategoryDetails{
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

func convertTagOut(in *model.Tag, acm map[int64]int) (out *blog.TagDetails) {
	out = &blog.TagDetails{
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

func findArticleCountGroupCategory(ctx context.Context, svcCtx *svc.ServiceContext, list []*model.Category) (acm map[int64]int, err error) {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}

	// 查询每个 category_id 的文章数量
	var results []struct {
		CategoryID   int64 `gorm:"column:category_id"`
		ArticleCount int   `gorm:"column:article_count"`
	}

	err = svcCtx.Gorm.Model(&model.Article{}).
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

func findArticleCountGroupTag(ctx context.Context, svcCtx *svc.ServiceContext, list []*model.Tag) (acm map[int64]int, err error) {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	// 查询每个 tag_id 的文章数量
	var results []struct {
		TagID        int64 `gorm:"column:tag_id"`
		ArticleCount int   `gorm:"column:article_count"`
	}

	err = svcCtx.Gorm.Model(&model.ArticleTag{}).
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

func findCategoryGroupArticle(ctx context.Context, svcCtx *svc.ServiceContext, list []*model.Article) (acm map[int64]*model.Category, err error) {
	var articleIds []int64
	for _, v := range list {
		articleIds = append(articleIds, v.Id)
	}

	records, err := svcCtx.CategoryModel.FindALL(ctx, "id IN ?", articleIds)
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]*model.Category)
	for _, result := range records {
		acm[result.Id] = result
	}

	return acm, nil
}
func findTagGroupArticle(ctx context.Context, svcCtx *svc.ServiceContext, list []*model.Article) (atm map[int64][]*model.Tag, err error) {
	var articleIds []int64
	for _, v := range list {
		articleIds = append(articleIds, v.Id)
	}

	ats, err := svcCtx.ArticleTagModel.FindALL(ctx, "article_id in (?)", articleIds)
	if err != nil {
		return nil, err
	}

	var tidm = make(map[int64][]int64)
	var tids []int64
	for _, v := range ats {
		tidm[v.ArticleId] = append(tidm[v.ArticleId], v.TagId)
		tids = append(tids, v.TagId)
	}

	ts, err := svcCtx.TagModel.FindALL(ctx, "id in ?", tids)
	if err != nil {
		return nil, err
	}

	atm = make(map[int64][]*model.Tag)
	for k, v := range tidm {
		var tags []*model.Tag
		for _, vv := range v {
			for _, tag := range ts {
				if tag.Id == vv {
					tags = append(tags, tag)
				}
			}
		}
		atm[k] = tags
	}

	return atm, nil
}
