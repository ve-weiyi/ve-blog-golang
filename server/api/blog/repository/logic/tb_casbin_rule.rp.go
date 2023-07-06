package logic

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
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
func (s *CasbinRuleRepository) CreateCasbinRule(casbinRule *entity.CasbinRule) (out *entity.CasbinRule, err error) {
	db := s.DbEngin
	err = db.Create(&casbinRule).Error
	if err != nil {
		return nil, err
	}
	return casbinRule, err
}

// 删除CasbinRule记录
func (s *CasbinRuleRepository) DeleteCasbinRule(casbinRule *entity.CasbinRule) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&casbinRule)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新CasbinRule记录
func (s *CasbinRuleRepository) UpdateCasbinRule(casbinRule *entity.CasbinRule) (out *entity.CasbinRule, err error) {
	db := s.DbEngin
	err = db.Save(&casbinRule).Error
	if err != nil {
		return nil, err
	}
	return casbinRule, err
}

// 根据id获取CasbinRule记录
func (s *CasbinRuleRepository) FindCasbinRule(id int) (out *entity.CasbinRule, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除CasbinRule记录
func (s *CasbinRuleRepository) DeleteCasbinRuleByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.CasbinRule{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取CasbinRule记录
func (s *CasbinRuleRepository) GetCasbinRuleList(page *request.PageInfo) (list []*entity.CasbinRule, total int64, err error) {
	limit := page.Limit()
	offset := page.Offset()
	// 创建db
	db := s.DbEngin
	var casbinRules []*entity.CasbinRule
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
	}

	err = db.Model(&casbinRules).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Find(&casbinRules).Error
	if err != nil {
		return nil, 0, err
	}

	return casbinRules, total, nil
}
