package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAvatarLogic {
	return &UpdateUserAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户头像
func (l *UpdateUserAvatarLogic) UpdateUserAvatar(in *account.UpdateUserAvatarReq) (*account.UserInfoResp, error) {
	ui, err := l.svcCtx.UserInformationModel.First(l.ctx, "user_id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	ui.Avatar = in.Avatar

	_, err = l.svcCtx.UserInformationModel.Update(l.ctx, ui)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUserInfoModelToPb(ui), nil
}
