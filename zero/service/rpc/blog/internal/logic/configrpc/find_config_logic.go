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

func (l *FindConfigLogic) FindConfig(in *blog.FindConfigReq) (*blog.FindConfigResp, error) {
	entity, err := l.svcCtx.WebsiteConfigModel.FindOneByKey(l.ctx, in.ConfigKey)
	if err != nil {
		return nil, err
	}

	return &blog.FindConfigResp{
		ConfigValue: entity.Config,
	}, nil
}
