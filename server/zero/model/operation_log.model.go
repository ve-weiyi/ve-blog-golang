package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameOperationLog = "operation_log"

type (
	// 接口定义
	IOperationLogModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out IOperationLogModel)
		// 增删改查
		Create(ctx context.Context, in *OperationLog) (out *OperationLog, err error)
		Update(ctx context.Context, in *OperationLog) (out *OperationLog, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *OperationLog, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*OperationLog) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*OperationLog, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*OperationLog, err error)
	}

	// 接口实现
	defaultOperationLogModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	OperationLog struct {
		ID             int64  `json:"id"`              // 主键id
		UserID         int64  `json:"user_id"`         // 用户id
		Nickname       string `json:"nickname"`        // 用户昵称
		IpAddress      string `json:"ip_address"`      // 操作ip
		IpSource       string `json:"ip_source"`       // 操作地址
		OptModule      string `json:"opt_module"`      // 操作模块
		OptDesc        string `json:"opt_desc"`        // 操作描述
		RequestURL     string `json:"request_url"`     // 请求地址
		RequestMethod  string `json:"request_method"`  // 请求方式
		RequestHeader  string `json:"request_header"`  // 请求头参数
		RequestData    string `json:"request_data"`    // 请求参数
		ResponseData   string `json:"response_data"`   // 返回数据
		ResponseStatus int64  `json:"response_status"` // 响应状态码
		Cost           string `json:"cost"`            // 耗时（ms）
		CreatedAt      int64  `json:"created_at"`      // 创建时间
		UpdatedAt      int64  `json:"updated_at"`      // 更新时间
	}
)

func NewOperationLogModel(db *gorm.DB, cache *redis.Client) IOperationLogModel {
	return &defaultOperationLogModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameOperationLog,
	}
}

// 切换事务操作
func (s *defaultOperationLogModel) WithTransaction(tx *gorm.DB) (out IOperationLogModel) {
	return NewOperationLogModel(tx, s.CacheEngin)
}

// 创建OperationLog记录
func (s *defaultOperationLogModel) Create(ctx context.Context, in *OperationLog) (out *OperationLog, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新OperationLog记录
func (s *defaultOperationLogModel) Update(ctx context.Context, in *OperationLog) (out *OperationLog, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除OperationLog记录
func (s *defaultOperationLogModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&OperationLog{})
	return query.RowsAffected, query.Error
}

// 查询OperationLog记录
func (s *defaultOperationLogModel) First(ctx context.Context, conditions string, args ...interface{}) (out *OperationLog, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(OperationLog)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询OperationLog记录
func (s *defaultOperationLogModel) BatchCreate(ctx context.Context, in ...*OperationLog) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询OperationLog记录
func (s *defaultOperationLogModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&OperationLog{})
	return query.RowsAffected, query.Error
}

// 查询OperationLog总数
func (s *defaultOperationLogModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

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

// 查询OperationLog列表
func (s *defaultOperationLogModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*OperationLog, err error) {
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

// 分页查询OperationLog记录
func (s *defaultOperationLogModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*OperationLog, err error) {
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
