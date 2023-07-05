package logic

import (
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
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
func (s *ArticleRepository) CreateArticle(article *entity.Article) (out *entity.Article, err error) {
	db := s.DbEngin
	err = db.Create(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 删除Article记录
func (s *ArticleRepository) DeleteArticle(article *entity.Article) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&article)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Article记录
func (s *ArticleRepository) UpdateArticle(article *entity.Article) (out *entity.Article, err error) {
	db := s.DbEngin
	err = db.Save(&article).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

// 根据id获取Article记录
func (s *ArticleRepository) FindArticle(id int) (out *entity.Article, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Article记录
func (s *ArticleRepository) DeleteArticleByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Article{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取Article记录
func (s *ArticleRepository) GetArticleList(page *request.PageInfo) (list []*entity.Article, total int64, err error) {
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

	// 查询总数
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

// 根据分类id获取文章
func (s *ArticleRepository) GetArticleListByCategoryId(categoryId int) (list []*entity.Article, total int64, err error) {
	db := s.DbEngin
	err = db.Model(&entity.Article{}).Where("category_id = ?", categoryId).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	total = int64(len(list))
	return list, total, nil
}

// 根据标签id获取文章
func (s *ArticleRepository) GetArticleListByTagId(tagId int) (list []*entity.Article, total int64, err error) {
	db := s.DbEngin

	// 获取文章标签映射
	var ats []*entity.ArticleTag
	err = db.Where("tag_id = ?", tagId).Find(&ats).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取文章id列表
	var ids []int
	for _, at := range ats {
		ids = append(ids, at.ArticleID)
	}

	// 获取文章列表
	err = db.Where("id in (?)", ids).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	total = int64(len(list))
	return list, total, nil
}

// 获取推荐文章,与id相同分类的文章
func (s *ArticleRepository) GetRecommendArticle(cateId int) (list []*entity.Article, err error) {
	db := s.DbEngin
	err = db.Where("category_id = ?", cateId).Limit(5).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取上一篇文章
func (s *ArticleRepository) GetLastArticle(id int) (out *entity.Article, err error) {
	db := s.DbEngin
	var lastArticle entity.Article
	err = db.Where("id < ?", id).Order("`id` desc").First(&lastArticle).Error
	if err != nil {
		return nil, nil
	}

	return &lastArticle, nil
}

// 获取下一篇文章
func (s *ArticleRepository) GetNextArticle(id int) (out *entity.Article, err error) {
	db := s.DbEngin
	var nextArticle entity.Article
	err = db.Where("id > ?", id).Order("`id` asc").First(&nextArticle).Error
	if err != nil {
		return nil, nil
	}

	return &nextArticle, nil
}
