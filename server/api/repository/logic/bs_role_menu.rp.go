package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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

// 更新RoleMenu记录
func (s *RoleMenuRepository) UpdateRoleMenu(ctx context.Context, roleMenu *entity.RoleMenu) (out *entity.RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&roleMenu).Error
	if err != nil {
		return nil, err
	}
	return roleMenu, err
}

// 删除RoleMenu记录
func (s *RoleMenuRepository) DeleteRoleMenu(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.RoleMenu{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询RoleMenu记录
func (s *RoleMenuRepository) FindRoleMenu(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.RoleMenu, err error) {
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

// 分页查询RoleMenu记录
func (s *RoleMenuRepository) FindRoleMenuList(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.RoleMenu, err error) {
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
func (s *RoleMenuRepository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.RoleMenu{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询RoleMenu记录——根据id
func (s *RoleMenuRepository) FindRoleMenuById(ctx context.Context, id int) (out *entity.RoleMenu, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 删除RoleMenu记录——根据id
func (s *RoleMenuRepository) DeleteRoleMenuById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.RoleMenu{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除RoleMenu记录——根据ids
func (s *RoleMenuRepository) DeleteRoleMenuByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.RoleMenu{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}