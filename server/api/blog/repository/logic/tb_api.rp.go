package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
)

type ApiRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewApiRepository(svcCtx *svc.RepositoryContext) *ApiRepository {
	return &ApiRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Api记录
func (s *ApiRepository) CreateApi(ctx context.Context, api *entity.Api) (out *entity.Api, err error) {
	db := s.DbEngin
	err = db.Create(&api).Error
	if err != nil {
		return nil, err
	}
	return api, err
}

// 删除Api记录
func (s *ApiRepository) DeleteApi(ctx context.Context, api *entity.Api) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&api)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Api记录
func (s *ApiRepository) UpdateApi(ctx context.Context, api *entity.Api) (out *entity.Api, err error) {
	db := s.DbEngin
	err = db.Save(&api).Error
	if err != nil {
		return nil, err
	}
	return api, err
}

// 查询Api记录
func (s *ApiRepository) GetApi(ctx context.Context, id int) (out *entity.Api, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Api记录
func (s *ApiRepository) DeleteApiByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Api{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Api记录
func (s *ApiRepository) FindApiList(ctx context.Context, page *request.PageInfo) (list []*entity.Api, total int64, err error) {
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
