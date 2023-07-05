package logic

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
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
func (s *ArticleTagRepository) CreateArticleTag(articleTag *entity.ArticleTag) (out *entity.ArticleTag, err error) {
	db := s.DbEngin
	err = db.Create(&articleTag).Error
	if err != nil {
		return nil, err
	}
	return articleTag, err
}

// 删除ArticleTag记录
func (s *ArticleTagRepository) DeleteArticleTag(articleTag *entity.ArticleTag) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&articleTag)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新ArticleTag记录
func (s *ArticleTagRepository) UpdateArticleTag(articleTag *entity.ArticleTag) (out *entity.ArticleTag, err error) {
	db := s.DbEngin
	err = db.Save(&articleTag).Error
	if err != nil {
		return nil, err
	}
	return articleTag, err
}

// 根据id获取ArticleTag记录
func (s *ArticleTagRepository) FindArticleTag(id int) (out *entity.ArticleTag, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除ArticleTag记录
func (s *ArticleTagRepository) DeleteArticleTagByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.ArticleTag{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取ArticleTag记录
func (s *ArticleTagRepository) GetArticleTagList(page *request.PageInfo) (list []*entity.ArticleTag, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
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

	// 查询表记录总数
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
