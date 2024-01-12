package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
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
func (s *WebsiteConfigRepository) Create(ctx context.Context, item *entity.WebsiteConfig) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item, err
}

// 更新WebsiteConfig记录
func (s *WebsiteConfigRepository) Update(ctx context.Context, item *entity.WebsiteConfig) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&item).Error
	if err != nil {
		return nil, err
	}
	return item, err
}

// 删除WebsiteConfig记录
func (s *WebsiteConfigRepository) Delete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	query := db.Delete(&entity.WebsiteConfig{})
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 查询WebsiteConfig记录
func (s *WebsiteConfigRepository) First(ctx context.Context, conditions string, args ...interface{}) (out *entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

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

func (s *WebsiteConfigRepository) FindALL(ctx context.Context, conditions string, args ...interface{}) (out []*entity.WebsiteConfig, err error) {
	db := s.DbEngin.WithContext(ctx)

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

// 分页查询WebsiteConfig记录
func (s *WebsiteConfigRepository) FindList(ctx context.Context, page int, size int, sorts string, conditions string, args ...interface{}) (list []*entity.WebsiteConfig, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

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
func (s *WebsiteConfigRepository) Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		db = db.Where(conditions, args...)
	}

	err = db.Model(&entity.WebsiteConfig{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
