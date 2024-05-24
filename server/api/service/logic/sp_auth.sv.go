package logic

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/feishu"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/qq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/weibo"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/temputil"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type AuthService struct {
	svcCtx *svc.ServiceContext
}

func NewAuthService(svcCtx *svc.ServiceContext) *AuthService {
	return &AuthService{
		svcCtx: svcCtx,
	}
}

func (l *AuthService) Login(reqCtx *request.Context, req *request.LoginReq) (resp *response.LoginResp, err error) {
	//获取用户
	account, err := l.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx, req.Username)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}

	//验证密码是否正确
	if !crypto.BcryptCheck(req.Password, account.Password) {
		return nil, apierr.ErrorUserPasswordError
	}

	//判断用户是否被禁用
	if account.Status == constant.UserStatusDisabled {
		return nil, apierr.ErrorUserDisabled
	}

	//生成token
	token, err := l.createToken(account.Id, account.Username, constant.LoginTypeEmail)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	info, err := l.getUserInfo(reqCtx, account)
	if err != nil {
		return nil, err
	}

	history := &entity.UserLoginHistory{
		UserId:    account.Id,
		LoginType: constant.LoginTypeEmail,
		IpAddress: reqCtx.IpAddress,
		IpSource:  reqCtx.GetIpSource(),
		Agent:     reqCtx.UserAgent,
		CreatedAt: time.Now(),
	}
	//保存此次登录记录
	_, err = l.svcCtx.UserLoginHistoryRepository.Create(reqCtx, history)
	if err != nil {
		return nil, err
	}

	// 更新用户登录信息
	_, _ = l.svcCtx.UserAccountRepository.Login(reqCtx, account)
	resp = &response.LoginResp{
		Token:        token,
		UserInfo:     info,
		LoginHistory: convertLoginHistory(history),
	}
	return resp, nil
}

func (l *AuthService) Logout(reqCtx *request.Context, req interface{}) (resp interface{}, err error) {
	glog.Info("用户登出")
	return l.svcCtx.UserAccountRepository.Logout(reqCtx, reqCtx.Uid)
}

func (l *AuthService) Logoff(reqCtx *request.Context, req interface{}) (resp interface{}, err error) {
	glog.Info("用户注销")

	return l.svcCtx.UserAccountRepository.Logoff(reqCtx, reqCtx.Uid)
}

func (l *AuthService) Register(reqCtx *request.Context, req *request.LoginReq) (resp *response.LoginResp, err error) {
	// 验证码校验
	if req.Code != "" {
		key := fmt.Sprintf("%s:%s", constant.Register, req.Username)
		if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, req.Code) {
			return nil, apierr.ErrorCaptchaVerify
		}
	}

	//获取用户
	exist, err := l.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx, req.Username)
	if exist != nil {
		return nil, apierr.ErrorUserAlreadyExist
	}

	account := &entity.UserAccount{
		Username:     req.Username,
		Password:     crypto.BcryptHash(req.Password),
		Status:       1,
		RegisterType: constant.LoginTypeEmail,
		IpAddress:    reqCtx.IpAddress,
		IpSource:     reqCtx.GetIpSource(),
	}

	_, _, err = l.svcCtx.UserAccountRepository.Register(reqCtx, account, &entity.UserInformation{})
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	info, err := l.getUserInfo(reqCtx, account)
	if err != nil {
		return nil, err
	}

	token, err := l.createToken(account.Id, account.Username, account.RegisterType)
	if err != nil {
		return nil, err
	}
	resp = &response.LoginResp{
		Token:        token,
		UserInfo:     info,
		LoginHistory: nil,
	}

	return resp, nil
}

