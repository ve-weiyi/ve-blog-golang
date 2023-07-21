package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type RemarkRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewRemarkRepository(svcCtx *svc.RepositoryContext) *RemarkRepository {
	return &RemarkRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Remark记录
func (s *RemarkRepository) CreateRemark(ctx context.Context, remark *entity.Remark) (out *entity.Remark, err error) {
	db := s.DbEngin
	err = db.Create(&remark).Error
	if err != nil {
		return nil, err
	}
	return remark, err
}

// 更新Remark记录
func (s *RemarkRepository) UpdateRemark(ctx context.Context, remark *entity.Remark) (out *entity.Remark, err error) {
	db := s.DbEngin
	err = db.Save(&remark).Error
	if err != nil {
		return nil, err
	}
	return remark, err
}

// 删除Remark记录
func (s *RemarkRepository) DeleteRemark(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&entity.Remark{}, "id = ?", id)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 查询Remark记录
func (s *RemarkRepository) FindRemark(ctx context.Context, id int) (out *entity.Remark, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Remark记录
func (s *RemarkRepository) DeleteRemarkByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&entity.Remark{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Remark记录
func (s *RemarkRepository) FindRemarkList(ctx context.Context, page *request.PageInfo) (list []*entity.Remark, total int64, err error) {
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
