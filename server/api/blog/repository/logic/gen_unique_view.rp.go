package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type UniqueViewRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUniqueViewRepository(svcCtx *svc.RepositoryContext) *UniqueViewRepository {
	return &UniqueViewRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建UniqueView记录
func (s *UniqueViewRepository) CreateUniqueView(ctx context.Context, uniqueView *entity.UniqueView) (out *entity.UniqueView, err error) {
	db := s.DbEngin
	err = db.Create(&uniqueView).Error
	if err != nil {
		return nil, err
	}
	return uniqueView, err
}

// 删除UniqueView记录
func (s *UniqueViewRepository) DeleteUniqueView(ctx context.Context, uniqueView *entity.UniqueView) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&uniqueView)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新UniqueView记录
func (s *UniqueViewRepository) UpdateUniqueView(ctx context.Context, uniqueView *entity.UniqueView) (out *entity.UniqueView, err error) {
	db := s.DbEngin
	err = db.Save(&uniqueView).Error
	if err != nil {
		return nil, err
	}
	return uniqueView, err
}

// 查询UniqueView记录
func (s *UniqueViewRepository) GetUniqueView(ctx context.Context, id int) (out *entity.UniqueView, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UniqueView记录
func (s *UniqueViewRepository) DeleteUniqueViewByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.UniqueView{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询UniqueView记录
func (s *UniqueViewRepository) FindUniqueViewList(ctx context.Context, page *request.PageInfo) (list []*entity.UniqueView, total int64, err error) {
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
