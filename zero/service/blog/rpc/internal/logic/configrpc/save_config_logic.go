package configrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveConfigLogic {
	return &SaveConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveConfigLogic) SaveConfig(in *blog.SaveConfigReq) (*blog.EmptyResp, error) {
	// 查找
	result, err := l.svcCtx.WebsiteConfigModel.FindOneByKey(l.ctx, in.ConfigKey)
	if err != nil {
		return nil, err
	}

	// 修改
	result.Config = in.ConfigValue
	_, err = l.svcCtx.WebsiteConfigModel.Save(l.ctx, result)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
