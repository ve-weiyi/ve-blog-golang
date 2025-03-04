package syslogrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVisitLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVisitLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVisitLogLogic {
	return &AddVisitLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建访问记录
func (l *AddVisitLogLogic) AddVisitLog(in *syslogrpc.VisitLogNewReq) (*syslogrpc.VisitLogDetails, error) {
	entity := convertVisitLogIn(in)

	_, err := l.svcCtx.TVisitLogModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertVisitLogOut(entity), nil
}

func convertVisitLogIn(in *syslogrpc.VisitLogNewReq) (out *model.TVisitLog) {
	out = &model.TVisitLog{
		Id:         0,
		UserId:     in.UserId,
		TerminalId: in.TerminalId,
		IpAddress:  in.IpAddress,
		IpSource:   in.IpSource,
		Os:         in.Os,
		Browser:    in.Browser,
		Page:       in.Page,
		//CreatedAt:      time.Unix(in.CreatedAt, 0),
		//UpdatedAt:      time.Unix(in.UpdatedAt, 0),
	}

	return out
}

func convertVisitLogOut(in *model.TVisitLog) (out *syslogrpc.VisitLogDetails) {
	out = &syslogrpc.VisitLogDetails{
		Id:         in.Id,
		UserId:     in.UserId,
		TerminalId: in.TerminalId,
		IpAddress:  in.IpAddress,
		IpSource:   in.IpSource,
		Os:         in.Os,
		Browser:    in.Browser,
		Page:       in.Page,
		CreatedAt:  in.CreatedAt.Unix(),
		UpdatedAt:  in.UpdatedAt.Unix(),
	}

	return out
}
