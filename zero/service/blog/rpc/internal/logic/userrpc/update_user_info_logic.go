package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
	ui, err := l.svcCtx.UserInformationModel.First(l.ctx, "user_id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	ui.Nickname = in.Nickname
	ui.Phone = in.Phone
	ui.Intro = in.Intro
	ui.Website = in.Website

	_, err = l.svcCtx.UserInformationModel.Update(l.ctx, ui)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
