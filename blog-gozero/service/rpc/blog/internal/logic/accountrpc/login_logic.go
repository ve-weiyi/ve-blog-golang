package accountrpclogic

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/bizerr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/metadatax"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *accountrpc.LoginReq) (*accountrpc.LoginResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, bizerr.NewBizError(bizerr.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	account, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, bizerr.NewBizError(bizerr.CodeUserNotExist, "用户不存在")
	}

	// 验证密码是否正确
	if !crypto.BcryptCheck(in.Password, account.Password) {
		return nil, bizerr.NewBizError(bizerr.CodeUserPasswordError, "密码不正确")
	}

	return onLogin(l.ctx, l.svcCtx, account)
}

func onLogin(ctx context.Context, svcCtx *svc.ServiceContext, user *model.TUser) (resp *accountrpc.LoginResp, err error) {
	// 判断用户是否被禁用
	if user.Status == constant.UserStatusDisabled {
		return nil, bizerr.NewBizError(bizerr.CodeUserDisabled, "用户已被禁用")
	}

	agent, _ := metadatax.GetRPCUserAgent(ctx)
	ip, _ := metadatax.GetRPCClientIP(ctx)
	is, _ := ipx.GetIpSourceByBaidu(ip)
	// 查找用户角色
	rList, err := getUserRoles(ctx, svcCtx, user.UserId)
	if err != nil {
		return nil, err
	}

	var roles []*accountrpc.UserRoleLabel
	for _, role := range rList {
		m := &accountrpc.UserRoleLabel{
			RoleId:      role.Id,
			RoleKey:     role.RoleKey,
			RoleLabel:   role.RoleLabel,
			RoleComment: role.RoleComment,
		}

		roles = append(roles, m)
	}

	resp = &accountrpc.LoginResp{
		UserId:    user.UserId,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Phone:     user.Phone,
		Info:      user.Info,
		LoginType: user.LoginType,
		IpAddress: ip,
		IpSource:  is,
		Roles:     roles,
	}

	// 登录记录
	history := &model.TUserLoginHistory{
		UserId:    user.UserId,
		LoginType: user.LoginType,
		IpAddress: ip,
		IpSource:  is,
		Agent:     agent,
		LoginAt:   time.Now(),
		LogoutAt:  time.Unix(0, 0),
	}

	// 保存此次登录记录
	_, err = svcCtx.TUserLoginHistoryModel.Insert(ctx, history)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
