package userservicelogic

import (
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/userrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
	"github.com/ve-weiyi/vkit/x/cryptox"
)

type DeactivateAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeactivateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeactivateAccountLogic {
	return &DeactivateAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 停用账号（进入冷静期）
func (l *DeactivateAccountLogic) DeactivateAccount(in *userrpc.DeactivateAccountRequest) (*userrpc.DeactivateAccountResponse, error) {
	// 从上下文获取用户ID
	userID, err := metax.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 查询用户信息
	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Password == "" {
		return nil, bizerr.NewBizError(bizcode.CodeAccountDisabled, "该账号未设置密码，无法自助注销")
	}
	if !cryptox.BcryptCheck(in.Password, user.Password) {
		return nil, bizerr.NewBizError(bizcode.CodeCredentialsInvalid, "密码不正确")
	}
	if strings.TrimSpace(in.Reason) == "" {
		return nil, bizerr.NewBizError(bizcode.CodeInvalidParam, "请填写注销原因")
	}

	// 设置冷静期（30天）
	const coolingPeriodDays = 30
	now := time.Now()
	canReactivateUntil := now.AddDate(0, 0, coolingPeriodDays)

	// 更新用户状态
	reason := strings.TrimSpace(in.Reason)
	user.Status = enums.UserStatusLogoff
	user.DeletedAt = &now
	_, err = l.svcCtx.TUserModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}

	// 注销原因只用于审计日志，不落到用户表
	l.Infof("用户注销: userId=%s reason=%s", userID, reason)

	return &userrpc.DeactivateAccountResponse{
		Success:            true,
		CoolingPeriodDays:  coolingPeriodDays,
		CanReactivateUntil: canReactivateUntil.UnixMilli(),
	}, nil
}
