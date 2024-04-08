package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuLogic {
	return &FindMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMenuLogic) FindMenu(in *blog.IdReq) (*blog.Menu, error) {
	// todo: add your logic here and delete this line

	return &blog.Menu{}, nil
}
