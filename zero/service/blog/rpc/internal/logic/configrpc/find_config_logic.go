package configrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindConfigLogic {
	return &FindConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindConfigLogic) FindConfig(in *blog.IdReq) (*blog.Config, error) {
	// todo: add your logic here and delete this line

	return &blog.Config{}, nil
}
