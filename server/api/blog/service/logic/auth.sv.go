package logic

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"

	"github.com/ve-weiyi/go-sdk/utils/crypto"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	templateUtil "github.com/ve-weiyi/go-sdk/utils/temp"
)

type AuthService struct {
	svcCtx *svc.ServiceContext
}

func NewAuthService(svcCtx *svc.ServiceContext) *AuthService {
	return &AuthService{
		svcCtx: svcCtx,
	}
}

func (s *AuthService) Login(reqCtx *request.Context, req *request.User) (resp *response.Login, err error) {
	//获取用户
	account, err := s.svcCtx.UserAccountRepository.LoadUserByUsername(req.Username)
	if err != nil {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户不存在！")
	}
	if req.LoginType == "" || req.LoginType == constant.LoginEmail {
		//验证密码是否正确
		if !crypto.BcryptCheck(req.Password, account.Password) {
			return nil, codes.NewError(codes.CodeForbiddenOperation, "密码错误！")
		}
	}

	//获取用户信息
	info, err := s.svcCtx.UserAccountRepository.GetUserinfo(account.ID)
	if err != nil {
		return nil, err
	}
	var history entity.UserLoginHistory

	//保存此次登录记录
	history = entity.UserLoginHistory{
		UserID:    account.ID,
		LoginType: 0,
		IpAddress: account.IpAddress,
		IpSource:  account.IpSource,
	}
	_, err = s.svcCtx.UserLoginHistoryRepository.CreateUserLoginHistory(reqCtx, &history)
	if err != nil {
		return nil, err
	}

	userinfo := response.UserDetail{
		ID:            account.ID,
		Username:      account.Username,
		Nickname:      info.Nickname,
		Avatar:        info.Avatar,
		Intro:         info.Intro,
		Email:         info.Email,
		LoginType:     history.LoginType,
		IpAddress:     history.IpAddress,
		IpSource:      history.IpSource,
		LastLoginTime: history.CreatedAt.String(),
	}

	token, err := s.generateToken(userinfo.ID)
	if err != nil {
		return nil, err
	}
	resp = &response.Login{
		UserInfo: userinfo,
		Token:    token,
	}

	return resp, nil
}

func (s *AuthService) Logout(reqCtx *request.Context, req interface{}) (resp interface{}, err error) {
	return true, nil
}

func (s *AuthService) Logoff(reqCtx *request.Context, req interface{}) (resp interface{}, err error) {
	s.svcCtx.Log.Info("用户注销")

	return s.svcCtx.UserAccountRepository.Logoff(reqCtx, reqCtx.UID)
}

func (s *AuthService) Register(reqCtx *request.Context, req *request.User) (resp *response.Login, err error) {
	// 验证码校验
	if req.LoginType == "" || req.LoginType == constant.LoginEmail {
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

	account := entity.UserAccount{
		Username:     req.Username,
		Password:     crypto.BcryptHash(req.Password),
		Status:       1,
		RegisterType: req.LoginType,
		IpAddress:    reqCtx.Ip,
		IpSource:     reqCtx.IpSource,
	}
	info := entity.UserInformation{}

	_, _, err = s.svcCtx.UserAccountRepository.Register(reqCtx, &account, &info)
	if err != nil {
		return nil, err
	}

	// 事务操作成功
	userinfo := response.UserDetail{
		ID:        account.ID,
		Username:  account.Username,
		Nickname:  info.Nickname,
		Avatar:    info.Avatar,
		Intro:     info.Intro,
		Email:     info.Email,
		LoginType: 0,
		IpAddress: account.IpAddress,
		IpSource:  account.IpSource,
	}

	token, err := s.generateToken(userinfo.ID)
	if err != nil {
		return nil, err
	}
	resp = &response.Login{
		UserInfo: userinfo,
		Token:    token,
	}

	return resp, nil
}

func (s *AuthService) ResetPassword(reqCtx *request.Context, req *request.ResetPasswordReq) (resp interface{}, err error) {
	// 验证code是否正确
	key := fmt.Sprintf("%s:%s", constant.ForgetPassword, req.Username)
	if !s.svcCtx.Captcha.VerifyCaptcha(key, req.Code) {
		return nil, codes.ErrorCaptchaVerify
	}

	// 验证用户是否存在
	account, err := s.svcCtx.UserAccountRepository.LoadUserByUsername(req.Username)
	if account == nil {
		return nil, codes.ErrorUserNotExist
	}

	// 更新密码
	account.Password = crypto.BcryptHash(req.Password)
	_, err = s.svcCtx.UserAccountRepository.UpdateUserAccount(reqCtx, account)
	if err != nil {
		return nil, err
	}

	return true, nil
}

func (s *AuthService) SendForgetPwdEmail(reqCtx *request.Context, req *request.UserEmail) (resp interface{}, err error) {
	// 验证用户是否存在
	account, err := s.svcCtx.UserAccountRepository.LoadUserByUsername(req.Username)
	if account == nil {
		return nil, codes.ErrorUserNotExist
	}

	// 获取code
	key := fmt.Sprintf("%s:%s", constant.ForgetPassword, req.Username)
	code := s.svcCtx.Captcha.GetCodeCaptcha(key)
	data := mail.CaptchaEmail{
		Username: req.Username,
		Code:     code,
	}

	// 组装邮件内容
	content, err := templateUtil.TempParseString(mail.TempForgetPassword, data)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{req.Username},
		Subject: "忘记密码",
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

func (s *AuthService) SendRegisterEmail(reqCtx *request.Context, req *request.UserEmail) (resp interface{}, err error) {
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

func (s *AuthService) OauthLogin(reqCtx *request.Context, req *request.OauthLoginReq) (resp *response.Login, err error) {
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
		pwd := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(16)
		userAccount := entity.UserAccount{
			ID:           0,
			Username:     info.OpenID,
			Password:     strconv.Itoa(pwd),
			RegisterType: req.Platform,
			IpAddress:    reqCtx.Ip,
			IpSource:     reqCtx.IpSource,
		}

		userInfo := entity.UserInformation{
			Avatar: info.AvatarURL,
		}

		// 注册用户
		_, _, err = s.svcCtx.UserAccountRepository.Register(reqCtx, &userAccount, &userInfo)
		if err != nil {
			return nil, err
		}

		// 绑定用户第三方信息
		userOauth = &entity.UserOauth{
			UserID:   userAccount.ID,
			OpenID:   info.OpenID,
			Platform: req.Platform,
		}

		_, err = s.svcCtx.UserOauthRepository.CreateUserOauth(reqCtx, userOauth)
		if err != nil {
			return nil, err
		}
	}

	// 用户已经注册,查询用户信息
	login, err := s.Login(reqCtx, &request.User{
		Username:  info.OpenID,
		Password:  "",
		Code:      "",
		LoginType: req.Platform,
	})
	if err != nil {
		return nil, err
	}

	return login, nil
}

func (s *AuthService) GetAuthorizeUrl(reqCtx *request.Context, req *request.OauthLoginReq) (resp *response.OauthLoginUrl, err error) {
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

func (s *AuthService) generateToken(userId int) (token string, err error) {
	account, err := s.svcCtx.UserAccountRepository.GetUserAccount(nil, userId)
	if err != nil {
		return "", err
	}

	roles, err := s.svcCtx.RoleRepository.FindUserRoles(userId)
	if err != nil {
		return "", err
	}

	var roleLabels []string
	for _, item := range roles {
		roleLabels = append(roleLabels, item.RoleName)
	}

	return s.svcCtx.Token.CreateClaims(account.ID, account.Username, roleLabels)
}
