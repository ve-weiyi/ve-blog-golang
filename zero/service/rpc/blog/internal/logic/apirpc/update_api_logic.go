package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新接口
func (l *UpdateApiLogic) UpdateApi(in *blog.Api) (*blog.Api, error) {
	entity := convert.ConvertApiPbToModel(in)

	_, err := l.svcCtx.ApiModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertApiModelToPb(entity), nil
}
