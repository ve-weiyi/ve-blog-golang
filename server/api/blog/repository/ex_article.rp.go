package repository

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/cache"
)

// 根据分类id获取文章
func (s *ArticleRepository) FindArticleListByCategoryId(ctx context.Context, categoryId int64) (list []*entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Model(&entity.Article{}).Where("category_id = ?", categoryId).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 根据标签id获取文章
func (s *ArticleRepository) FindArticleListByTagId(ctx context.Context, tagId int64) (list []*entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 获取文章标签映射
	var ats []*entity.ArticleTag
	err = db.Where("tag_id = ?", tagId).Find(&ats).Error
	if err != nil {
		return nil, err
	}

	// 获取文章id列表
	ids := make([]int64, 0)
	for _, at := range ats {
		ids = append(ids, at.ArticleId)
	}

	// 获取文章列表
	err = db.Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 获取推荐文章,与id相同分类的文章
func (s *ArticleRepository) FindRecommendArticle(ctx context.Context, cateId int64) (list []*entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Where("category_id = ?", cateId).Limit(5).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取上一篇文章
func (s *ArticleRepository) FindLastArticle(ctx context.Context, id int64) (out *entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	var lastArticle entity.Article
	err = db.Where("id < ?", id).Order("`id` desc").First(&lastArticle).Error
	if err != nil {
		return nil, nil
	}

	return &lastArticle, nil
}

// 获取下一篇文章
func (s *ArticleRepository) FindNextArticle(ctx context.Context, id int64) (out *entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	var nextArticle entity.Article
	err = db.Where("id > ?", id).Order("`id` asc").First(&nextArticle).Error
	if err != nil {
		return nil, nil
	}

	return &nextArticle, nil
}

// 修改文章删除状态
func (s *ArticleRepository) UpdateArticleDelete(ctx context.Context, id int64, delete int64) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Model(&entity.Article{}).Where("id = ?", id).Update("is_delete", delete)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 修改文章置顶状态
func (s *ArticleRepository) UpdateArticleTop(ctx context.Context, id int64, top int64) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Model(&entity.Article{}).Where("id = ?", id).Update("is_top", top)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 点赞评论
func (s *ArticleRepository) LikeArticle(ctx context.Context, uid int64, articleId int64) (data interface{}, err error) {
	// 用户点赞的评论列表
	articleUserLikeKey := cache.WrapCacheKey(cache.ArticleUserLike, uid)
	// 当前评论的点赞量
	articleLikeCountKey := cache.WrapCacheKey(cache.ArticleLikeCount, articleId)

	// 判断是否已经点赞
	if s.Cache.SIsMember(ctx, articleUserLikeKey, articleId).Val() {
		// 点过赞则删除评论id
		s.Cache.SRem(ctx, articleUserLikeKey, articleId)
		// 评论点赞量-1
		s.Cache.Decr(ctx, articleLikeCountKey)
	} else {
		// 未点赞则增加评论id
		s.Cache.SAdd(ctx, articleUserLikeKey, articleId)
		// 评论点赞量+1
		s.Cache.Incr(ctx, articleLikeCountKey)
	}

	return data, nil
}

// 获取用户点赞记录
func (s *ArticleRepository) FindUserLikeArticle(ctx context.Context, uid int64) (data []string, err error) {
	// 用户点赞的评论列表
	articleUserLikeKey := cache.WrapCacheKey(cache.ArticleUserLike, uid)
	return s.Cache.SMembers(ctx, articleUserLikeKey).Result()
}
