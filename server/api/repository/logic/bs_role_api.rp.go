package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type RoleApiRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewRoleApiRepository(svcCtx *svc.RepositoryContext) *RoleApiRepository {
	return &RoleApiRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建RoleApi记录
func (s *RoleApiRepository) CreateRoleApi(ctx context.Context, roleApi *entity.RoleApi) (out *entity.RoleApi, err error) {
	db := s.DbEngin
	err = db.Create(&roleApi).Error
	if err != nil {
		return nil, err
	}
	return roleApi, err
}

// 删除RoleApi记录
func (s *RoleApiRepository) DeleteRoleApi(ctx context.Context, roleApi *entity.RoleApi) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&roleApi)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新RoleApi记录
func (s *RoleApiRepository) UpdateRoleApi(ctx context.Context, roleApi *entity.RoleApi) (out *entity.RoleApi, err error) {
	db := s.DbEngin
	err = db.Save(&roleApi).Error
	if err != nil {
		return nil, err
	}
	return roleApi, err
}

// 查询RoleApi记录
func (s *RoleApiRepository) FindRoleApi(ctx context.Context, id int) (out *entity.RoleApi, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除RoleApi记录
func (s *RoleApiRepository) DeleteRoleApiByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.RoleApi{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询RoleApi记录
func (s *RoleApiRepository) FindRoleApiList(ctx context.Context, page *request.PageQuery) (list []*entity.RoleApi, total int64, err error) {
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
