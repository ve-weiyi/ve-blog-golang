package configrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateConfigLogic {
	return &CreateConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateConfigLogic) CreateConfig(in *blog.Config) (*blog.Config, error) {
	// todo: add your logic here and delete this line

	return &blog.Config{}, nil
}
