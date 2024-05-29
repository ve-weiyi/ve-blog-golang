package repository

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

// 批量创建category, 如果category已经存在则返回已经存在的category
func (s *CategoryRepository) CreateCategoryNotExist(ctx context.Context, categoryName string) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)

	var category entity.Category
	err = db.Where("category_name = ?", categoryName).First(&category).Error
	if err != nil {
		return nil, err
	}

	if category.Id != 0 {
		return &category, nil
	}

	category.CategoryName = categoryName
	err = db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
