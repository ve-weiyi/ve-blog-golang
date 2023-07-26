package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
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
func (s *WebsiteConfigRepository) CreateWebsiteConfig(ctx context.Context, websiteConfig *entity.WebsiteConfig) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin
	err = db.Create(&websiteConfig).Error
	if err != nil {
		return nil, err
	}
	return websiteConfig, err
}

// 删除WebsiteConfig记录
func (s *WebsiteConfigRepository) DeleteWebsiteConfig(ctx context.Context, websiteConfig *entity.WebsiteConfig) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&websiteConfig)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新WebsiteConfig记录
func (s *WebsiteConfigRepository) UpdateWebsiteConfig(ctx context.Context, websiteConfig *entity.WebsiteConfig) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin
	err = db.Save(&websiteConfig).Error
	if err != nil {
		return nil, err
	}
	return websiteConfig, err
}

// 查询WebsiteConfig记录
func (s *WebsiteConfigRepository) FindWebsiteConfig(ctx context.Context, key string) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin
	err = db.Where("`key` = ?", key).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除WebsiteConfig记录
func (s *WebsiteConfigRepository) DeleteWebsiteConfigByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.WebsiteConfig{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询WebsiteConfig记录
func (s *WebsiteConfigRepository) FindWebsiteConfigList(ctx context.Context, page *request.PageQuery) (list []*entity.WebsiteConfig, total int64, err error) {
	// 创建db
	db := s.DbEngin

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
