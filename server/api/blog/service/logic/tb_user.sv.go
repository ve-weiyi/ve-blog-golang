package logic

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/ve-weiyi/go-sdk/utils/crypto"
	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	templateUtil "github.com/ve-weiyi/go-sdk/utils/temp"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/mail"
)

type UserService struct {
	svcCtx *svc.ServiceContext
}

func NewUserService(svcCtx *svc.ServiceContext) *UserService {
	return &UserService{
		svcCtx: svcCtx,
	}
}

// 分页获取UserAccount记录
func (s *UserService) GetUserList(reqCtx *request.Context, page *request.PageQuery) (list []*response.UserDetail, total int64, err error) {
	userAccounts, total, err := s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, page)

	for _, ua := range userAccounts {
		ui, err := s.GetUserInfo(reqCtx, ua.ID)
		if err != nil {
			continue
		}
		list = append(list, ui)
	}

	return list, total, nil
}

func (s *UserService) GetUserInfo(reqCtx *request.Context, userId int) (result *response.UserDetail, err error) {
	account, err := s.svcCtx.UserAccountRepository.FindUserAccount(reqCtx, userId)
	if err != nil {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户不存在！")
	}

	info, err := s.svcCtx.UserAccountRepository.FindUserInfo(account.ID)
	if err != nil {
		return nil, err
	}

	history, err := s.svcCtx.UserAccountRepository.FindLastLoginHistory(reqCtx, account.ID)
	if err != nil {
		return nil, err
	}

	roles, err := s.svcCtx.RoleRepository.FindUserRoles(userId)
	if err != nil {
		return nil, err
	}

	userinfo := &response.UserDetail{
		ID:            account.ID,
		Username:      account.Username,
		Status:        account.Status,
		CreatedAt:     account.CreatedAt.Format(time.DateTime),
		Nickname:      info.Nickname,
		Avatar:        info.Avatar,
		Intro:         info.Intro,
		Email:         info.Email,
		LoginType:     history.LoginType,
		IpAddress:     history.IpAddress,
		IpSource:      history.IpSource,
		LastLoginTime: history.CreatedAt.Format(time.DateTime),
		Roles:         roles,
	}

	return userinfo, nil
}

func (s *UserService) GetUserAreas(reqCtx *request.Context, page *request.PageQuery) (result []*response.UserArea, total int64, err error) {
	list, total, err := s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 分类
	AreaMap := make(map[string]int)
	for _, item := range list {
		key := item.IpSource
		if _, ok := AreaMap[key]; ok {
			AreaMap[key]++
		} else {
			AreaMap[key] = 1
		}
	}

	for k, v := range AreaMap {
		result = append(result, &response.UserArea{
			Name:  k,
			Value: v,
		})
	}
	return result, int64(len(result)), nil
}

func (s *UserService) GetLoginHistory(reqCtx *request.Context, page *request.PageQuery) (result []*response.LoginHistory, total int64, err error) {
	//获取用户
	account, err := s.svcCtx.UserAccountRepository.FindUserAccount(reqCtx, reqCtx.UID)
	if err != nil {
		return nil, 0, codes.NewError(codes.CodeForbiddenOperation, "用户不存在！")
	}

	page.Conditions = append(page.Conditions, &request.Condition{
		Flag:  "AND",
		Field: "user_id",
		Value: account.ID,
		Rule:  "=",
	})

	histories, total, err := s.svcCtx.UserLoginHistoryRepository.FindUserLoginHistoryList(reqCtx, page)

	for _, item := range histories {
		his := &response.LoginHistory{
			LoginType: item.LoginType,
			IpAddress: item.IpAddress,
			IpSource:  item.IpSource,
			LoginTime: item.CreatedAt.String(),
		}
		result = append(result, his)
	}
	return result, total, nil
}

func (s *UserService) SendForgetPwdEmail(reqCtx *request.Context, req *request.UserEmail) (resp interface{}, err error) {
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

func (s *UserService) ResetPassword(reqCtx *request.Context, req *request.ResetPasswordReq) (resp interface{}, err error) {
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

// 修改用户角色
func (s *UserService) UpdateUserAvatar(reqCtx *request.Context, file *multipart.FileHeader) (data interface{}, err error) {
	label := "avatar"
	url, err := s.svcCtx.Uploader.UploadFile(label, file)
	if err != nil {
		return nil, err
	}

	// 保存上传记录
	up := &entity.Upload{
		UserID:   reqCtx.UID,
		Label:    label,
		FileName: file.Filename,
		FileSize: int(file.Size),
		FileMd5:  crypto.MD5V([]byte(file.Filename)),
		FileUrl:  url,
	}
	s.svcCtx.UploadRepository.CreateUpload(reqCtx, up)

	// 更新用户信息
	information, err := s.svcCtx.UserInformationRepository.FindUserInformation(reqCtx, reqCtx.UID)
	if err != nil {
		return nil, err
	}
	information.Avatar = url

	return s.svcCtx.UserInformationRepository.UpdateUserInformation(reqCtx, information)
}

// 修改用户角色
func (s *UserService) UpdateUserRoles(reqCtx *request.Context, req *request.UpdateUserRoles) (data interface{}, err error) {

	return s.svcCtx.RoleRepository.UpdateUserRoles(reqCtx, req.UserId, req.RoleIds)
}

// 修改用户状态
func (s *UserService) UpdateUserStatus(reqCtx *request.Context, req *entity.UserAccount) (data *entity.UserAccount, err error) {
	// 创建db
	account, err := s.svcCtx.UserAccountRepository.FindUserAccount(reqCtx, req.ID)
	if err != nil {
		return nil, err
	}

	account.Status = req.Status
	_, err = s.svcCtx.UserAccountRepository.UpdateUserAccount(reqCtx, account)
	if err != nil {
		return nil, err
	}

	return account, err
}

// 修改用户信息
func (s *UserService) UpdateUserInfo(reqCtx *request.Context, req *entity.UserInformation) (data *entity.UserInformation, err error) {
	// 创建db
	info, err := s.svcCtx.UserInformationRepository.FindUserInformation(reqCtx, req.ID)
	if err != nil {
		return nil, err
	}

	info.Nickname = req.Nickname
	info.Intro = req.Intro
	info.WebSite = req.WebSite
	_, err = s.svcCtx.UserInformationRepository.UpdateUserInformation(reqCtx, info)
	if err != nil {
		return nil, err
	}

	return info, err
}
