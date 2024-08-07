package service

import (
	"math/rand"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

func convertLoginHistory(entity *entity.UserLoginHistory) *dto.LoginHistory {
	return &dto.LoginHistory{
		Id:        entity.Id,
		LoginType: entity.LoginType,
		Agent:     entity.Agent,
		IpAddress: entity.IpAddress,
		IpSource:  entity.IpSource,
		LoginTime: entity.CreatedAt.String(),
	}
}

func convertArticle(article *entity.Article) dto.ArticleDTO {
	out := dto.ArticleDTO{
		Id:             article.Id,
		ArticleCover:   article.ArticleCover,
		ArticleTitle:   article.ArticleTitle,
		ArticleContent: article.ArticleContent,
		LikeCount:      100,
		ViewsCount:     200,
		IsTop:          article.IsTop,
		IsDelete:       article.IsDelete,
		Type:           article.Type,
		Status:         article.Status,
		OriginalUrl:    article.OriginalUrl,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
	}
	return out
}

func convertArticlePreviewList(list []*entity.Article) []*dto.ArticlePreviewDTO {
	var out []*dto.ArticlePreviewDTO
	for _, item := range list {
		at := &dto.ArticlePreviewDTO{
			Id:           item.Id,
			ArticleCover: item.ArticleCover,
			ArticleTitle: item.ArticleTitle,
			CreatedAt:    item.CreatedAt,
		}
		out = append(out, at)
	}

	return out
}

func convertArticlePreview(article *entity.Article) *dto.ArticlePreviewDTO {
	if article == nil {
		return nil
	}
	return &dto.ArticlePreviewDTO{
		Id:           article.Id,
		ArticleCover: article.ArticleCover,
		ArticleTitle: article.ArticleTitle,
	}
}

func convertArticleStatisticsList(list []*entity.Article) []*dto.ArticleStatisticsDTO {
	var dateMap = make(map[string]int64)
	for _, item := range list {
		key := item.CreatedAt.Format("2006-01-02")
		if _, ok := dateMap[key]; ok {
			dateMap[key]++
		} else {
			dateMap[key] = 1
		}
	}
	// 日期统计
	var out []*dto.ArticleStatisticsDTO
	for k, v := range dateMap {
		at := &dto.ArticleStatisticsDTO{
			Day:   k,
			Count: v,
		}
		out = append(out, at)
	}

	return out
}

func convertUniqueViewList(list []*entity.UniqueView) []*dto.UniqueViewDTO {
	var out []*dto.UniqueViewDTO
	for _, item := range list {
		at := &dto.UniqueViewDTO{
			Day:   item.CreatedAt.Format("2006-01-02"),
			Count: rand.Int63n(100),
		}
		out = append(out, at)
	}

	return out
}

func convertArticleRankList(list []*entity.Article) []*dto.ArticleViewRankDTO {
	var out []*dto.ArticleViewRankDTO
	for _, item := range list {
		at := &dto.ArticleViewRankDTO{
			Id:           item.Id,
			ArticleTitle: item.ArticleTitle,
			Count:        rand.Int63n(100),
		}
		out = append(out, at)
	}

	return out
}

func convertCategory(entity *entity.Category) *dto.CategoryDTO {
	if entity == nil {
		return &dto.CategoryDTO{}
	}
	return &dto.CategoryDTO{
		Id:           entity.Id,
		CategoryName: entity.CategoryName,
	}
}

func convertTagList(list []*entity.Tag) []*dto.TagDTO {
	var tagList []*dto.TagDTO
	for _, tag := range list {
		tagList = append(tagList, &dto.TagDTO{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}
	return tagList
}

func convertCategoryList(list []*entity.Category) []*dto.CategoryDTO {
	var categoryList []*dto.CategoryDTO

	for _, in := range list {
		data := &dto.CategoryDTO{
			Id:           in.Id,
			CategoryName: in.CategoryName,
		}
		categoryList = append(categoryList, data)
	}
	return categoryList
}

func convertPageList(list []*entity.Page) []*dto.PageDTO {
	var pageList []*dto.PageDTO

	for _, in := range list {
		data := &dto.PageDTO{
			Id:        in.Id,
			PageName:  in.PageName,
			PageLabel: in.PageLabel,
			PageCover: in.PageCover,
		}
		pageList = append(pageList, data)
	}
	return pageList
}

func convertRoleList(list []*entity.Role) []*dto.RoleDTO {
	var roleList []*dto.RoleDTO

	for _, in := range list {
		data := &dto.RoleDTO{
			RoleName:    in.RoleName,
			RoleComment: in.RoleComment,
		}
		roleList = append(roleList, data)
	}
	return roleList
}

func convertMenu(entity *entity.Menu) *dto.MenuDetailsDTO {
	if entity == nil {
		return nil
	}

	meta := dto.Meta{}
	jsonconv.JsonToObject(entity.Extra, &meta)

	return &dto.MenuDetailsDTO{
		Id:        entity.Id,
		Type:      entity.Type,
		Name:      entity.Name,
		Path:      entity.Path,
		Component: entity.Component,
		Title:     entity.Title,
		ParentId:  entity.ParentId,
		Meta:      meta,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		Children:  nil,
	}
}
