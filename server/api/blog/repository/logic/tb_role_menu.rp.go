package logic

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
)

type RoleMenuRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewRoleMenuRepository(svcCtx *svc.RepositoryContext) *RoleMenuRepository {
	return &RoleMenuRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建RoleMenu记录
func (s *RoleMenuRepository) CreateRoleMenu(roleMenu *entity.RoleMenu) (out *entity.RoleMenu, err error) {
	db := s.DbEngin
	err = db.Create(&roleMenu).Error
	if err != nil {
		return nil, err
	}
	return roleMenu, err
}

// 删除RoleMenu记录
func (s *RoleMenuRepository) DeleteRoleMenu(roleMenu *entity.RoleMenu) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&roleMenu)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新RoleMenu记录
func (s *RoleMenuRepository) UpdateRoleMenu(roleMenu *entity.RoleMenu) (out *entity.RoleMenu, err error) {
	db := s.DbEngin
	err = db.Save(&roleMenu).Error
	if err != nil {
		return nil, err
	}
	return roleMenu, err
}

// 根据id获取RoleMenu记录
func (s *RoleMenuRepository) FindRoleMenu(id int) (out *entity.RoleMenu, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除RoleMenu记录
func (s *RoleMenuRepository) DeleteRoleMenuByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.RoleMenu{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取RoleMenu记录
func (s *RoleMenuRepository) GetRoleMenuList(page *request.PageInfo) (list []*entity.RoleMenu, total int64, err error) {
	limit := page.Limit()
	offset := page.Offset()
	// 创建db
	db := s.DbEngin
	var roleMenus []*entity.RoleMenu
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
	}

	err = db.Model(&roleMenus).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Find(&roleMenus).Error
	if err != nil {
		return nil, 0, err
	}

	return roleMenus, total, nil
}
