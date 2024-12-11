package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ TFileUploadModel = (*defaultTFileUploadModel)(nil)

type (
	// 接口定义
	TFileUploadModel interface {
		TableName() string
		// 在事务中操作
		WithTransaction(tx *gorm.DB) (out TFileUploadModel)
		// 插入
		Insert(ctx context.Context, in *TFileUpload) (rows int64, err error)
		Inserts(ctx context.Context, in ...*TFileUpload) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *TFileUpload) (rows int64, err error)
		Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (rows int64, err error)
		// 保存
		Save(ctx context.Context, in *TFileUpload) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *TFileUpload, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*TFileUpload, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TFileUpload, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		// add extra method in here
	}

	// 表字段定义
	TFileUpload struct {
		Id        int64     `json:"id" gorm:"column:id" `                 // id
		UserId    string    `json:"user_id" gorm:"column:user_id" `       // 用户id
		FilePath  string    `json:"file_path" gorm:"column:file_path" `   // 文件路径
		FileName  string    `json:"file_name" gorm:"column:file_name" `   // 文件名称
		FileType  string    `json:"file_type" gorm:"column:file_type" `   // 文件类型
		FileSize  int64     `json:"file_size" gorm:"column:file_size" `   // 文件大小
		FileMd5   string    `json:"file_md5" gorm:"column:file_md5" `     // 文件md5值
		FileUrl   string    `json:"file_url" gorm:"column:file_url" `     // 上传路径
		CreatedAt time.Time `json:"created_at" gorm:"column:created_at" ` // 创建时间
		UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" ` // 更新时间
	}

	// 接口实现
	defaultTFileUploadModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewTFileUploadModel(db *gorm.DB, cache *redis.Client) TFileUploadModel {
	return &defaultTFileUploadModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`t_file_upload`",
	}
}

func (m *defaultTFileUploadModel) TableName() string {
	return m.table
}

// 在事务中操作
func (m *defaultTFileUploadModel) WithTransaction(tx *gorm.DB) (out TFileUploadModel) {
	return NewTFileUploadModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultTFileUploadModel) Insert(ctx context.Context, in *TFileUpload) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录（批量操作）
func (m *defaultTFileUploadModel) Inserts(ctx context.Context, in ...*TFileUpload) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultTFileUploadModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&TFileUpload{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录（批量操作）
func (m *defaultTFileUploadModel) Deletes(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&TFileUpload{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultTFileUploadModel) Update(ctx context.Context, in *TFileUpload) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（批量操作）
func (m *defaultTFileUploadModel) Updates(ctx context.Context, columns map[string]interface{}, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Where(conditions, args...).Updates(columns)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 保存记录（更新零值）
func (m *defaultTFileUploadModel) Save(ctx context.Context, in *TFileUpload) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Omit("created_at").Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultTFileUploadModel) FindOne(ctx context.Context, id int64) (out *TFileUpload, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询列表
func (m *defaultTFileUploadModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*TFileUpload, err error) {
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
func (m *defaultTFileUploadModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*TFileUpload, err error) {
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

// 查询总数
func (m *defaultTFileUploadModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&TFileUpload{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// add extra method in here
