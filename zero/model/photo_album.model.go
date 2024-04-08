package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNamePhotoAlbum = "photo_album"

type (
	// 接口定义
	PhotoAlbumModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out PhotoAlbumModel)
		// 增删改查
		Create(ctx context.Context, in *PhotoAlbum) (out *PhotoAlbum, err error)
		Update(ctx context.Context, in *PhotoAlbum) (out *PhotoAlbum, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *PhotoAlbum, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*PhotoAlbum) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*PhotoAlbum, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*PhotoAlbum, err error)
	}

	// 接口实现
	defaultPhotoAlbumModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	PhotoAlbum struct {
		Id         int64     `json:"id"`          // 主键
		AlbumName  string    `json:"album_name"`  // 相册名
		AlbumDesc  string    `json:"album_desc"`  // 相册描述
		AlbumCover string    `json:"album_cover"` // 相册封面
		IsDelete   int64     `json:"is_delete"`   // 是否删除
		Status     int64     `json:"status"`      // 状态值 1公开 2私密
		CreatedAt  time.Time `json:"created_at"`  // 创建时间
		UpdatedAt  time.Time `json:"updated_at"`  // 更新时间
	}
)

func NewPhotoAlbumModel(db *gorm.DB, cache *redis.Client) PhotoAlbumModel {
	return &defaultPhotoAlbumModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNamePhotoAlbum,
	}
}

// 切换事务操作
func (s *defaultPhotoAlbumModel) WithTransaction(tx *gorm.DB) (out PhotoAlbumModel) {
	return NewPhotoAlbumModel(tx, s.CacheEngin)
}

// 创建PhotoAlbum记录
func (s *defaultPhotoAlbumModel) Create(ctx context.Context, in *PhotoAlbum) (out *PhotoAlbum, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新PhotoAlbum记录
func (s *defaultPhotoAlbumModel) Update(ctx context.Context, in *PhotoAlbum) (out *PhotoAlbum, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除PhotoAlbum记录
func (s *defaultPhotoAlbumModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&PhotoAlbum{})
	return query.RowsAffected, query.Error
}

// 查询PhotoAlbum记录
func (s *defaultPhotoAlbumModel) First(ctx context.Context, conditions string, args ...interface{}) (out *PhotoAlbum, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(PhotoAlbum)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询PhotoAlbum记录
func (s *defaultPhotoAlbumModel) BatchCreate(ctx context.Context, in ...*PhotoAlbum) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询PhotoAlbum记录
func (s *defaultPhotoAlbumModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&PhotoAlbum{})
	return query.RowsAffected, query.Error
}

// 查询PhotoAlbum总数
func (s *defaultPhotoAlbumModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&PhotoAlbum{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询PhotoAlbum列表
func (s *defaultPhotoAlbumModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*PhotoAlbum, err error) {
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

// 分页查询PhotoAlbum记录
func (s *defaultPhotoAlbumModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*PhotoAlbum, err error) {
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
