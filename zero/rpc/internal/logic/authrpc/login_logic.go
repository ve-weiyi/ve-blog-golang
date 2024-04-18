package authrpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *LoginLogic) Login(req *account.LoginReq) (*account.LoginResp, error) {
	//验证码校验
	if req.Code != "" {
		key := fmt.Sprintf("%s:%s", constant.Register, req.Username)
		if !l.svcCtx.CaptchaRepository.VerifyCaptcha(key, req.Code) {
			return nil, apierr.ErrorCaptchaVerify
		}
	}

	//获取用户
	user, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	//验证密码是否正确
	if !crypto.BcryptCheck(req.Password, user.Password) {
		return nil, apierr.ErrorUserPasswordError
	}

	//判断用户是否被禁用
	if user.Status == constant.UserStatusDisabled {
		return nil, apierr.ErrorUserDisabled
	}

	// 获取用户信息
	info, err := l.svcCtx.UserInformationModel.FindOneByUserId(l.ctx, user.Id)
	if err != nil {
		return nil, err
	}

	resp := &account.LoginResp{
		UserId:   user.Id,
		Username: user.Username,
		Nickname: info.Nickname,
		Avatar:   info.Avatar,
		Intro:    info.Intro,
		Website:  info.Website,
		Email:    info.Email,
	}

	return resp, nil
}
