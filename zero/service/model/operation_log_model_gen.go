package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ OperationLogModel = (*defaultOperationLogModel)(nil)

type (
	// 接口定义
	OperationLogModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out OperationLogModel)
		// 插入
		Insert(ctx context.Context, in *OperationLog) (rows int64, err error)
		InsertBatch(ctx context.Context, in ...*OperationLog) (rows int64, err error)
		// 更新
		Update(ctx context.Context, in *OperationLog) (rows int64, err error)
		UpdateNotEmpty(ctx context.Context, in *OperationLog) (rows int64, err error)
		// 删除
		Delete(ctx context.Context, id int64) (rows int64, err error)
		DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		FindOne(ctx context.Context, id int64) (out *OperationLog, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *OperationLog, err error)
		FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*OperationLog, err error)
		FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*OperationLog, err error)
		// add extra method in here
	}

	// 表字段定义
	OperationLog struct {
		Id             int64     `json:"id" gorm:"column:id" `                           // 主键id
		UserId         int64     `json:"user_id" gorm:"column:user_id" `                 // 用户id
		Nickname       string    `json:"nickname" gorm:"column:nickname" `               // 用户昵称
		IpAddress      string    `json:"ip_address" gorm:"column:ip_address" `           // 操作ip
		IpSource       string    `json:"ip_source" gorm:"column:ip_source" `             // 操作地址
		OptModule      string    `json:"opt_module" gorm:"column:opt_module" `           // 操作模块
		OptDesc        string    `json:"opt_desc" gorm:"column:opt_desc" `               // 操作描述
		RequestUrl     string    `json:"request_url" gorm:"column:request_url" `         // 请求地址
		RequestMethod  string    `json:"request_method" gorm:"column:request_method" `   // 请求方式
		RequestHeader  string    `json:"request_header" gorm:"column:request_header" `   // 请求头参数
		RequestData    string    `json:"request_data" gorm:"column:request_data" `       // 请求参数
		ResponseData   string    `json:"response_data" gorm:"column:response_data" `     // 返回数据
		ResponseStatus int64     `json:"response_status" gorm:"column:response_status" ` // 响应状态码
		Cost           string    `json:"cost" gorm:"column:cost" `                       // 耗时（ms）
		CreatedAt      time.Time `json:"created_at" gorm:"column:created_at" `           // 创建时间
		UpdatedAt      time.Time `json:"updated_at" gorm:"column:updated_at" `           // 更新时间
	}

	// 接口实现
	defaultOperationLogModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		table      string
	}
)

func NewOperationLogModel(db *gorm.DB, cache *redis.Client) OperationLogModel {
	return &defaultOperationLogModel{
		DbEngin:    db,
		CacheEngin: cache,
		table:      "`operation_log`",
	}
}

// 切换事务操作
func (m *defaultOperationLogModel) WithTransaction(tx *gorm.DB) (out OperationLogModel) {
	return NewOperationLogModel(tx, m.CacheEngin)
}

// 插入记录 (返回的是受影响行数，如需获取自增id，请通过data参数获取)
func (m *defaultOperationLogModel) Insert(ctx context.Context, in *OperationLog) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Create(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 插入记录
func (m *defaultOperationLogModel) InsertBatch(ctx context.Context, in ...*OperationLog) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.CreateInBatches(&in, len(in))
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（不更新零值）
func (m *defaultOperationLogModel) Update(ctx context.Context, in *OperationLog) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Save(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 更新记录（更新零值）
func (m *defaultOperationLogModel) UpdateNotEmpty(ctx context.Context, in *OperationLog) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	result := db.Updates(&in)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 删除记录
func (m *defaultOperationLogModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	db = db.Where("id = ?", id)

	result := db.Delete(&OperationLog{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultOperationLogModel) DeleteBatch(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	result := db.Delete(&OperationLog{})
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, err
}

// 查询记录
func (m *defaultOperationLogModel) FindOne(ctx context.Context, id int64) (out *OperationLog, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	err = db.Where("`id` = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}

	return out, err
}

// 查询记录
func (m *defaultOperationLogModel) First(ctx context.Context, conditions string, args ...interface{}) (out *OperationLog, err error) {
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
func (m *defaultOperationLogModel) FindCount(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := m.DbEngin.WithContext(ctx).Table(m.table)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&OperationLog{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (m *defaultOperationLogModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*OperationLog, err error) {
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
func (m *defaultOperationLogModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*OperationLog, err error) {
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
