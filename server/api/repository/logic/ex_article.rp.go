package logic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

// 根据分类id获取文章
func (s *ArticleRepository) FindArticleListByCategoryId(ctx context.Context, categoryId int) (list []*entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Model(&entity.Article{}).Where("category_id = ?", categoryId).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 根据标签id获取文章
func (s *ArticleRepository) FindArticleListByTagId(ctx context.Context, tagId int) (list []*entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 获取文章标签映射
	var ats []*entity.ArticleTag
	err = db.Where("tag_id = ?", tagId).Find(&ats).Error
	if err != nil {
		return nil, err
	}

	// 获取文章id列表
	var ids []int
	for _, at := range ats {
		ids = append(ids, at.ArticleID)
	}

	// 获取文章列表
	err = db.Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 获取推荐文章,与id相同分类的文章
func (s *ArticleRepository) FindRecommendArticle(ctx context.Context, cateId int) (list []*entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Where("category_id = ?", cateId).Limit(5).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取上一篇文章
func (s *ArticleRepository) FindLastArticle(ctx context.Context, id int) (out *entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	var lastArticle entity.Article
	err = db.Where("id < ?", id).Order("`id` desc").First(&lastArticle).Error
	if err != nil {
		return nil, nil
	}

	return &lastArticle, nil
}

// 获取下一篇文章
func (s *ArticleRepository) FindNextArticle(ctx context.Context, id int) (out *entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)
	var nextArticle entity.Article
	err = db.Where("id > ?", id).Order("`id` asc").First(&nextArticle).Error
	if err != nil {
		return nil, nil
	}

	return &nextArticle, nil
}

// 修改文章删除状态
func (s *ArticleRepository) UpdateArticleDelete(ctx context.Context, id int, delete int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Model(&entity.Article{}).Where("id = ?", id).Update("is_delete", delete)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 修改文章置顶状态
func (s *ArticleRepository) UpdateArticleTop(ctx context.Context, id int, top int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Model(&entity.Article{}).Where("id = ?", id).Update("is_top", top)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}
