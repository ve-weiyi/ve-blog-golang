package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
func (l *LoginLogic) Login(in *blog.LoginReq) (*blog.LoginResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, apierr.ErrorInvalidParam
	}

	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	//验证密码是否正确
	if !crypto.BcryptCheck(in.Password, account.Password) {
		return nil, apierr.ErrorUserPasswordError
	}

	//判断用户是否被禁用
	if account.Status == constant.UserStatusDisabled {
		return nil, apierr.ErrorUserDisabled
	}

	// 获取用户信息
	info, err := l.svcCtx.UserInformationModel.FindOneByUserId(l.ctx, account.Id)
	if err != nil {
		return nil, err
	}

	resp := &blog.LoginResp{
		UserId:   account.Id,
		Username: account.Username,
		Nickname: info.Nickname,
		Avatar:   info.Avatar,
		Intro:    info.Intro,
		Website:  info.Website,
		Email:    info.Email,
	}

	return resp, nil
}
