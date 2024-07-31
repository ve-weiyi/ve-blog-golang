package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogoutAtLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogoutAtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogoutAtLogic {
	return &GetLogoutAtLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户登录时间
func (l *GetLogoutAtLogic) GetLogoutAt(in *blog.GetLogoutAtReq) (*blog.LogoutResp, error) {
	find, err := l.svcCtx.UserAccountModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &blog.LogoutResp{
		UserId:   in.UserId,
		LogoutAt: find.LogoutAt.Unix(),
	}, nil
}
