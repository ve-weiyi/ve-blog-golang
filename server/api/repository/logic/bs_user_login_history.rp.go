package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type UserLoginHistoryRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUserLoginHistoryRepository(svcCtx *svc.RepositoryContext) *UserLoginHistoryRepository {
	return &UserLoginHistoryRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建UserLoginHistory记录
func (s *UserLoginHistoryRepository) CreateUserLoginHistory(ctx context.Context, userLoginHistory *entity.UserLoginHistory, conditions ...*request.Condition) (out *entity.UserLoginHistory, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&userLoginHistory).Error
	if err != nil {
		return nil, err
	}
	return userLoginHistory, err
}

// 更新UserLoginHistory记录
func (s *UserLoginHistoryRepository) UpdateUserLoginHistory(ctx context.Context, userLoginHistory *entity.UserLoginHistory, conditions ...*request.Condition) (out *entity.UserLoginHistory, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&userLoginHistory).Error
	if err != nil {
		return nil, err
	}
	return userLoginHistory, err
}

// 删除UserLoginHistory记录
func (s *UserLoginHistoryRepository) DeleteUserLoginHistory(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.UserLoginHistory{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询UserLoginHistory记录
func (s *UserLoginHistoryRepository) FindUserLoginHistory(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.UserLoginHistory, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UserLoginHistory记录
func (s *UserLoginHistoryRepository) DeleteUserLoginHistoryByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.UserLoginHistory{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询UserLoginHistory记录
func (s *UserLoginHistoryRepository) FindUserLoginHistoryList(ctx context.Context, page *request.PageQuery) (list []*entity.UserLoginHistory, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Sorts) != 0 {
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
		return nil, err
	}

	return list, nil
}

// 查询总数
func (s *UserLoginHistoryRepository) Count(ctx context.Context, conditions ...*request.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.UserLoginHistory{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
