package syslogrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLoginLogLogic {
	return &AddLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建登录记录
func (l *AddLoginLogLogic) AddLoginLog(in *syslogrpc.LoginLogNewReq) (*syslogrpc.LoginLogDetails, error) {
	agent, _ := rpcutils.GetUserAgentFromCtx(l.ctx)
	ip, _ := rpcutils.GetUserClientIPFromCtx(l.ctx)
	is, _ := ipx.GetIpSourceByBaidu(ip)

	now := time.Now()
	entity := &model.TLoginLog{
		Id:        0,
		UserId:    in.UserId,
		LoginType: in.LoginType,
		Agent:     agent,
		IpAddress: ip,
		IpSource:  is,
		LoginAt:   now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := l.svcCtx.TLoginLogModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertLoginLogOut(entity), nil
}

func convertLoginLogOut(in *model.TLoginLog) (out *syslogrpc.LoginLogDetails) {
	out = &syslogrpc.LoginLogDetails{
		Id:        in.Id,
		UserId:    in.UserId,
		LoginType: in.LoginType,
		Agent:     in.Agent,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		LoginAt:   in.LoginAt.Unix(),
		LogoutAt:  in.LogoutAt.Unix(),
	}

	return out
}
