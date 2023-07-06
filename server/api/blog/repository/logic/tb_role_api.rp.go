package logic

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
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
func (s *RoleApiRepository) CreateRoleApi(roleApi *entity.RoleApi) (out *entity.RoleApi, err error) {
	db := s.DbEngin
	err = db.Create(&roleApi).Error
	if err != nil {
		return nil, err
	}
	return roleApi, err
}

// 删除RoleApi记录
func (s *RoleApiRepository) DeleteRoleApi(roleApi *entity.RoleApi) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&roleApi)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新RoleApi记录
func (s *RoleApiRepository) UpdateRoleApi(roleApi *entity.RoleApi) (out *entity.RoleApi, err error) {
	db := s.DbEngin
	err = db.Save(&roleApi).Error
	if err != nil {
		return nil, err
	}
	return roleApi, err
}

// 根据id获取RoleApi记录
func (s *RoleApiRepository) FindRoleApi(id int) (out *entity.RoleApi, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除RoleApi记录
func (s *RoleApiRepository) DeleteRoleApiByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.RoleApi{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取RoleApi记录
func (s *RoleApiRepository) GetRoleApiList(page *request.PageInfo) (list []*entity.RoleApi, total int64, err error) {
	limit := page.Limit()
	offset := page.Offset()
	// 创建db
	db := s.DbEngin
	var roleApis []*entity.RoleApi
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
	}

	err = db.Model(&roleApis).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Find(&roleApis).Error
	if err != nil {
		return nil, 0, err
	}

	return roleApis, total, nil
}
