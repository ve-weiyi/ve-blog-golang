package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
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
	db := s.DbEngin.WithContext(ctx)

	err = db.Create(&photoAlbum).Error
	if err != nil {
		return nil, err
	}
	return photoAlbum, err
}

// 更新PhotoAlbum记录
func (s *PhotoAlbumRepository) UpdatePhotoAlbum(ctx context.Context, photoAlbum *entity.PhotoAlbum) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Save(&photoAlbum).Error
	if err != nil {
		return nil, err
	}
	return photoAlbum, err
}

// 删除PhotoAlbum记录
func (s *PhotoAlbumRepository) DeletePhotoAlbum(ctx context.Context, conditions ...*sqlx.Condition) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	query := db.Delete(&entity.PhotoAlbum{})
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 查询PhotoAlbum记录
func (s *PhotoAlbumRepository) FindPhotoAlbum(ctx context.Context, conditions ...*sqlx.Condition) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 分页查询PhotoAlbum记录
func (s *PhotoAlbumRepository) FindPhotoAlbumList(ctx context.Context, page *sqlx.PageLimit, sorts []*sqlx.Sort, conditions ...*sqlx.Condition) (list []*entity.PhotoAlbum, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	// 如果有搜索条件
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(sorts) != 0 {
		db = db.Order(sqlx.OrderClause(sorts))
	}

	// 如果有分页参数
	if page != nil && page.IsValid() {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 查询总数
func (s *PhotoAlbumRepository) Count(ctx context.Context, conditions ...*sqlx.Condition) (count int64, err error) {
	db := s.DbEngin.WithContext(ctx)

	// 如果有条件语句
	if len(conditions) != 0 {
		query, args := sqlx.ConditionClause(conditions)
		db = db.Where(query, args...)
	}

	err = db.Model(&entity.PhotoAlbum{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 查询PhotoAlbum记录——根据id
func (s *PhotoAlbumRepository) FindPhotoAlbumById(ctx context.Context, id int) (out *entity.PhotoAlbum, err error) {
	db := s.DbEngin.WithContext(ctx)

	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 删除PhotoAlbum记录——根据id
func (s *PhotoAlbumRepository) DeletePhotoAlbumById(ctx context.Context, id int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.PhotoAlbum{}, "id = ?", id)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}

// 批量删除PhotoAlbum记录——根据ids
func (s *PhotoAlbumRepository) DeletePhotoAlbumByIds(ctx context.Context, ids []int) (rows int, err error) {
	db := s.DbEngin.WithContext(ctx)

	query := db.Delete(&entity.PhotoAlbum{}, "id in ?", ids)
	err = query.Error
	rows = int(query.RowsAffected)
	return rows, err
}
