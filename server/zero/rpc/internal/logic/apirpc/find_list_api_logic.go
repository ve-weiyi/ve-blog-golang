package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindListApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindListApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindListApiLogic {
	return &FindListApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindListApiLogic) FindListApi(in *blog.PageQuery) (*blog.PageResult, error) {
	// todo: add your logic here and delete this line

	return &blog.PageResult{}, nil
}
