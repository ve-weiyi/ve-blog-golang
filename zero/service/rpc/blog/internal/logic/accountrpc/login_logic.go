package accountrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr/codex"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/rpcutil"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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
		return nil, apierr.NewApiError(codex.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, apierr.NewApiError(codex.CodeUserNotExist, "用户不存在")
	}

	// 验证密码是否正确
	if !crypto.BcryptCheck(in.Password, account.Password) {
		return nil, apierr.NewApiError(codex.CodeUserPasswordError, "密码不正确")
	}

	return onLogin(l.svcCtx, l.ctx, account)
}

func onLogin(svcCtx *svc.ServiceContext, ctx context.Context, user *model.UserAccount) (resp *accountrpc.LoginResp, err error) {
	// 判断用户是否被禁用
	if user.Status == constant.UserStatusDisabled {
		return nil, apierr.NewApiError(codex.CodeUserDisabled, "用户已被禁用")
	}

	resp = &accountrpc.LoginResp{
		UserId:   user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Info:     user.Info,
	}

	agent, _ := rpcutil.GetRPCUserAgent(ctx)
	ip, _ := rpcutil.GetRPCClientIP(ctx)
	is, _ := ipx.GetIpInfoByBaidu(ip)
	// 登录记录
	history := &model.UserLoginHistory{
		UserId:    user.Id,
		LoginType: constant.LoginTypeOauth,
		IpAddress: ip,
		IpSource:  is.Location,
		Agent:     agent,
		LoginAt:   time.Now(),
		LogoutAt:  time.Unix(0, 0),
	}

	// 保存此次登录记录
	_, err = svcCtx.UserLoginHistoryModel.Insert(ctx, history)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
