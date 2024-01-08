package logic

import (
	"context"
)

// 定义泛型接口
type Repository[T any] interface {
	// 增删改查
	Create(ctx context.Context, item T) (out T, err error)
	Update(ctx context.Context, item T) (out T, err error)
	Delete(ctx context.Context, id int) (rows int64, err error)
	Find(ctx context.Context, id int) (out T, err error)
	// 查询列表
	Count(ctx context.Context, query string, args ...interface{}) (rows int64, err error)
	FindList(ctx context.Context, limit int, offset int, sorts string, query string, args ...interface{}) (list []T, err error)
	DeleteByConditions(ctx context.Context, query string, args ...interface{}) (rows int, err error)
	DeleteByIDS(ctx context.Context, ids []int) (rows int64, err error)
}
