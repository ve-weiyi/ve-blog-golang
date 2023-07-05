package logic

import (
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
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
func (s *PhotoAlbumRepository) CreatePhotoAlbum(photoAlbum *entity.PhotoAlbum) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin
	err = db.Create(&photoAlbum).Error
	if err != nil {
		return nil, err
	}
	return photoAlbum, err
}

// 删除PhotoAlbum记录
func (s *PhotoAlbumRepository) DeletePhotoAlbum(photoAlbum *entity.PhotoAlbum) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&photoAlbum)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新PhotoAlbum记录
func (s *PhotoAlbumRepository) UpdatePhotoAlbum(photoAlbum *entity.PhotoAlbum) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin
	err = db.Save(&photoAlbum).Error
	if err != nil {
		return nil, err
	}
	return photoAlbum, err
}

// 查询PhotoAlbum记录
func (s *PhotoAlbumRepository) GetPhotoAlbum(id int) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除PhotoAlbum记录
func (s *PhotoAlbumRepository) DeletePhotoAlbumByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.PhotoAlbum{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询PhotoAlbum记录
func (s *PhotoAlbumRepository) FindPhotoAlbumList(page *request.PageInfo) (list []*entity.PhotoAlbum, total int64, err error) {
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
