package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
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
func (s *RoleMenuRepository) CreateRoleMenu(ctx context.Context, roleMenu *entity.RoleMenu) (out *entity.RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Create(&roleMenu).Error
	if err != nil {
		return nil, err
	}
	return roleMenu, err
}

// 删除RoleMenu记录
func (s *RoleMenuRepository) DeleteRoleMenu(ctx context.Context, roleMenu *entity.RoleMenu) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Delete(&roleMenu)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新RoleMenu记录
func (s *RoleMenuRepository) UpdateRoleMenu(ctx context.Context, roleMenu *entity.RoleMenu) (out *entity.RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Save(&roleMenu).Error
	if err != nil {
		return nil, err
	}
	return roleMenu, err
}

// 查询RoleMenu记录
func (s *RoleMenuRepository) FindRoleMenu(ctx context.Context, id int) (out *entity.RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx)
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除RoleMenu记录
func (s *RoleMenuRepository) DeleteRoleMenuByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)
	query := db.Delete(&[]entity.RoleMenu{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询RoleMenu记录
func (s *RoleMenuRepository) FindRoleMenuList(ctx context.Context, page *request.PageQuery) (list []*entity.RoleMenu, total int64, err error) {
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
