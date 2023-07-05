package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type ArticleTagRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewArticleTagRepository(svcCtx *svc.RepositoryContext) *ArticleTagRepository {
	return &ArticleTagRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建ArticleTag记录
func (s *ArticleTagRepository) CreateArticleTag(ctx context.Context, articleTag *entity.ArticleTag, conditions ...*request.Condition) (out *entity.ArticleTag, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&articleTag).Error
	if err != nil {
		return nil, err
	}
	return articleTag, err
}

// 更新ArticleTag记录
func (s *ArticleTagRepository) UpdateArticleTag(ctx context.Context, articleTag *entity.ArticleTag, conditions ...*request.Condition) (out *entity.ArticleTag, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&articleTag).Error
	if err != nil {
		return nil, err
	}
	return articleTag, err
}

// 删除ArticleTag记录
func (s *ArticleTagRepository) DeleteArticleTag(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.ArticleTag{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询ArticleTag记录
func (s *ArticleTagRepository) FindArticleTag(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.ArticleTag, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除ArticleTag记录
func (s *ArticleTagRepository) DeleteArticleTagByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.ArticleTag{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询ArticleTag记录
func (s *ArticleTagRepository) FindArticleTagList(ctx context.Context, page *request.PageQuery, conditions ...*request.Condition) (list []*entity.ArticleTag, total int64, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

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
