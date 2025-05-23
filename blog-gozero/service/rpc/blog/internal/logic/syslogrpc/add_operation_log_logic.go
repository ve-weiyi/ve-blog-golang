package syslogrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOperationLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOperationLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOperationLogLogic {
	return &AddOperationLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建操作记录
func (l *AddOperationLogLogic) AddOperationLog(in *syslogrpc.OperationLogNewReq) (*syslogrpc.EmptyResp, error) {
	entity := &model.TOperationLog{
		Id:             0,
		UserId:         in.UserId,
		TerminalId:     in.TerminalId,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		OptModule:      in.OptModule,
		OptDesc:        in.OptDesc,
		RequestUri:     in.RequestUri,
		RequestMethod:  in.RequestMethod,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
		//CreatedAt:      time.Unix(in.CreatedAt, 0),
		//UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	_, err := l.svcCtx.TOperationLogModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.EmptyResp{}, nil
}