func (l *AuthService) SendRegisterEmail(reqCtx *request.Context, req *request.UserEmailReq) (resp interface{}, err error) {
	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx, req.Username)
	if account != nil {
		return nil, apierr.ErrorUserAlreadyExist
	}

	// 验证code是否正确
	key := fmt.Sprintf("%s:%s", constant.Register, req.Username)
	code := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
	data := mail.CaptchaEmail{
		Username: req.Username,
		Code:     code,
	}
	// 组装邮件内容
	content, err := temputil.TempParseString(mail.TempRegister, data)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{req.Username},
		Subject: "注册邮件提醒",
		Content: content,
		Type:    0,
	}
	// 发送邮件
	err = l.svcCtx.EmailPublisher.PublishMessage([]byte(jsonconv.ObjectToJson(msg)))
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (l *AuthService) GetAuthorizeUrl(reqCtx *request.Context, req *request.OauthLoginReq) (resp *response.OauthLoginUrl, err error) {
	var auth oauth.Oauth
	cfg := l.svcCtx.Config.Oauth
	switch req.Platform {
	case constant.OauthQQ:
		auth = qq.NewAuthQq(convertAuthConfig(cfg.QQ))
	case constant.OauthWeibo:
		auth = weibo.NewAuthWb(convertAuthConfig(cfg.Weibo))
	case constant.OauthFeishu:
		auth = feishu.NewAuthFeishu(convertAuthConfig(cfg.Feishu))
	default:
		auth = qq.NewAuthQq(convertAuthConfig(cfg.QQ))
	}

	resp = &response.OauthLoginUrl{
		Url: auth.GetRedirectUrl(req.State),
	}
	return resp, nil
}

func (l *AuthService) OauthLogin(reqCtx *request.Context, req *request.OauthLoginReq) (resp *response.LoginResp, err error) {
	var auth oauth.Oauth
	cfg := l.svcCtx.Config.Oauth
	switch req.Platform {
	case constant.OauthQQ:
		auth = qq.NewAuthQq(convertAuthConfig(cfg.QQ))
	case constant.OauthWeibo:
		auth = weibo.NewAuthWb(convertAuthConfig(cfg.Weibo))
	case constant.OauthFeishu:
		auth = feishu.NewAuthFeishu(convertAuthConfig(cfg.Feishu))
	default:
		auth = qq.NewAuthQq(convertAuthConfig(cfg.QQ))
	}

	// 获取第三方用户信息
	info, err := auth.GetUserOpenInfo(req.Code)
	glog.JsonIndent("第三方用户信息", info)
	if err != nil {
		return nil, err
	}

	// 查询用户是否存在
	userOauth, err := l.svcCtx.UserAccountRepository.FindUserOauthByOpenid(reqCtx, info.OpenId, req.Platform)
	if userOauth == nil {
		// 用户未注册,先注册用户
		userOauth, err = l.oauthRegister(reqCtx, req, info)
		if err != nil {
			return nil, err
		}
	}

	// 用户已经注册,查询用户信息
	return l.oauthLogin(reqCtx, userOauth)
}

func (l *AuthService) oauthRegister(reqCtx *request.Context, req *request.OauthLoginReq, info *oauth.UserResult) (resp *entity.UserOauth, err error) {
	// 用户未注册,先注册用户
	pwd := crypto.BcryptHash(info.EnName)
	username := info.Email
	if username == "" {
		username = info.Mobile
	}
	userAccount := entity.UserAccount{
		Username:     username,
		Password:     pwd,
		RegisterType: req.Platform,
		IpAddress:    reqCtx.IpAddress,
		IpSource:     reqCtx.GetIpSource(),
	}

	userInfo := entity.UserInformation{
		Nickname: info.Name,
		Avatar:   info.Avatar,
		Email:    info.Email,
	}

	// 注册用户
	_, _, err = l.svcCtx.UserAccountRepository.Register(reqCtx, &userAccount, &userInfo)
	if err != nil {
		return nil, err
	}

	// 绑定用户第三方信息
	userOauth := &entity.UserOauth{
		UserId:   userAccount.Id,
		OpenId:   info.OpenId,
		Platform: req.Platform,
	}

	_, err = l.svcCtx.UserOauthRepository.Create(reqCtx, userOauth)
	if err != nil {
		return nil, err
	}

	return userOauth, nil
}

