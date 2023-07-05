package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type PhotoRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewPhotoRepository(svcCtx *svc.RepositoryContext) *PhotoRepository {
	return &PhotoRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Photo记录
func (s *PhotoRepository) CreatePhoto(ctx context.Context, photo *entity.Photo, conditions ...*request.Condition) (out *entity.Photo, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&photo).Error
	if err != nil {
		return nil, err
	}
	return photo, err
}

// 更新Photo记录
func (s *PhotoRepository) UpdatePhoto(ctx context.Context, photo *entity.Photo, conditions ...*request.Condition) (out *entity.Photo, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&photo).Error
	if err != nil {
		return nil, err
	}
	return photo, err
}

// 删除Photo记录
func (s *PhotoRepository) DeletePhoto(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Photo{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询Photo记录
func (s *PhotoRepository) FindPhoto(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.Photo, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Photo记录
func (s *PhotoRepository) DeletePhotoByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Photo{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询Photo记录
func (s *PhotoRepository) FindPhotoList(ctx context.Context, page *request.PageQuery, conditions ...*request.Condition) (list []*entity.Photo, total int64, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

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
