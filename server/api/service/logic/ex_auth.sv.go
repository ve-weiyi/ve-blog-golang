package logic

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	entity2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	request2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth/result"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	templateUtil "github.com/ve-weiyi/ve-blog-golang/server/utils/temp"
)

type AuthService struct {
	svcCtx *svc.ServiceContext
}

func NewAuthService(svcCtx *svc.ServiceContext) *AuthService {
	return &AuthService{
		svcCtx: svcCtx,
	}
}

func (s *AuthService) Login(reqCtx *request2.Context, req *request2.User) (resp *response.Login, err error) {
	//获取用户
	account, err := s.svcCtx.UserAccountRepository.LoadUserByUsername(req.Username)
	if err != nil {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户不存在！")
	}
	//判断用户是否被禁用
	if account.Status == constant.UserStatusDisabled {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户已被禁用！")
	}
	if req.Code != "" {
		//验证密码是否正确
		if !crypto.BcryptCheck(req.Password, account.Password) {
			return nil, codes.NewError(codes.CodeForbiddenOperation, "密码错误！")
		}
	}

	//获取用户信息
	info, err := s.svcCtx.UserAccountRepository.FindUserInfo(account.ID)
	if err != nil {
		return nil, err
	}

	history := &entity2.UserLoginHistory{
		UserID:    account.ID,
		LoginType: constant.LoginEmail,
		IpAddress: reqCtx.IpAddress,
		IpSource:  reqCtx.IpSource,
		CreatedAt: time.Now(),
	}
	//保存此次登录记录
	_, err = s.svcCtx.UserLoginHistoryRepository.CreateUserLoginHistory(reqCtx, history)
	if err != nil {
		return nil, err
	}

	//生成token
	token, err := s.svcCtx.Token.CreateClaims(account.ID, account.Username, history.LoginType)
	if err != nil {
		return nil, err
	}

	resp = &response.Login{
		UserInfo: convertUserDetails(account, info, history),
		Token:    token,
	}
	return resp, nil
}

func (s *AuthService) Logout(reqCtx *request2.Context, req interface{}) (resp interface{}, err error) {
	return true, nil
}

func (s *AuthService) Logoff(reqCtx *request2.Context, req interface{}) (resp interface{}, err error) {
	s.svcCtx.Log.Info("用户注销")

	return s.svcCtx.UserAccountRepository.Logoff(reqCtx, reqCtx.UID)
}

func (s *AuthService) Register(reqCtx *request2.Context, req *request2.User) (resp *response.Login, err error) {
	// 验证码校验
	if req.Code != "" {
		key := fmt.Sprintf("%s:%s", constant.Register, req.Username)
		if !s.svcCtx.Captcha.VerifyCaptcha(key, req.Code) {
			return nil, codes.ErrorCaptchaVerify
		}
	}

	//获取用户
	_, err = s.svcCtx.UserAccountRepository.LoadUserByUsername(req.Username)
	if err == nil {
		return nil, codes.ErrorUserAlreadyExist
	}

	account := &entity2.UserAccount{
		Username:     req.Username,
		Password:     crypto.BcryptHash(req.Password),
		Status:       1,
		RegisterType: constant.LoginEmail,
		IpAddress:    reqCtx.IpAddress,
		IpSource:     reqCtx.IpSource,
	}
	info := &entity2.UserInformation{}

	_, _, err = s.svcCtx.UserAccountRepository.Register(reqCtx, account, info)
	if err != nil {
		return nil, err
	}

	// 事务操作成功
	userinfo := &response.UserDetail{
		ID:        account.ID,
		Username:  account.Username,
		Nickname:  info.Nickname,
		Avatar:    info.Avatar,
		Intro:     info.Intro,
		Email:     info.Email,
		LoginType: account.RegisterType,
		IpAddress: account.IpAddress,
		IpSource:  account.IpSource,
	}

	token, err := s.svcCtx.Token.CreateClaims(account.ID, account.Username, account.RegisterType)
	if err != nil {
		return nil, err
	}
	resp = &response.Login{
		UserInfo: userinfo,
		Token:    token,
	}

	return resp, nil
}

