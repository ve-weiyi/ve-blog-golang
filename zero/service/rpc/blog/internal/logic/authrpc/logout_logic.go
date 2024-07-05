package authrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登出
func (l *LogoutLogic) Logout(in *blog.LogoutReq) (*blog.EmptyResp, error) {
	find, err := l.svcCtx.UserAccountModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	find.LogoutAt = time.Now()
	// 修改登出时间
	_, err = l.svcCtx.UserAccountModel.Update(l.ctx, find)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
