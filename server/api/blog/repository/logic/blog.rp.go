package logic

import (
	"github.com/redis/go-redis/v9"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/repository/svc"
	"gorm.io/gorm"
)

type BlogRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewBlogRepository(svcCtx *svc.RepositoryContext) *BlogRepository {
	return &BlogRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}
