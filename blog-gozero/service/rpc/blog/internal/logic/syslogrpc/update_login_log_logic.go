package syslogrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLoginLogLogic {
	return &UpdateLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新登录记录
func (l *UpdateLoginLogLogic) UpdateLoginLog(in *syslogrpc.LoginLogUpdateReq) (*syslogrpc.LoginLogDeleteResp, error) {
	exist, err := l.svcCtx.TLoginLogModel.FindOne(l.ctx, in.UserId, "logout_at = ''")
	if err != nil {
		return nil, err
	}

	exist.LogoutAt = time.Unix(in.LogoutAt, 0)
	exist.UpdatedAt = time.Now()
	_, err = l.svcCtx.TLoginLogModel.Save(l.ctx, exist)
	if err != nil {
		return nil, err
	}

	return &syslogrpc.LoginLogDeleteResp{}, nil
}
