package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type UserAccountRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUserAccountRepository(svcCtx *svc.RepositoryContext) *UserAccountRepository {
	return &UserAccountRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建UserAccount记录
func (s *UserAccountRepository) CreateUserAccount(ctx context.Context, userAccount *entity.UserAccount) (out *entity.UserAccount, err error) {
	db := s.DbEngin
	err = db.Create(&userAccount).Error
	if err != nil {
		return nil, err
	}
	return userAccount, err
}

// 删除UserAccount记录
func (s *UserAccountRepository) DeleteUserAccount(ctx context.Context, userAccount *entity.UserAccount) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&userAccount)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新UserAccount记录
func (s *UserAccountRepository) UpdateUserAccount(ctx context.Context, userAccount *entity.UserAccount) (out *entity.UserAccount, err error) {
	db := s.DbEngin
	err = db.Save(&userAccount).Error
	if err != nil {
		return nil, err
	}
	return userAccount, err
}

// 查询UserAccount记录
func (s *UserAccountRepository) FindUserAccount(ctx context.Context, id int) (out *entity.UserAccount, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UserAccount记录
func (s *UserAccountRepository) DeleteUserAccountByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.UserAccount{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询UserAccount记录
func (s *UserAccountRepository) FindUserAccountList(ctx context.Context, page *request.PageQuery) (list []*entity.UserAccount, total int64, err error) {
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
