package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type CasbinRuleRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewCasbinRuleRepository(svcCtx *svc.RepositoryContext) *CasbinRuleRepository {
	return &CasbinRuleRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建CasbinRule记录
func (s *CasbinRuleRepository) CreateCasbinRule(ctx context.Context, casbinRule *entity.CasbinRule, conditions ...*request.Condition) (out *entity.CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&casbinRule).Error
	if err != nil {
		return nil, err
	}
	return casbinRule, err
}

// 更新CasbinRule记录
func (s *CasbinRuleRepository) UpdateCasbinRule(ctx context.Context, casbinRule *entity.CasbinRule, conditions ...*request.Condition) (out *entity.CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&casbinRule).Error
	if err != nil {
		return nil, err
	}
	return casbinRule, err
}

// 删除CasbinRule记录
func (s *CasbinRuleRepository) DeleteCasbinRule(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.CasbinRule{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询CasbinRule记录
func (s *CasbinRuleRepository) FindCasbinRule(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.CasbinRule, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除CasbinRule记录
func (s *CasbinRuleRepository) DeleteCasbinRuleByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.CasbinRule{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询CasbinRule记录
func (s *CasbinRuleRepository) FindCasbinRuleList(ctx context.Context, page *request.PageQuery, conditions ...*request.Condition) (list []*entity.CasbinRule, total int64, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Sorts) != 0 {
		db = db.Order(page.OrderClause())
	}

	// 查询总数,要在使用limit之前
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果有分页参数
	if page.Page != 0 || page.PageSize != 0 {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