func (s *AuthService) SendRegisterEmail(reqCtx *request2.Context, req *request2.UserEmail) (resp interface{}, err error) {
	// 验证用户是否存在
	account, err := s.svcCtx.UserAccountRepository.LoadUserByUsername(req.Username)
	if account != nil {
		return nil, codes.ErrorUserAlreadyExist
	}

	// 获取code
	key := fmt.Sprintf("%s:%s", constant.Register, req.Username)
	code := s.svcCtx.Captcha.GetCodeCaptcha(key)
	data := mail.CaptchaEmail{
		Username: req.Username,
		Code:     code,
	}
	// 组装邮件内容
	content, err := templateUtil.TempParseString(mail.TempRegister, data)
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
	err = s.svcCtx.EmailPublisher.SendMessage(jsonconv.ObjectToJson(msg))
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (s *AuthService) OauthLogin(reqCtx *request2.Context, req *request2.OauthLoginReq) (resp *response.Login, err error) {
	var auth oauth.Oauth
	cfg := s.svcCtx.Config.Oauth
	switch req.Platform {
	case constant.LoginQQ:
		auth = oauth.NewAuthQq(convertAuthConfig(cfg.QQ))
	case constant.LoginWeibo:
		auth = oauth.NewAuthWb(convertAuthConfig(cfg.Weibo))
	case constant.LoginFeishu:
		auth = oauth.NewAuthFeishu(convertAuthConfig(cfg.Feishu))
	default:
		auth = oauth.NewAuthQq(convertAuthConfig(cfg.QQ))
	}
	// 获取access_token
	token, err := auth.GetAccessToken(req.Code)
	if err != nil {
		return nil, err
	}

	// 获取第三方用户信息
	info, err := auth.GetUserInfo(token.AccessToken)
	if err != nil {
		return nil, err
	}

	// 查询用户是否存在
	userOauth, err := s.svcCtx.UserAccountRepository.FindUserOauthByOpenid(info.OpenID, req.Platform)
	if userOauth == nil {
		// 用户未注册,先注册用户
		_, err = s.oauthRegister(reqCtx, req, info)
		if err != nil {
			return nil, err
		}
	}

	// 用户已经注册,查询用户信息
	return s.oauthLogin(reqCtx, userOauth)
}

func (s *AuthService) oauthRegister(reqCtx *request2.Context, req *request2.OauthLoginReq, info *result.UserResult) (resp *response.Login, err error) {
	// 用户未注册,先注册用户
	pwd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(16)
	userAccount := entity2.UserAccount{
		ID:           0,
		Username:     info.OpenID,
		Password:     strconv.Itoa(pwd),
		RegisterType: req.Platform,
		IpAddress:    reqCtx.IpAddress,
		IpSource:     reqCtx.IpSource,
	}

	userInfo := entity2.UserInformation{
		Avatar: info.AvatarURL,
	}

	// 注册用户
	_, _, err = s.svcCtx.UserAccountRepository.Register(reqCtx, &userAccount, &userInfo)
	if err != nil {
		return nil, err
	}

	// 绑定用户第三方信息
	userOauth := &entity2.UserOauth{
		UserID:   userAccount.ID,
		OpenID:   info.OpenID,
		Platform: req.Platform,
	}

	_, err = s.svcCtx.UserOauthRepository.CreateUserOauth(reqCtx, userOauth)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AuthService) oauthLogin(reqCtx *request2.Context, req *entity2.UserOauth) (resp *response.Login, err error) {

	//获取用户
	account, err := s.svcCtx.UserAccountRepository.FindUserAccount(reqCtx, req.UserID)
	if err != nil {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户不存在！")
	}
	//判断用户是否被禁用
	if account.Status == constant.UserStatusDisabled {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户已被禁用！")
	}

	//获取用户信息
	info, err := s.svcCtx.UserAccountRepository.FindUserInfo(req.UserID)
	if err != nil {
		return nil, err
	}

	history := &entity2.UserLoginHistory{
		UserID:    account.ID,
		LoginType: req.Platform,
		IpAddress: reqCtx.IpAddress,
		IpSource:  reqCtx.IpSource,
		CreatedAt: time.Now(),
	}
	//保存此次登录记录
	_, err = s.svcCtx.UserLoginHistoryRepository.CreateUserLoginHistory(reqCtx, history)
	if err != nil {
		return nil, err
	}

	//生成token
	token, err := s.svcCtx.Token.CreateClaims(account.ID, account.Username, req.Platform)
	if err != nil {
		return nil, err
	}

	resp = &response.Login{
		UserInfo: convertUserDetails(account, info, history),
		Token:    token,
	}
	return resp, nil
}

func (s *AuthService) GetAuthorizeUrl(reqCtx *request2.Context, req *request2.OauthLoginReq) (resp *response.OauthLoginUrl, err error) {
	var auth oauth.Oauth
	cfg := s.svcCtx.Config.Oauth
	switch req.Platform {
	case constant.LoginQQ:
		auth = oauth.NewAuthQq(convertAuthConfig(cfg.QQ))
	case constant.LoginWeibo:
		auth = oauth.NewAuthWb(convertAuthConfig(cfg.Weibo))
	case constant.LoginFeishu:
		auth = oauth.NewAuthFeishu(convertAuthConfig(cfg.Feishu))
	default:
		auth = oauth.NewAuthQq(convertAuthConfig(cfg.QQ))
	}

	resp = &response.OauthLoginUrl{
		Url: auth.GetRedirectUrl(req.State),
	}
	return resp, nil
}

//func (s *AuthService) generateToken(userId int) (token string, err error) {
//	account, err := s.svcCtx.UserAccountRepository.FindUserAccount(nil, userId)
//	if err != nil {
//		return "", err
//	}
//
//	roles, err := s.svcCtx.RoleRepository.FindUserRoles(userId)
//	if err != nil {
//		return "", err
//	}
//
//	var roleLabels []string
//	for _, item := range roles {
//		roleLabels = append(roleLabels, item.RoleName)
//	}
//
//	return s.svcCtx.Token.CreateClaims(account.ID, account.Username, roleLabels)
//}

func convertUserDetails(user *entity2.UserAccount, info *entity2.UserInformation, history *entity2.UserLoginHistory) *response.UserDetail {
	userinfo := response.UserDetail{
		ID:            user.ID,
		Username:      user.Username,
		Nickname:      info.Nickname,
		Avatar:        info.Avatar,
		Intro:         info.Intro,
		Email:         info.Email,
		LoginType:     history.LoginType,
		IpAddress:     history.IpAddress,
		IpSource:      history.IpSource,
		LastLoginTime: history.CreatedAt.String(),
	}
	return &userinfo
}
