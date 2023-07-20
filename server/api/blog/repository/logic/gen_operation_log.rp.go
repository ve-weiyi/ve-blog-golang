package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
)

type OperationLogRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewOperationLogRepository(svcCtx *svc.RepositoryContext) *OperationLogRepository {
	return &OperationLogRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建OperationLog记录
func (s *OperationLogRepository) CreateOperationLog(ctx context.Context, operationLog *entity.OperationLog) (out *entity.OperationLog, err error) {
	db := s.DbEngin
	err = db.Create(&operationLog).Error
	if err != nil {
		return nil, err
	}
	return operationLog, err
}

// 删除OperationLog记录
func (s *OperationLogRepository) DeleteOperationLog(ctx context.Context, operationLog *entity.OperationLog) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&operationLog)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新OperationLog记录
func (s *OperationLogRepository) UpdateOperationLog(ctx context.Context, operationLog *entity.OperationLog) (out *entity.OperationLog, err error) {
	db := s.DbEngin
	err = db.Save(&operationLog).Error
	if err != nil {
		return nil, err
	}
	return operationLog, err
}

// 查询OperationLog记录
func (s *OperationLogRepository) GetOperationLog(ctx context.Context, id int) (out *entity.OperationLog, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除OperationLog记录
func (s *OperationLogRepository) DeleteOperationLogByIds(ctx context.Context, ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.OperationLog{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询OperationLog记录
func (s *OperationLogRepository) FindOperationLogList(ctx context.Context, page *request.PageInfo) (list []*entity.OperationLog, total int64, err error) {
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
