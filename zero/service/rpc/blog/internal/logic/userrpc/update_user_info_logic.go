package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户信息
func (l *UpdateUserInfoLogic) UpdateUserInfo(in *blog.UpdateUserInfoReq) (*blog.EmptyResp, error) {
	ui, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	ui.Nickname = in.Nickname
	ui.Avatar = in.Avatar
	ui.Info = in.Info

	_, err = l.svcCtx.UserAccountModel.Update(l.ctx, ui)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
