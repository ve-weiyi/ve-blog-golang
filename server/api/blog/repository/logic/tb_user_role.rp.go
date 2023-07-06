package logic

import (
	"fmt"

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
func (s *UserRoleRepository) CreateUserRole(userRole *entity.UserRole) (out *entity.UserRole, err error) {
	db := s.DbEngin
	err = db.Create(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, err
}

// 删除UserRole记录
func (s *UserRoleRepository) DeleteUserRole(userRole *entity.UserRole) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&userRole)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新UserRole记录
func (s *UserRoleRepository) UpdateUserRole(userRole *entity.UserRole) (out *entity.UserRole, err error) {
	db := s.DbEngin
	err = db.Save(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, err
}

// 根据id获取UserRole记录
func (s *UserRoleRepository) FindUserRole(id int) (out *entity.UserRole, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UserRole记录
func (s *UserRoleRepository) DeleteUserRoleByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.UserRole{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取UserRole记录
func (s *UserRoleRepository) GetUserRoleList(page *request.PageInfo) (list []*entity.UserRole, total int64, err error) {
	limit := page.Limit()
	offset := page.Offset()
	// 创建db
	db := s.DbEngin
	var userRoles []*entity.UserRole
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
	}

	err = db.Model(&userRoles).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Find(&userRoles).Error
	if err != nil {
		return nil, 0, err
	}

	return userRoles, total, nil
}
