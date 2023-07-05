package logic

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
)

type PageRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewPageRepository(svcCtx *svc.RepositoryContext) *PageRepository {
	return &PageRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Page记录
func (s *PageRepository) CreatePage(page *entity.Page) (out *entity.Page, err error) {
	db := s.DbEngin
	err = db.Create(&page).Error
	if err != nil {
		return nil, err
	}
	return page, err
}

// 删除Page记录
func (s *PageRepository) DeletePage(page *entity.Page) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&page)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Page记录
func (s *PageRepository) UpdatePage(page *entity.Page) (out *entity.Page, err error) {
	db := s.DbEngin
	err = db.Save(&page).Error
	if err != nil {
		return nil, err
	}
	return page, err
}

// 根据id获取Page记录
func (s *PageRepository) FindPage(id int) (out *entity.Page, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Page记录
func (s *PageRepository) DeletePageByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Page{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取Page记录
func (s *PageRepository) GetPageList(page *request.PageInfo) (list []*entity.Page, total int64, err error) {
	limit := page.Limit()
	offset := page.Offset()
	// 创建db
	db := s.DbEngin
	var pages []*entity.Page
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
	}

	err = db.Model(&pages).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Find(&pages).Error
	if err != nil {
		return nil, 0, err
	}

	return pages, total, nil
}
