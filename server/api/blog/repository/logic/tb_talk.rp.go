package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type TalkRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewTalkRepository(svcCtx *svc.RepositoryContext) *TalkRepository {
	return &TalkRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Talk记录
func (s *TalkRepository) CreateTalk(ctx context.Context, talk *entity.Talk) (out *entity.Talk, err error) {
	db := s.DbEngin
	err = db.Create(&talk).Error
	if err != nil {
		return nil, err
	}
	return talk, err
}

// 删除Talk记录
func (s *TalkRepository) DeleteTalk(ctx context.Context, talk *entity.Talk) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&talk)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Talk记录
func (s *TalkRepository) UpdateTalk(ctx context.Context, talk *entity.Talk) (out *entity.Talk, err error) {
	db := s.DbEngin
	err = db.Save(&talk).Error
	if err != nil {
		return nil, err
	}
	return talk, err
}

// 查询Talk记录
func (s *TalkRepository) GetTalk(ctx context.Context, id int) (out *entity.Talk, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Talk记录
func (s *TalkRepository) DeleteTalkByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Talk{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Talk记录
func (s *TalkRepository) FindTalkList(ctx context.Context, page *request.PageInfo) (list []*entity.Talk, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Orders) != 0 {
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
