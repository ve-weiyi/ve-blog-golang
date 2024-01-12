package logic

import (
	"context"
)

// 定义泛型接口
type Repository[T any] interface {
	// 增删改查
	Create(ctx context.Context, item T) (out T, err error)
	Update(ctx context.Context, item T) (out T, err error)
	Delete(ctx context.Context, conditions string, args ...interface{}) (rows int64, err error)
	First(ctx context.Context, conditions string, args ...interface{}) (out T, err error)
	FindALL(ctx context.Context, conditions string, args ...interface{}) (list []T, err error)
	FindList(ctx context.Context, limit int, offset int, sorts string, conditions string, args ...interface{}) (list []T, err error)
	Count(ctx context.Context, conditions string, args ...interface{}) (count int64, err error)
}
