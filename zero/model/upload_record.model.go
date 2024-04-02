package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameUploadRecord = "upload_record"

type (
	// 接口定义
	UploadRecordModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out UploadRecordModel)
		// 增删改查
		Create(ctx context.Context, in *UploadRecord) (out *UploadRecord, err error)
		Update(ctx context.Context, in *UploadRecord) (out *UploadRecord, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *UploadRecord, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*UploadRecord) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*UploadRecord, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UploadRecord, err error)
	}

	// 接口实现
	defaultUploadRecordModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	UploadRecord struct {
		Id        int64     `json:"id"`         // id
		UserId    int64     `json:"user_id"`    // 用户id
		Label     string    `json:"label"`      // 标签
		FileName  string    `json:"file_name"`  // 文件名称
		FileSize  int64     `json:"file_size"`  // 文件大小
		FileMd5   string    `json:"file_md5"`   // 文件md5值
		FileUrl   string    `json:"file_url"`   // 上传路径
		CreatedAt time.Time `json:"created_at"` // 创建时间
		UpdatedAt time.Time `json:"updated_at"` // 更新时间
	}
)

func NewUploadRecordModel(db *gorm.DB, cache *redis.Client) UploadRecordModel {
	return &defaultUploadRecordModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameUploadRecord,
	}
}

// 切换事务操作
func (s *defaultUploadRecordModel) WithTransaction(tx *gorm.DB) (out UploadRecordModel) {
	return NewUploadRecordModel(tx, s.CacheEngin)
}

// 创建UploadRecord记录
func (s *defaultUploadRecordModel) Create(ctx context.Context, in *UploadRecord) (out *UploadRecord, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新UploadRecord记录
func (s *defaultUploadRecordModel) Update(ctx context.Context, in *UploadRecord) (out *UploadRecord, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除UploadRecord记录
func (s *defaultUploadRecordModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&UploadRecord{})
	return query.RowsAffected, query.Error
}

// 查询UploadRecord记录
func (s *defaultUploadRecordModel) First(ctx context.Context, conditions string, args ...interface{}) (out *UploadRecord, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(UploadRecord)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询UploadRecord记录
func (s *defaultUploadRecordModel) BatchCreate(ctx context.Context, in ...*UploadRecord) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询UploadRecord记录
func (s *defaultUploadRecordModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&UploadRecord{})
	return query.RowsAffected, query.Error
}

// 查询UploadRecord总数
func (s *defaultUploadRecordModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&UploadRecord{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询UploadRecord列表
func (s *defaultUploadRecordModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*UploadRecord, err error) {
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

// 分页查询UploadRecord记录
func (s *defaultUploadRecordModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*UploadRecord, err error) {
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
