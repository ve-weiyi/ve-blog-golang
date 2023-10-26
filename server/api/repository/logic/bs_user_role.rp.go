package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, err
}

// 更新UserRole记录
func (s *UserRoleRepository) UpdateUserRole(ctx context.Context, userRole *entity.UserRole) (out *entity.UserRole, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, err
}

// 删除UserRole记录
func (s *UserRoleRepository) DeleteUserRole(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.UserRole{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询UserRole记录
func (s *UserRoleRepository) FindUserRole(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.UserRole, err error) {
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

// 分页查询UserRole记录
func (s *UserRoleRepository) FindUserRoleList(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.UserRole, err error) {
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
func (s *UserRoleRepository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.UserRole{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UserRole记录——根据id
func (s *UserRoleRepository) FindUserRoleById(ctx context.Context, id int) (out *entity.UserRole, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 删除UserRole记录——根据id
func (s *UserRoleRepository) DeleteUserRoleById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.UserRole{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除UserRole记录——根据ids
func (s *UserRoleRepository) DeleteUserRoleByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.UserRole{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}
