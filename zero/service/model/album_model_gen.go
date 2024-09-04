package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ AlbumModel = (*defaultAlbumModel)(nil)

type (
	// 接口定义
	AlbumModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out AlbumModel)
		// 插入
		Insert(ctx context.Context, in *Album) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*Album) (rows int64, err error)
		// 更新
		Save(ctx context.Context, in *Album) (rows int64, err error)
		Update(ctx context.Context, in *Album) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *Album, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Album, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Album, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Album, err error)
		// add extra method in here
	}

	// 表字段定义
	Album struct {
		Id         int64     `json:"id" gorm:"column:id" `                   // 主键
		AlbumName  string    `json:"album_name" gorm:"column:album_name" `   // 相册名
		AlbumDesc  string    `json:"album_desc" gorm:"column:album_desc" `   // 相册描述
		AlbumCover string    `json:"album_cover" gorm:"column:album_cover" ` // 相册封面
		IsDelete   int64     `json:"is_delete" gorm:"column:is_delete" `     // 是否删除
		Status     int64     `json:"status" gorm:"column:status" `           // 状态值 1公开 2私密
		CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" `   // 创建时间
		UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" `   // 更新时间
	}

	// 接口实现
	defaultAlbumModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewAlbumModel(db *gorm.DB, cache *redis.Client) AlbumModel {
	return &defaultAlbumModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`album`",
	}
}

// 切换事务操作
func (m *defaultAlbumModel) WithTransaction(tx *gorm.DB) (out AlbumModel) {
	return NewAlbumModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultAlbumModel) Insert(ctx context.Context, in *Album) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultAlbumModel) InsertBatch(ctx context.Context, in ...*Album) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultAlbumModel) Save(ctx context.Context, in *Album) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultAlbumModel) Update(ctx context.Context, in *Album) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultAlbumModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&Album{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultAlbumModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&Album{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultAlbumModel) FindOne(ctx context.Context, id int64) (out *Album, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultAlbumModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Album, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询总数
func (m *defaultAlbumModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Album{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultAlbumModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Album, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

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

// 分页查询记录
func (m *defaultAlbumModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*Album, err error) {
	// 插入db
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有搜索条件
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sorts)
	}

	// 如果有分页参数
	if page > 0 && size > 0 {
		limit := size
		offset := (page - 1) * limit
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// add extra method in here
