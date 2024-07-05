package authrpclogic

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/rpcutil"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 第三方登录
func (l *OauthLoginLogic) OauthLogin(in *blog.OauthLoginReq) (*blog.LoginResp, error) {
	var auth oauth.Oauth
	for platform, v := range l.svcCtx.Oauth {
		if platform == in.Platform {
			auth = v
		}
	}

	if auth == nil {
		return nil, fmt.Errorf("platform %s is not support", in.Platform)
	}

	// 获取第三方用户信息
	info, err := auth.GetUserOpenInfo(in.Code)
	if err != nil {
		return nil, err
	}

	if info.OpenId == "" {
		return nil, fmt.Errorf("open_id is empty")
	}
	// 查询用户是否存在
	userOauth, err := l.svcCtx.UserOauthModel.FindOneByOpenIdPlatform(l.ctx, info.OpenId, in.Platform)
	if userOauth == nil {
		// 用户未注册,先注册用户
		err = l.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
			userOauth, err = l.oauthRegister(tx, in.Platform, info)
			return err
		})
		if err != nil {
			return nil, err
		}
	}

	// 用户已经注册,查询用户信息

	return l.oauthLogin(userOauth)
}

func (l *OauthLoginLogic) oauthRegister(tx *gorm.DB, platform string, info *oauth.UserResult) (resp *model.UserOauth, err error) {
	// 用户未注册,先注册用户
	pwd := crypto.BcryptHash(info.EnName)
	username := info.Email
	if username == "" {
		username = info.Mobile
	}

	// 用户账号
	userAccount := &model.UserAccount{
		Username:     username,
		Password:     pwd,
		RegisterType: platform,
		IpAddress:    "",
		IpSource:     "",
	}

	// 用户信息
	userInfo := &model.UserInformation{
		Nickname: info.Name,
		Avatar:   info.Avatar,
		Email:    info.Email,
	}

	// 绑定用户第三方信息
	userOauth := &model.UserOauth{
		OpenId:   info.OpenId,
		Platform: platform,
	}

	/** 创建用户 **/
	_, err = l.svcCtx.UserAccountModel.WithTransaction(tx).Insert(l.ctx, userAccount)
	if err != nil {
		return nil, err
	}

	/** 创建用户信息 **/
	userInfo.UserId = userAccount.Id
	_, err = l.svcCtx.UserInformationModel.WithTransaction(tx).Insert(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}

	/** 创建用户第三方信息 **/
	userOauth.UserId = userAccount.Id
	_, err = l.svcCtx.UserOauthModel.WithTransaction(tx).Insert(l.ctx, userOauth)
	if err != nil {
		return nil, err
	}

	return userOauth, nil
}

func (l *OauthLoginLogic) oauthLogin(ua *model.UserOauth) (resp *blog.LoginResp, err error) {

	//获取用户
	account, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", ua.UserId)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
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

	resp = &blog.LoginResp{
		UserId:   account.Id,
		Username: account.Username,
		Nickname: info.Nickname,
		Avatar:   info.Avatar,
		Intro:    info.Intro,
		Website:  info.Website,
		Email:    info.Email,
	}

	agent, _ := rpcutil.GetRPCUserAgent(l.ctx)
	ip, _ := rpcutil.GetRPCClientIP(l.ctx)
	is, _ := ipx.GetIpInfoByBaidu(ip)
	//登录记录
	history := &model.UserLoginHistory{
		UserId:    account.Id,
		LoginType: constant.LoginTypeOauth,
		IpAddress: ip,
		IpSource:  is.Location,
		Agent:     agent,
		CreatedAt: time.Now(),
	}

	//保存此次登录记录
	_, err = l.svcCtx.UserLoginHistoryModel.Insert(l.ctx, history)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
