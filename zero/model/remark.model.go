package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameRemark = "remark"

type (
	// 接口定义
	RemarkModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out RemarkModel)
		// 增删改查
		Create(ctx context.Context, in *Remark) (out *Remark, err error)
		Update(ctx context.Context, in *Remark) (out *Remark, err error)
		Delete(ctx context.Context, id int64) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *Remark, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*Remark) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*Remark, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Remark, err error)
	}

	// 接口实现
	defaultRemarkModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	Remark struct {
		Id             int64     `json:"id"`              // 主键id
		Nickname       string    `json:"nickname"`        // 昵称
		Avatar         string    `json:"avatar"`          // 头像
		MessageContent string    `json:"message_content"` // 留言内容
		IpAddress      string    `json:"ip_address"`      // 用户ip
		IpSource       string    `json:"ip_source"`       // 用户地址
		Time           int64     `json:"time"`            // 弹幕速度
		IsReview       int64     `json:"is_review"`       // 是否审核
		CreatedAt      time.Time `json:"created_at"`      // 发布时间
		UpdatedAt      time.Time `json:"updated_at"`      // 更新时间
	}
)

func NewRemarkModel(db *gorm.DB, cache *redis.Client) RemarkModel {
	return &defaultRemarkModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameRemark,
	}
}

// 切换事务操作
func (s *defaultRemarkModel) WithTransaction(tx *gorm.DB) (out RemarkModel) {
	return NewRemarkModel(tx, s.CacheEngin)
}

// 创建Remark记录
func (s *defaultRemarkModel) Create(ctx context.Context, in *Remark) (out *Remark, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新Remark记录
func (s *defaultRemarkModel) Update(ctx context.Context, in *Remark) (out *Remark, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除Remark记录
func (s *defaultRemarkModel) Delete(ctx context.Context, id int64) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&Remark{})
	return query.RowsAffected, query.Error
}

// 查询Remark记录
func (s *defaultRemarkModel) First(ctx context.Context, conditions string, args ...interface{}) (out *Remark, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(Remark)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询Remark记录
func (s *defaultRemarkModel) BatchCreate(ctx context.Context, in ...*Remark) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询Remark记录
func (s *defaultRemarkModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&Remark{})
	return query.RowsAffected, query.Error
}

// 查询Remark总数
func (s *defaultRemarkModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&Remark{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询Remark列表
func (s *defaultRemarkModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*Remark, err error) {
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

// 分页查询Remark记录
func (s *defaultRemarkModel) FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*Remark, err error) {
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
