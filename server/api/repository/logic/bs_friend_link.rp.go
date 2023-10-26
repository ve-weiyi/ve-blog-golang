package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&friendLink).Error
	if err != nil {
		return nil, err
	}
	return friendLink, err
}

// 更新FriendLink记录
func (s *FriendLinkRepository) UpdateFriendLink(ctx context.Context, friendLink *entity.FriendLink) (out *entity.FriendLink, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&friendLink).Error
	if err != nil {
		return nil, err
	}
	return friendLink, err
}

// 删除FriendLink记录
func (s *FriendLinkRepository) DeleteFriendLink(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.FriendLink{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询FriendLink记录
func (s *FriendLinkRepository) FindFriendLink(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.FriendLink, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 分页查询FriendLink记录
func (s *FriendLinkRepository) FindFriendLinkList(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.FriendLink, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sqlx.OrderClause(sorts))
	}

	// 如果有分页参数
	if page != nil && page.IsValid() {
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
func (s *FriendLinkRepository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.FriendLink{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询FriendLink记录——根据id
func (s *FriendLinkRepository) FindFriendLinkById(ctx context.Context, id int) (out *entity.FriendLink, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 删除FriendLink记录——根据id
func (s *FriendLinkRepository) DeleteFriendLinkById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.FriendLink{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除FriendLink记录——根据ids
func (s *FriendLinkRepository) DeleteFriendLinkByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.FriendLink{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}
