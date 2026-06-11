package permissionservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/model"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
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

// 更新 API
func (l *UpdateApiLogic) UpdateApi(in *permissionrpc.UpdateApiRequest) (*permissionrpc.UpdateApiResponse, error) {
	data := &model.TApi{
		Id:        in.Id,
		ParentId:  in.ParentId,
		Path:      in.Path,
		Name:      in.Name,
		Method:    in.Method,
		Traceable: in.Traceable,
		Status:    in.Status,
	}
	_, err := l.svcCtx.TApiModel.Update(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.UpdateApiResponse{Success: true}, nil
}
