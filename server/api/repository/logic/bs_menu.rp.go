package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type MenuRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewMenuRepository(svcCtx *svc.RepositoryContext) *MenuRepository {
	return &MenuRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Menu记录
func (s *MenuRepository) CreateMenu(ctx context.Context, menu *entity.Menu) (out *entity.Menu, err error) {
	db := s.DbEngin
	err = db.Create(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, err
}

// 更新Menu记录
func (s *MenuRepository) UpdateMenu(ctx context.Context, menu *entity.Menu) (out *entity.Menu, err error) {
	db := s.DbEngin
	err = db.Save(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, err
}

// 删除Menu记录
func (s *MenuRepository) DeleteMenu(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&entity.Menu{}, "id = ?", id)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 查询Menu记录
func (s *MenuRepository) FindMenu(ctx context.Context, id int) (out *entity.Menu, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Menu记录
func (s *MenuRepository) DeleteMenuByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&entity.Menu{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Menu记录
func (s *MenuRepository) FindMenuList(ctx context.Context, page *request.PageQuery) (list []*entity.Menu, total int64, err error) {
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
