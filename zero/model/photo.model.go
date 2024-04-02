package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNamePhoto = "photo"

type (
	// 接口定义
	PhotoModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out PhotoModel)
		// 增删改查
		Create(ctx context.Context, in *Photo) (out *Photo, err error)
		Update(ctx context.Context, in *Photo) (out *Photo, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Photo, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Photo) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Photo, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Photo, err error)
	}

	// 接口实现
	defaultPhotoModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Photo struct {
		Id        int64     `json:"id"`         // 主键
		AlbumId   int64     `json:"album_id"`   // 相册id
		PhotoName string    `json:"photo_name"` // 照片名
		PhotoDesc string    `json:"photo_desc"` // 照片描述
		PhotoSrc  string    `json:"photo_src"`  // 照片地址
		IsDelete  int64     `json:"is_delete"`  // 是否删除
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewPhotoModel(db *gorm.DB, cache *redis.Client) PhotoModel {
	return &defaultPhotoModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNamePhoto,
	}
}

// 切换事务操作
func (s *defaultPhotoModel) WithTransaction(tx *gorm.DB) (out PhotoModel) {
	return NewPhotoModel(tx, s.CacheEngin)
}

// 创建Photo记录
func (s *defaultPhotoModel) Create(ctx context.Context, in *Photo) (out *Photo, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Photo记录
func (s *defaultPhotoModel) Update(ctx context.Context, in *Photo) (out *Photo, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Photo记录
func (s *defaultPhotoModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Photo{})
	return query.RowsAffected, query.Error
}

// 查询Photo记录
func (s *defaultPhotoModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Photo, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Photo)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Photo记录
func (s *defaultPhotoModel) BatchCreate(ctx context.Context, in ...*Photo) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Photo记录
func (s *defaultPhotoModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Photo{})
	return query.RowsAffected, query.Error
}

// 查询Photo总数
func (s *defaultPhotoModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Photo{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Photo列表
func (s *defaultPhotoModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Photo, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Find(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 分页查询Photo记录
func (s *defaultPhotoModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Photo, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有搜索条件
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sorts)
	}

	// 如果有分页参数
	if limit > 0 && offset > 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
