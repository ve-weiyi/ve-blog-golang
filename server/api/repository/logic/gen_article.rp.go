package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
)

<<<<<<<< HEAD:server/api/blog/repository/logic/gen_article.rp.go
type ArticleRepository struct {
========
type RemarkRepository struct {
>>>>>>>> 7bb7cd01 (优化 (#3)):server/api/blog/repository/logic/tb_remark.rp.go
	DbEngin *gorm.DB
	Cache   *redis.Client
}

<<<<<<<< HEAD:server/api/blog/repository/logic/gen_article.rp.go
func NewArticleRepository(svcCtx *svc.RepositoryContext) *ArticleRepository {
	return &ArticleRepository{
========
func NewRemarkRepository(svcCtx *svc.RepositoryContext) *RemarkRepository {
	return &RemarkRepository{
>>>>>>>> 7bb7cd01 (优化 (#3)):server/api/blog/repository/logic/tb_remark.rp.go
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

<<<<<<<< HEAD:server/api/blog/repository/logic/gen_article.rp.go
// 创建Article记录
func (s *ArticleRepository) CreateArticle(ctx context.Context, article *entity.Article) (out *entity.Article, err error) {
	db := s.DbEngin
	err = db.Create(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 更新Article记录
func (s *ArticleRepository) UpdateArticle(ctx context.Context, article *entity.Article) (out *entity.Article, err error) {
	db := s.DbEngin
	err = db.Save(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 删除Article记录
func (s *ArticleRepository) DeleteArticle(ctx context.Context, id int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&entity.Article{}, "id = ?", id)
========
// 创建Remark记录
func (s *RemarkRepository) CreateRemark(ctx context.Context, remark *entity.Remark) (out *entity.Remark, err error) {
	db := s.DbEngin
	err = db.Create(&remark).Error
	if err != nil {
		return nil, err
	}
	return remark, err
}

// 删除Remark记录
func (s *RemarkRepository) DeleteRemark(ctx context.Context, remark *entity.Remark) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&remark)
>>>>>>>> 7bb7cd01 (优化 (#3)):server/api/blog/repository/logic/tb_remark.rp.go
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

<<<<<<<< HEAD:server/api/blog/repository/logic/gen_article.rp.go
// 查询Article记录
func (s *ArticleRepository) FindArticle(ctx context.Context, id int) (out *entity.Article, err error) {
========
// 更新Remark记录
func (s *RemarkRepository) UpdateRemark(ctx context.Context, remark *entity.Remark) (out *entity.Remark, err error) {
	db := s.DbEngin
	err = db.Save(&remark).Error
	if err != nil {
		return nil, err
	}
	return remark, err
}

// 查询Remark记录
func (s *RemarkRepository) GetRemark(ctx context.Context, id int) (out *entity.Remark, err error) {
>>>>>>>> 7bb7cd01 (优化 (#3)):server/api/blog/repository/logic/tb_remark.rp.go
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

<<<<<<<< HEAD:server/api/blog/repository/logic/gen_article.rp.go
// 批量删除Article记录
func (s *ArticleRepository) DeleteArticleByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&entity.Article{}, "id in ?", ids)
========
// 批量删除Remark记录
func (s *RemarkRepository) DeleteRemarkByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Remark{}, "id in ?", ids)
>>>>>>>> 7bb7cd01 (优化 (#3)):server/api/blog/repository/logic/tb_remark.rp.go
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

<<<<<<<< HEAD:server/api/blog/repository/logic/gen_article.rp.go
// 分页查询Article记录
func (s *ArticleRepository) FindArticleList(ctx context.Context, page *request.PageQuery) (list []*entity.Article, total int64, err error) {
========
// 分页查询Remark记录
func (s *RemarkRepository) FindRemarkList(ctx context.Context, page *request.PageInfo) (list []*entity.Remark, total int64, err error) {
>>>>>>>> 7bb7cd01 (优化 (#3)):server/api/blog/repository/logic/tb_remark.rp.go
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
