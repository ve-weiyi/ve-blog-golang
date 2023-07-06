package logic

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type CategoryRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewCategoryRepository(svcCtx *svc.RepositoryContext) *CategoryRepository {
	return &CategoryRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Category记录
func (s *CategoryRepository) CreateCategory(category *entity.Category) (out *entity.Category, err error) {
	db := s.DbEngin
	err = db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

// 删除Category记录
func (s *CategoryRepository) DeleteCategory(category *entity.Category) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&category)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Category记录
func (s *CategoryRepository) UpdateCategory(category *entity.Category) (out *entity.Category, err error) {
	db := s.DbEngin
	err = db.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

// 根据id获取Category记录
func (s *CategoryRepository) FindCategory(id int) (out *entity.Category, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Category记录
func (s *CategoryRepository) DeleteCategoryByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Category{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取Category记录
func (s *CategoryRepository) GetCategoryList(page *request.PageInfo) (list []*entity.Category, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(page.OrderClause())
	}

	// 如果有分页参数
	if page.Page != 0 || page.PageSize != 0 {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询总数
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
