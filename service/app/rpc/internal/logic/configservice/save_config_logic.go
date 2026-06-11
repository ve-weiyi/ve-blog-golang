package configservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/model"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/configrpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
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

// 保存配置
func (l *SaveConfigLogic) SaveConfig(in *configrpc.SaveConfigRequest) (*configrpc.SaveConfigResponse, error) {
	entity := &model.TConfig{
		Key:    in.ConfigKey,
		Config: in.ConfigValue,
	}

	result, _ := l.svcCtx.TConfigModel.FindOneByKey(l.ctx, in.ConfigKey)
	if result != nil {
		entity.Id = result.Id
	}

	_, err := l.svcCtx.TConfigModel.Save(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &configrpc.SaveConfigResponse{
		Success: true,
	}, nil
}
