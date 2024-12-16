package rediskey

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
)

type RedisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisLogic {
	return &RedisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
