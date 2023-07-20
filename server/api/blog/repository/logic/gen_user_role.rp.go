package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type UserRoleRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUserRoleRepository(svcCtx *svc.RepositoryContext) *UserRoleRepository {
	return &UserRoleRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建UserRole记录
func (s *UserRoleRepository) CreateUserRole(ctx context.Context, userRole *entity.UserRole) (out *entity.UserRole, err error) {
	db := s.DbEngin
	err = db.Create(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, err
}

// 删除UserRole记录
func (s *UserRoleRepository) DeleteUserRole(ctx context.Context, userRole *entity.UserRole) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&userRole)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新UserRole记录
func (s *UserRoleRepository) UpdateUserRole(ctx context.Context, userRole *entity.UserRole) (out *entity.UserRole, err error) {
	db := s.DbEngin
	err = db.Save(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, err
}

// 查询UserRole记录
func (s *UserRoleRepository) GetUserRole(ctx context.Context, id int) (out *entity.UserRole, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UserRole记录
func (s *UserRoleRepository) DeleteUserRoleByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.UserRole{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询UserRole记录
func (s *UserRoleRepository) FindUserRoleList(ctx context.Context, page *request.PageInfo) (list []*entity.UserRole, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Sorts) != 0 {
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
