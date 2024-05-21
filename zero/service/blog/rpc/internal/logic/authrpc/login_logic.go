package authrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
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
func (l *LoginLogic) Login(req *blog.LoginReq) (*blog.LoginResp, error) {
	//获取用户
	user, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}

	//验证密码是否正确
	if !crypto.BcryptCheck(req.Password, user.Password) {
		return nil, err
	}

	//判断用户是否被禁用
	if user.Status == constant.UserStatusDisabled {
		return nil, err
	}

	// 获取用户信息
	info, err := l.svcCtx.UserInformationModel.FindOneByUserId(l.ctx, user.Id)
	if err != nil {
		return nil, err
	}

	resp := &blog.LoginResp{
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
