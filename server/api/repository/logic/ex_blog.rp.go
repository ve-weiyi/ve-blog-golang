package logic

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/repository/svc"
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
