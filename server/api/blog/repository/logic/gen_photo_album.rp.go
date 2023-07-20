package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type PhotoAlbumRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewPhotoAlbumRepository(svcCtx *svc.RepositoryContext) *PhotoAlbumRepository {
	return &PhotoAlbumRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建PhotoAlbum记录
func (s *PhotoAlbumRepository) CreatePhotoAlbum(ctx context.Context, photoAlbum *entity.PhotoAlbum) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin
	err = db.Create(&photoAlbum).Error
	if err != nil {
		return nil, err
	}
	return photoAlbum, err
}

// 删除PhotoAlbum记录
func (s *PhotoAlbumRepository) DeletePhotoAlbum(ctx context.Context, photoAlbum *entity.PhotoAlbum) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&photoAlbum)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新PhotoAlbum记录
func (s *PhotoAlbumRepository) UpdatePhotoAlbum(ctx context.Context, photoAlbum *entity.PhotoAlbum) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin
	err = db.Save(&photoAlbum).Error
	if err != nil {
		return nil, err
	}
	return photoAlbum, err
}

// 查询PhotoAlbum记录
func (s *PhotoAlbumRepository) GetPhotoAlbum(ctx context.Context, id int) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除PhotoAlbum记录
func (s *PhotoAlbumRepository) DeletePhotoAlbumByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.PhotoAlbum{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询PhotoAlbum记录
func (s *PhotoAlbumRepository) FindPhotoAlbumList(ctx context.Context, page *request.PageInfo) (list []*entity.PhotoAlbum, total int64, err error) {
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
