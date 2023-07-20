package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type CommentRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewCommentRepository(svcCtx *svc.RepositoryContext) *CommentRepository {
	return &CommentRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Comment记录
func (s *CommentRepository) CreateComment(ctx context.Context, comment *entity.Comment) (out *entity.Comment, err error) {
	db := s.DbEngin
	err = db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, err
}

// 删除Comment记录
func (s *CommentRepository) DeleteComment(ctx context.Context, comment *entity.Comment) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&comment)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Comment记录
func (s *CommentRepository) UpdateComment(ctx context.Context, comment *entity.Comment) (out *entity.Comment, err error) {
	db := s.DbEngin
	err = db.Save(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, err
}

// 查询Comment记录
func (s *CommentRepository) GetComment(ctx context.Context, id int) (out *entity.Comment, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Comment记录
func (s *CommentRepository) DeleteCommentByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Comment{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Comment记录
func (s *CommentRepository) FindCommentList(ctx context.Context, page *request.PageInfo) (list []*entity.Comment, total int64, err error) {
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
