package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiLogic {
	return &FindApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindApiLogic) FindApi(in *blog.IdReq) (*blog.Api, error) {
	// todo: add your logic here and delete this line

	return &blog.Api{}, nil
}
