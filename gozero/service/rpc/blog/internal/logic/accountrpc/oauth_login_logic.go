package accountrpclogic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
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
func (l *OauthLoginLogic) OauthLogin(in *accountrpc.OauthLoginReq) (*accountrpc.LoginResp, error) {
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
	userOauth, err := l.svcCtx.TUserOauthModel.FindOneByOpenIdPlatform(l.ctx, info.OpenId, in.Platform)
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
	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, userOauth.UserId)
	if err != nil {
		return nil, apierr.NewApiError(apierr.CodeUserNotExist, err.Error())
	}

	return onLogin(l.ctx, l.svcCtx, user)
}

func (l *OauthLoginLogic) oauthRegister(tx *gorm.DB, platform string, info *oauth.UserResult) (out *model.TUserOauth, err error) {
	// 用户未注册,先注册用户
	uid := uuid.NewString()
	// 使用第三方注册时，username需要唯一, 用户不能使用username登录，所以使用uuid生成。
	username := uid

	// 用户账号
	user := &model.TUser{
		UserId:    uid,
		Username:  username,
		Password:  crypto.BcryptHash(info.EnName),
		Nickname:  info.NickName,
		Avatar:    info.Avatar,
		Email:     info.Email,
		Phone:     info.Mobile,
		Info:      "",
		Status:    model.UserStatusNormal,
		LoginType: platform,
		IpAddress: "",
		IpSource:  "",
	}

	/** 创建用户 **/
	ua, err := onRegister(l.ctx, l.svcCtx, tx, user)
	if err != nil {
		return nil, err
	}

	// 绑定用户第三方信息
	userOauth := &model.TUserOauth{
		UserId:   ua.UserId,
		OpenId:   info.OpenId,
		Platform: platform,
	}

	/** 创建用户第三方信息 **/
	_, err = l.svcCtx.TUserOauthModel.WithTransaction(tx).Insert(l.ctx, userOauth)
	if err != nil {
		return nil, err
	}

	return userOauth, nil
}
