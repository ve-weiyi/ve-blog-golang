package userservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/userrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/auth/passwordx"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type UpdateMePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMePasswordLogic {
	return &UpdateMePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新当前用户密码
func (l *UpdateMePasswordLogic) UpdateMePassword(in *userrpc.UpdateMePasswordRequest) (*userrpc.UpdateMePasswordResponse, error) {
	// 从上下文获取用户ID
	userID, err := metax.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 查询用户
	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, userID)
	if err != nil {
		return nil, err
	}

	if err := passwordx.ValidatePasswordConfirm(in.NewPassword, in.ConfirmPassword); err != nil {
		return nil, err
	}
	if err := passwordx.ValidatePassword(in.NewPassword); err != nil {
		return nil, err
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		return nil, bizerr.NewBizError(bizcode.CodeCredentialsInvalid, "旧密码不正确")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 更新密码
	_, err = l.svcCtx.TUserModel.UpdateFields(l.ctx, map[string]interface{}{
		"password": string(hashedPassword),
	}, "user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return &userrpc.UpdateMePasswordResponse{
		Success: true,
	}, nil
}
