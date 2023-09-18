package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

func (s *TagRepository) FindArticleTagList(articleId int) (list []*entity.Tag, err error) {
	// 创建db
	db := s.DbEngin
	var ats []*entity.ArticleTag
	var tags []*entity.Tag

	err = db.Where("article_id = ?", articleId).Find(&ats).Error
	if err != nil {
		return nil, err
	}

	var tagIds []int
	for _, at := range ats {
		tagIds = append(tagIds, at.TagID)
	}

	err = db.Where("id in (?)", tagIds).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
