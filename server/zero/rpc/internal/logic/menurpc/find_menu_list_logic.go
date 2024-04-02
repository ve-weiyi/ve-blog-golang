package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuListLogic {
	return &FindMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMenuListLogic) FindMenuList(in *blog.PageQuery) (*blog.PageResult, error) {
	// todo: add your logic here and delete this line

	return &blog.PageResult{}, nil
}
