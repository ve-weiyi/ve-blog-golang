package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type UserInformationRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewUserInformationRepository(svcCtx *svc.RepositoryContext) *UserInformationRepository {
	return &UserInformationRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建UserInformation记录
func (s *UserInformationRepository) CreateUserInformation(ctx context.Context, userInformation *entity.UserInformation) (out *entity.UserInformation, err error) {
	db := s.DbEngin
	err = db.Create(&userInformation).Error
	if err != nil {
		return nil, err
	}
	return userInformation, err
}

// 删除UserInformation记录
func (s *UserInformationRepository) DeleteUserInformation(ctx context.Context, userInformation *entity.UserInformation) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&userInformation)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新UserInformation记录
func (s *UserInformationRepository) UpdateUserInformation(ctx context.Context, userInformation *entity.UserInformation) (out *entity.UserInformation, err error) {
	db := s.DbEngin
	err = db.Save(&userInformation).Error
	if err != nil {
		return nil, err
	}
	return userInformation, err
}

// 查询UserInformation记录
func (s *UserInformationRepository) GetUserInformation(ctx context.Context, id int) (out *entity.UserInformation, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除UserInformation记录
func (s *UserInformationRepository) DeleteUserInformationByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.UserInformation{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询UserInformation记录
func (s *UserInformationRepository) FindUserInformationList(ctx context.Context, page *request.PageInfo) (list []*entity.UserInformation, total int64, err error) {
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
