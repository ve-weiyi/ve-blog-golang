package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
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
func (s *UserLoginHistoryRepository) CreateUserLoginHistory(ctx context.Context, userLoginHistory *entity.UserLoginHistory) (out *entity.UserLoginHistory, err error) {
	db := s.DbEngin
	err = db.Create(&userLoginHistory).Error
	if err != nil {
		return nil, err
	}
	return userLoginHistory, err
}

// 删除UserLoginHistory记录
func (s *UserLoginHistoryRepository) DeleteUserLoginHistory(ctx context.Context, userLoginHistory *entity.UserLoginHistory) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&userLoginHistory)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新UserLoginHistory记录
func (s *UserLoginHistoryRepository) UpdateUserLoginHistory(ctx context.Context, userLoginHistory *entity.UserLoginHistory) (out *entity.UserLoginHistory, err error) {
	db := s.DbEngin
	err = db.Save(&userLoginHistory).Error
	if err != nil {
		return nil, err
	}
	return userLoginHistory, err
}

// 查询UserLoginHistory记录
func (s *UserLoginHistoryRepository) GetUserLoginHistory(ctx context.Context, id int) (out *entity.UserLoginHistory, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UserLoginHistory记录
func (s *UserLoginHistoryRepository) DeleteUserLoginHistoryByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.UserLoginHistory{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询UserLoginHistory记录
func (s *UserLoginHistoryRepository) FindUserLoginHistoryList(ctx context.Context, page *request.PageInfo) (list []*entity.UserLoginHistory, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Orders) != 0 {
		db = db.Order(page.OrderClause())
	}

	// 查询总数,要在使用limit之前
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
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

	return list, total, nil
}
