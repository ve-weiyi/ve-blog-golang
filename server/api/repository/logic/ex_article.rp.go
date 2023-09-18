package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

// 根据分类id获取文章
func (s *ArticleRepository) FindArticleListByCategoryId(categoryId int) (list []*entity.Article, total int64, err error) {
	db := s.DbEngin
	err = db.Model(&entity.Article{}).Where("category_id = ?", categoryId).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	total = int64(len(list))
	return list, total, nil
}

// 根据标签id获取文章
func (s *ArticleRepository) FindArticleListByTagId(tagId int) (list []*entity.Article, total int64, err error) {
	db := s.DbEngin

	// 获取文章标签映射
	var ats []*entity.ArticleTag
	err = db.Where("tag_id = ?", tagId).Find(&ats).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取文章id列表
	var ids []int
	for _, at := range ats {
		ids = append(ids, at.ArticleID)
	}

	// 获取文章列表
	err = db.Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	total = int64(len(list))
	return list, total, nil
}

// 获取推荐文章,与id相同分类的文章
func (s *ArticleRepository) FindRecommendArticle(cateId int) (list []*entity.Article, err error) {
	db := s.DbEngin
	err = db.Where("category_id = ?", cateId).Limit(5).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取上一篇文章
func (s *ArticleRepository) FindLastArticle(id int) (out *entity.Article, err error) {
	db := s.DbEngin
	var lastArticle entity.Article
	err = db.Where("id < ?", id).Order("`id` desc").First(&lastArticle).Error
	if err != nil {
		return nil, nil
	}

	return &lastArticle, nil
}

// 获取下一篇文章
func (s *ArticleRepository) FindNextArticle(id int) (out *entity.Article, err error) {
	db := s.DbEngin
	var nextArticle entity.Article
	err = db.Where("id > ?", id).Order("`id` asc").First(&nextArticle).Error
	if err != nil {
		return nil, nil
	}

	return &nextArticle, nil
}