func (l *AuthService) oauthLogin(reqCtx *request.Context, req *entity.UserOauth) (resp *response.LoginResp, err error) {

	//获取用户
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", req.UserId)
	if err != nil {
		return nil, apierr.ErrorUserNotExist
	}
	//判断用户是否被禁用
	if account.Status == constant.UserStatusDisabled {
		return nil, apierr.ErrorUserDisabled
	}

	history := &entity.UserLoginHistory{
		UserId:    account.Id,
		LoginType: req.Platform,
		IpAddress: reqCtx.IpAddress,
		IpSource:  reqCtx.GetIpSource(),
		Agent:     reqCtx.UserAgent,
		CreatedAt: time.Now(),
	}
	//保存此次登录记录
	_, err = l.svcCtx.UserLoginHistoryRepository.Create(reqCtx, history)
	if err != nil {
		return nil, err
	}

	//生成token
	token, err := l.createToken(account.Id, account.Username, req.Platform)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	info, err := l.getUserInfo(reqCtx, account)
	if err != nil {
		return nil, err
	}
	resp = &response.LoginResp{
		Token:        token,
		UserInfo:     info,
		LoginHistory: convertLoginHistory(history),
	}
	return resp, nil
}

func (l *AuthService) getUserInfo(reqCtx *request.Context, account *entity.UserAccount) (resp *response.UserInfo, err error) {
	//获取用户信息
	info, err := l.svcCtx.UserAccountRepository.FindUserInfo(reqCtx, account.Id)
	if err != nil {
		return nil, err
	}

	//accountLikeSet, _ := l.svcCtx.ArticleRepository.FindUserLikeArticle(reqCtx, account.Id)
	//commentLikeSet, _ := l.svcCtx.CommentRepository.FindUserLikeComment(reqCtx, account.Id)
	//talkLikeSet, _ := l.svcCtx.TalkRepository.FindUserLikeTalk(reqCtx, account.Id)

	roles, err := l.svcCtx.RoleRepository.FindUserRoles(reqCtx, account.Id)
	resp = &response.UserInfo{
		UserId:   account.Id,
		Username: account.Username,
		Nickname: info.Nickname,
		Avatar:   info.Avatar,
		Intro:    info.Intro,
		Website:  info.Website,
		Email:    info.Email,
		//ArticleLikeSet: accountLikeSet,
		//CommentLikeSet: commentLikeSet,
		//TalkLikeSet:    talkLikeSet,
		Roles: convertRoleList(roles),
	}

	return resp, nil
}

func (l *AuthService) createToken(uid int, username string, loginType string) (token *response.Token, err error) {
	now := time.Now().Unix()
	expiresIn := time.Now().Add(7 * 24 * time.Hour).Unix()
	refreshExpiresIn := time.Now().Add(30 * 24 * time.Hour).Unix()
	issuer := "blog"

	accessToken, err := l.svcCtx.Token.CreateToken(
		jjwt.TokenExt{
			Uid:       uid,
			Username:  username,
			LoginType: loginType,
		},
		jwt.StandardClaims{
			ExpiresAt: expiresIn,
			IssuedAt:  now,
			Issuer:    issuer,
		})

	refreshToken, err := l.svcCtx.Token.CreateToken(
		jjwt.TokenExt{
			Uid:       uid,
			Username:  username,
			LoginType: loginType,
		},
		jwt.StandardClaims{
			ExpiresAt: refreshExpiresIn,
			IssuedAt:  now,
			Issuer:    issuer,
		})

	token = &response.Token{
		TokenType:        "Bearer",
		AccessToken:      accessToken,
		ExpiresIn:        expiresIn,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshExpiresIn,
		UserId:           uid,
	}

	//生成token
	return token, nil
}
