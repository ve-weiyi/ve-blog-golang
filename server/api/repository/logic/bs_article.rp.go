package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

type ArticleRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewArticleRepository(svcCtx *svc.RepositoryContext) *ArticleRepository {
	return &ArticleRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Article记录
func (s *ArticleRepository) CreateArticle(ctx context.Context, article *entity.Article, conditions ...*request.Condition) (out *entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Create(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 更新Article记录
func (s *ArticleRepository) UpdateArticle(ctx context.Context, article *entity.Article, conditions ...*request.Condition) (out *entity.Article, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	err = db.Save(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 删除Article记录
func (s *ArticleRepository) DeleteArticle(ctx context.Context, id int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Article{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询Article记录
func (s *ArticleRepository) FindArticle(ctx context.Context, id int, conditions ...*request.Condition) (out *entity.Article, err error) {
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

// 批量删除Article记录
func (s *ArticleRepository) DeleteArticleByIds(ctx context.Context, ids []int, conditions ...*request.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := request.WhereConditions(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.Article{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 分页查询Article记录
func (s *ArticleRepository) FindArticleList(ctx context.Context, page *request.PageQuery, conditions ...*request.Condition) (list []*entity.Article, total int64, err error) {
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
