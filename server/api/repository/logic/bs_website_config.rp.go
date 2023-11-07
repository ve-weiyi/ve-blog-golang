package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type WebsiteConfigRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewWebsiteConfigRepository(svcCtx *svc.RepositoryContext) *WebsiteConfigRepository {
	return &WebsiteConfigRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建WebsiteConfig记录
func (s *WebsiteConfigRepository) CreateWebsiteConfig(ctx context.Context, websiteConfig *entity.WebsiteConfig, conditions ...*request.Condition) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&websiteConfig).Error
	if err != nil {
		return nil, err
	}
	return websiteConfig, err
}

// 更新WebsiteConfig记录
func (s *WebsiteConfigRepository) UpdateWebsiteConfig(ctx context.Context, websiteConfig *entity.WebsiteConfig, conditions ...*request.Condition) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&websiteConfig).Error
	if err != nil {
		return nil, err
	}
	return websiteConfig, err
}

// 删除WebsiteConfig记录
func (s *WebsiteConfigRepository) DeleteWebsiteConfig(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.WebsiteConfig{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询WebsiteConfig记录
func (s *WebsiteConfigRepository) FindWebsiteConfig(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.WebsiteConfig, err error) {
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

// 批量删除WebsiteConfig记录
func (s *WebsiteConfigRepository) DeleteWebsiteConfigByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.WebsiteConfig{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询WebsiteConfig记录
func (s *WebsiteConfigRepository) FindWebsiteConfigList(ctx context.Context, page *request.PageQuery) (list []*entity.WebsiteConfig, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Sorts) != 0 {
		db = db.Order(page.OrderClause())
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
		return nil, err
	}

	return list, nil
}

// 查询总数
func (s *WebsiteConfigRepository) Count(ctx context.Context, conditions ...*request.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.ArticleTag{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询WebsiteConfig记录
func (s *WebsiteConfigRepository) FindWebsiteConfigByKey(ctx context.Context, key string, conditions ...*request.Condition) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Where("`key` = ?", key).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}
