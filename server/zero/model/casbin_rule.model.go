package model

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const TableNameCasbinRule = "casbin_rule"

type (
	// 接口定义
	ICasbinRuleModel interface {
		// 切换事务操作
		WithTransaction(tx *gorm.DB) (out ICasbinRuleModel)
		// 增删改查
		Create(ctx context.Context, in *CasbinRule) (out *CasbinRule, err error)
		Update(ctx context.Context, in *CasbinRule) (out *CasbinRule, err error)
		Delete(ctx context.Context, id int) (rows int64, err error)
		First(ctx context.Context, conditions string, args ...interface{}) (out *CasbinRule, err error)
		// 批量操作
		BatchCreate(ctx context.Context, in ...*CasbinRule) (rows int64, err error)
		BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
		// 查询
		Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
		FindALL(ctx context.Context, conditions string, args ...interface{}) (list []*CasbinRule, err error)
		FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []*CasbinRule, err error)
	}

	// 接口实现
	defaultCasbinRuleModel struct {
		DbEngin    *gorm.DB
		CacheEngin *redis.Client
		tableName  string
	}

	// 表字段定义
	CasbinRule struct {
		ID    int64  `json:"id"`
		Ptype string `json:"ptype"`
		V0    string `json:"v0"`
		V1    string `json:"v1"`
		V2    string `json:"v2"`
		V3    string `json:"v3"`
		V4    string `json:"v4"`
		V5    string `json:"v5"`
	}
)

func NewCasbinRuleModel(db *gorm.DB, cache *redis.Client) ICasbinRuleModel {
	return &defaultCasbinRuleModel{
		DbEngin:    db,
		CacheEngin: cache,
		tableName:  TableNameCasbinRule,
	}
}

// 切换事务操作
func (s *defaultCasbinRuleModel) WithTransaction(tx *gorm.DB) (out ICasbinRuleModel) {
	return NewCasbinRuleModel(tx, s.CacheEngin)
}

// 创建CasbinRule记录
func (s *defaultCasbinRuleModel) Create(ctx context.Context, in *CasbinRule) (out *CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Create(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 更新CasbinRule记录
func (s *defaultCasbinRuleModel) Update(ctx context.Context, in *CasbinRule) (out *CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.Save(&in).Error
	if err != nil {
		return nil, err
	}
	return in, err
}

// 删除CasbinRule记录
func (s *defaultCasbinRuleModel) Delete(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	db = db.Where("id = ?", id)

	query := db.Delete(&CasbinRule{})
	return query.RowsAffected, query.Error
}

// 查询CasbinRule记录
func (s *defaultCasbinRuleModel) First(ctx context.Context, conditions string, args ...interface{}) (out *CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	out = new(CasbinRule)
	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 查询CasbinRule记录
func (s *defaultCasbinRuleModel) BatchCreate(ctx context.Context, in ...*CasbinRule) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	err = db.CreateInBatches(&in, len(in)).Error
	if err != nil {
		return 0, err
	}
	return rows, err
}

// 查询CasbinRule记录
func (s *defaultCasbinRuleModel) BatchDelete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&CasbinRule{})
	return query.RowsAffected, query.Error
}

// 查询CasbinRule总数
func (s *defaultCasbinRuleModel) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx).Table(s.tableName)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&CasbinRule{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询CasbinRule列表
func (s *defaultCasbinRuleModel) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*CasbinRule, err error) {
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

// 分页查询CasbinRule记录
func (s *defaultCasbinRuleModel) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*CasbinRule, err error) {
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
