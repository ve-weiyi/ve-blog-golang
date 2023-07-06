package logic

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
)

type TagRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewTagRepository(svcCtx *svc.RepositoryContext) *TagRepository {
	return &TagRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Tag记录
func (s *TagRepository) CreateTag(tag *entity.Tag) (out *entity.Tag, err error) {
	db := s.DbEngin
	err = db.Create(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, err
}

// 删除Tag记录
func (s *TagRepository) DeleteTag(tag *entity.Tag) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&tag)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Tag记录
func (s *TagRepository) UpdateTag(tag *entity.Tag) (out *entity.Tag, err error) {
	db := s.DbEngin
	err = db.Save(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, err
}

// 根据id获取Tag记录
func (s *TagRepository) FindTag(id int) (out *entity.Tag, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Tag记录
func (s *TagRepository) DeleteTagByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Tag{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页获取Tag记录
func (s *TagRepository) GetTagList(page *request.PageInfo) (list []*entity.Tag, total int64, err error) {
	limit := page.Limit()
	offset := page.Offset()
	// 创建db
	db := s.DbEngin
	var tags []*entity.Tag
	// 如果有条件搜索 下方会自动创建搜索语句
	if page.Order != "" && page.OrderKey != "" {
		db = db.Order(fmt.Sprintf("`%v` %v", page.Order, page.OrderKey))
	}

	err = db.Model(&tags).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(limit).Offset(offset).Find(&tags).Error
	if err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}

func (s *TagRepository) GetArticleTagList(articleId int) (list []*entity.Tag, err error) {
	// 创建db
	db := s.DbEngin
	var ats []*entity.ArticleTag
	var tags []*entity.Tag

	err = db.Where("article_id = ?", articleId).Find(&ats).Error
	if err != nil {
		return nil, err
	}

	var tagIds []int
	for _, at := range ats {
		tagIds = append(tagIds, at.TagID)
	}

	err = db.Where("id in (?)", tagIds).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}
