package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type FriendLinkRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewFriendLinkRepository(svcCtx *svc.RepositoryContext) *FriendLinkRepository {
	return &FriendLinkRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建FriendLink记录
func (s *FriendLinkRepository) CreateFriendLink(ctx context.Context, friendLink *entity.FriendLink) (out *entity.FriendLink, err error) {
	db := s.DbEngin
	err = db.Create(&friendLink).Error
	if err != nil {
		return nil, err
	}
	return friendLink, err
}

// 删除FriendLink记录
func (s *FriendLinkRepository) DeleteFriendLink(ctx context.Context, friendLink *entity.FriendLink) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&friendLink)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新FriendLink记录
func (s *FriendLinkRepository) UpdateFriendLink(ctx context.Context, friendLink *entity.FriendLink) (out *entity.FriendLink, err error) {
	db := s.DbEngin
	err = db.Save(&friendLink).Error
	if err != nil {
		return nil, err
	}
	return friendLink, err
}

// 查询FriendLink记录
func (s *FriendLinkRepository) GetFriendLink(ctx context.Context, id int) (out *entity.FriendLink, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除FriendLink记录
func (s *FriendLinkRepository) DeleteFriendLinkByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.FriendLink{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询FriendLink记录
func (s *FriendLinkRepository) FindFriendLinkList(ctx context.Context, page *request.PageInfo) (list []*entity.FriendLink, total int64, err error) {
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
