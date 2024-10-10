package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ TUploadRecordModel = (*defaultTUploadRecordModel)(nil)

type (
	// 接口定义
	TUploadRecordModel interface {
		TableName() string
		// 在事务中操作
		WithTransaction(tx *gorm.DB) (out TUploadRecordModel)
		// 插入
		Insert(ctx context.Context, in *TUploadRecord) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*TUploadRecord) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *TUploadRecord) (rows int64, err error)
		Save(ctx context.Context, in *TUploadRecord) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *TUploadRecord, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *TUploadRecord, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*TUploadRecord, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TUploadRecord, err error)
		// add extra method in here
	}

	// 表字段定义
	TUploadRecord struct {
		Id        int64     `json:"id" gorm:"column:id" `                 // id
		UserId    int64     `json:"user_id" gorm:"column:user_id" `       // 用户id
		Label     string    `json:"label" gorm:"column:label" `           // 标签
		FileName  string    `json:"file_name" gorm:"column:file_name" `   // 文件名称
		FileSize  int64     `json:"file_size" gorm:"column:file_size" `   // 文件大小
		FileMd5   string    `json:"file_md5" gorm:"column:file_md5" `     // 文件md5值
		FileUrl   string    `json:"file_url" gorm:"column:file_url" `     // 上传路径
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" ` // 创建时间
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" ` // 更新时间
	}

	// 接口实现
	defaultTUploadRecordModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewTUploadRecordModel(db *gorm.DB, cache *redis.Client) TUploadRecordModel {
	return &defaultTUploadRecordModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`t_upload_record`",
	}
}

func (m *defaultTUploadRecordModel) TableName() string {
	return m.table
}

// 在事务中操作
func (m *defaultTUploadRecordModel) WithTransaction(tx *gorm.DB) (out TUploadRecordModel) {
	return NewTUploadRecordModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultTUploadRecordModel) Insert(ctx context.Context, in *TUploadRecord) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultTUploadRecordModel) InsertBatch(ctx context.Context, in ...*TUploadRecord) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultTUploadRecordModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&TUploadRecord{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultTUploadRecordModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&TUploadRecord{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 保存记录（更新零值）
func (m *defaultTUploadRecordModel) Save(ctx context.Context, in *TUploadRecord) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultTUploadRecordModel) Update(ctx context.Context, in *TUploadRecord) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultTUploadRecordModel) FindOne(ctx context.Context, id int64) (out *TUploadRecord, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultTUploadRecordModel) First(ctx context.Context, conditions string, args ...interface{}) (out *TUploadRecord, err error) {
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
func (m *defaultTUploadRecordModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&TUploadRecord{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultTUploadRecordModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*TUploadRecord, err error) {
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
func (m *defaultTUploadRecordModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TUploadRecord, err error) {
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
