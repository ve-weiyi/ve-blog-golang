package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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
func (s *CategoryRepository) CreateCategory(ctx context.Context, category *entity.Category) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

// 更新Category记录
func (s *CategoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, err
}

// 删除Category记录
func (s *CategoryRepository) DeleteCategory(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Category{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询Category记录
func (s *CategoryRepository) FindCategory(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 分页查询Category记录
func (s *CategoryRepository) FindCategoryList(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.Category, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sqlx.OrderClause(sorts))
	}

	// 如果有分页参数
	if page != nil && page.IsValid() {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 查询总数
func (s *CategoryRepository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.Category{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Category记录——根据id
func (s *CategoryRepository) FindCategoryById(ctx context.Context, id int) (out *entity.Category, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 删除Category记录——根据id
func (s *CategoryRepository) DeleteCategoryById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.Category{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除Category记录——根据ids
func (s *CategoryRepository) DeleteCategoryByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.Category{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}
