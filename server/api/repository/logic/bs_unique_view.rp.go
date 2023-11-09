package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
)

type UniqueViewRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUniqueViewRepository(svcCtx *svc.RepositoryContext) *UniqueViewRepository {
	return &UniqueViewRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建UniqueView记录
func (s *UniqueViewRepository) CreateUniqueView(ctx context.Context, uniqueView *entity.UniqueView) (out *entity.UniqueView, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&uniqueView).Error
	if err != nil {
		return nil, err
	}
	return uniqueView, err
}

// 更新UniqueView记录
func (s *UniqueViewRepository) UpdateUniqueView(ctx context.Context, uniqueView *entity.UniqueView) (out *entity.UniqueView, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&uniqueView).Error
	if err != nil {
		return nil, err
	}
	return uniqueView, err
}

// 删除UniqueView记录
func (s *UniqueViewRepository) DeleteUniqueView(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.UniqueView{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询UniqueView记录
func (s *UniqueViewRepository) FindUniqueView(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.UniqueView, err error) {
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

// 分页查询UniqueView记录
func (s *UniqueViewRepository) FindUniqueViewList(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.UniqueView, err error) {
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
func (s *UniqueViewRepository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.UniqueView{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UniqueView记录——根据id
func (s *UniqueViewRepository) FindUniqueViewById(ctx context.Context, id int) (out *entity.UniqueView, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 删除UniqueView记录——根据id
func (s *UniqueViewRepository) DeleteUniqueViewById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.UniqueView{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除UniqueView记录——根据ids
func (s *UniqueViewRepository) DeleteUniqueViewByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.UniqueView{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}
