package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type RoleRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewRoleRepository(svcCtx *svc.RepositoryContext) *RoleRepository {
	return &RoleRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Role记录
func (s *RoleRepository) CreateRole(ctx context.Context, role *entity.Role) (out *entity.Role, err error) {
	db := s.DbEngin
	err = db.Create(&role).Error
	if err != nil {
		return nil, err
	}
	return role, err
}

// 删除Role记录
func (s *RoleRepository) DeleteRole(ctx context.Context, role *entity.Role) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&role)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Role记录
func (s *RoleRepository) UpdateRole(ctx context.Context, role *entity.Role) (out *entity.Role, err error) {
	db := s.DbEngin
	err = db.Save(&role).Error
	if err != nil {
		return nil, err
	}
	return role, err
}

// 查询Role记录
func (s *RoleRepository) GetRole(ctx context.Context, id int) (out *entity.Role, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Role记录
func (s *RoleRepository) DeleteRoleByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Role{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Role记录
func (s *RoleRepository) FindRoleList(ctx context.Context, page *request.PageInfo) (list []*entity.Role, total int64, err error) {
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
