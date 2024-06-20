package service

import (
	"mime/multipart"
	"path"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr/httperr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/cache"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/temputil"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
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
func (l *UserService) FindUserList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.UserDTO, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询账号信息
	userAccounts, err := l.svcCtx.UserAccountRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.UserAccountRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	var ids []int64
	for _, ua := range userAccounts {
		ids = append(ids, ua.Id)
	}

	//获取用户信息
	infos, err := l.svcCtx.UserInformationRepository.FindALL(reqCtx, "id in (?)", ids)
	if err != nil {
		return nil, 0, err
	}

	var infoMap = make(map[int64]*entity.UserInformation)
	for _, info := range infos {
		infoMap[info.Id] = info
	}

	for _, account := range userAccounts {
		info := infoMap[account.Id]
		// 查询账号角色信息
		roles, _ := l.svcCtx.RoleRepository.FindUserRoles(reqCtx, account.Id)

		item := &dto.UserDTO{
			Id:           account.Id,
			Username:     account.Username,
			Nickname:     info.Nickname,
			Status:       account.Status,
			Avatar:       info.Avatar,
			Intro:        info.Intro,
			Website:      info.Website,
			Email:        info.Email,
			RegisterType: account.RegisterType,
			IpAddress:    account.IpAddress,
			IpSource:     account.IpSource,
			CreatedAt:    account.CreatedAt,
			UpdatedAt:    account.UpdatedAt,
			Roles:        convertRoleList(roles),
		}

		list = append(list, item)
	}

	return list, total, nil
}

// 获取在线用户列表
func (l *UserService) FindOnlineUserList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.UserDTO, total int64, err error) {
	p, s := page.PageClause()
	keys, err := l.svcCtx.UserAccountRepository.Online(reqCtx, p, s)
	if err != nil {
		return nil, 0, err
	}

	glog.JsonIndent("names", keys)
	page.Limit.Page = 0
	page.Limit.PageSize = 0
	page.Conditions = append(page.Conditions, &dto.PageCondition{Field: "id", Value: keys, Operator: "in", Logic: "AND"})
	return l.FindUserList(reqCtx, page)
}

func (l *UserService) FindUserAreaList(reqCtx *request.Context, page *dto.PageQuery) (result []*dto.UserAreaDTO, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	list, err := l.svcCtx.UserAccountRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.UserAccountRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	// 分类
	AreaMap := make(map[string]int64)
	for _, item := range list {
		key := item.IpSource
		if _, ok := AreaMap[key]; ok {
			AreaMap[key]++
		} else {
			AreaMap[key] = 1
		}
	}

	for k, v := range AreaMap {
		result = append(result, &dto.UserAreaDTO{
			Name:  k,
			Value: v,
		})
	}
	return result, int64(len(result)), nil
}

func (l *UserService) FindUserLoginHistoryList(reqCtx *request.Context, page *dto.PageQuery) (result []*dto.LoginHistory, total int64, err error) {
	//获取用户
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", reqCtx.Uid)
	if err != nil {
		return nil, 0, apierr.NewApiError(httperr.CodeForbidden, "用户不存在！")
	}

	// 添加用户id条件
	c := &dto.PageCondition{Field: "user_id", Value: account.Id, Operator: "=", Logic: "AND"}
	page.Conditions = append(page.Conditions, c)

	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()
	histories, err := l.svcCtx.UserLoginHistoryRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.UserLoginHistoryRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	for _, item := range histories {
		his := convertLoginHistory(item)
		result = append(result, his)
	}

	return result, total, nil
}

func (l *UserService) DeleteUserLoginHistoryList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	//获取用户
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", reqCtx.Uid)
	if err != nil {
		return 0, apierr.NewApiError(httperr.CodeForbidden, "用户不存在！")
	}

	// 添加用户id条件
	return l.svcCtx.UserLoginHistoryRepository.Delete(reqCtx, "id in (?) and user_id = ?", req.Ids, account.Id)
}

func (l *UserService) SendForgetPwdEmail(reqCtx *request.Context, req *dto.UserEmailReq) (resp interface{}, err error) {
	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx, req.Username)
	if account == nil {
		return nil, apierr.ErrorUserNotExist
	}

	// 获取code
	key := cache.WrapCacheKey(constant.ForgetPassword, req.Username)
	code := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
	data := mail.CaptchaEmail{
		Username: req.Username,
		Code:     code,
	}

	// 组装邮件内容
	content, err := temputil.TempParseString(mail.TempForgetPassword, data)
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
	err = l.svcCtx.EmailPublisher.PublishMessage([]byte(jsonconv.ObjectToJson(msg)))
	if err != nil {
		return nil, err
	}
	return true, nil
}

func (l *UserService) ResetPassword(reqCtx *request.Context, req *dto.ResetPasswordReq) (resp interface{}, err error) {
	// 验证code是否正确
	key := cache.WrapCacheKey(constant.ForgetPassword, req.Username)
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, req.Code) {
		return nil, apierr.ErrorCaptchaVerify
	}

	// 验证用户是否存在
	account, err := l.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx, req.Username)
	if account == nil {
		return nil, apierr.ErrorUserNotExist
	}

	// 更新密码
	account.Password = crypto.BcryptHash(req.Password)
	_, err = l.svcCtx.UserAccountRepository.Update(reqCtx, account)
	if err != nil {
		return nil, err
	}

	return true, nil
}

// 修改用户角色
func (l *UserService) UpdateUserAvatar(reqCtx *request.Context, file *multipart.FileHeader) (data *entity.UserInformation, err error) {
	label := "avatar"
	url, err := l.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(reqCtx.Uid), label), file)
	if err != nil {
		return nil, err
	}

	// 保存上传记录
	up := &entity.UploadRecord{
		UserId:   reqCtx.Uid,
		Label:    label,
		FileName: file.Filename,
		FileSize: file.Size,
		FileMd5:  crypto.Md5v(file.Filename, ""),
		FileUrl:  url,
	}

	_, err = l.svcCtx.UploadRecordRepository.Create(reqCtx, up)
	if err != nil {
		return nil, err
	}

	// 更新用户信息
	information, err := l.svcCtx.UserInformationRepository.First(reqCtx, "id = ?", reqCtx.Uid)
	if err != nil {
		return nil, err
	}
	information.Avatar = url

	return l.svcCtx.UserInformationRepository.Update(reqCtx, information)
}

// 修改用户角色
func (l *UserService) UpdateUserRoles(reqCtx *request.Context, req *dto.UpdateUserRolesReq) (data interface{}, err error) {

	return l.svcCtx.RoleRepository.UpdateUserRoles(reqCtx, req.UserId, req.RoleIds)
}

// 修改用户状态
func (l *UserService) UpdateUserStatus(reqCtx *request.Context, req *entity.UserAccount) (data *entity.UserAccount, err error) {
	// 创建db
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	account.Status = req.Status
	_, err = l.svcCtx.UserAccountRepository.Update(reqCtx, account)
	if err != nil {
		return nil, err
	}

	return account, err
}

// 修改用户信息
func (l *UserService) UpdateUserInfo(reqCtx *request.Context, req *dto.UserInfoReq) (data *entity.UserInformation, err error) {
	info, err := l.svcCtx.UserAccountRepository.FindUserInfo(reqCtx, reqCtx.Uid)
	if err != nil {
		return nil, err
	}

	info.Nickname = req.Nickname
	info.Intro = req.Intro
	info.Website = req.Website
	info.Avatar = req.Avatar
	_, err = l.svcCtx.UserInformationRepository.Update(reqCtx, info)
	if err != nil {
		return nil, err
	}

	return info, err
}

func (l *UserService) GetUserInfo(reqCtx *request.Context, userId int64) (data *dto.UserInfo, err error) {
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", userId)
	if err != nil {
		return nil, apierr.NewApiError(httperr.CodeForbidden, "用户不存在！")
	}

	//获取用户信息
	info, err := l.svcCtx.UserAccountRepository.FindUserInfo(reqCtx, account.Id)
	if err != nil {
		return nil, err
	}

	//accountLikeSet, _ := l.svcCtx.ArticleRepository.FindUserLikeArticle(reqCtx, account.Id)
	//commentLikeSet, _ := l.svcCtx.CommentRepository.FindUserLikeComment(reqCtx, account.Id)
	//talkLikeSet, _ := l.svcCtx.TalkRepository.FindUserLikeTalk(reqCtx, account.Id)

	roles, err := l.svcCtx.RoleRepository.FindUserRoles(reqCtx, account.Id)
	data = &dto.UserInfo{
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

	return data, nil
}

func (l *UserService) GetUserMenus(reqCtx *request.Context, req interface{}) (data []*dto.MenuDetailsDTO, err error) {
	//查询用户信息
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", reqCtx.Uid)
	if err != nil {
		return nil, err
	}

	//查询用户角色
	roles, err := l.svcCtx.RoleRepository.FindUserRoles(reqCtx, account.Id)
	if err != nil {
		return nil, err
	}

	//查询角色权限,取交集
	menuMaps := make(map[int64]*entity.Menu)
	for _, item := range roles {
		menus, err := l.svcCtx.RoleRepository.FindRoleMenus(reqCtx, item.Id)
		if err != nil {
			return nil, err
		}
		// 去重
		for _, m := range menus {
			if _, ok := menuMaps[m.Id]; !ok {
				menuMaps[m.Id] = m
			}
		}
	}

	var list []*entity.Menu
	for _, v := range menuMaps {
		list = append(list, v)
	}

	var out dto.MenuDetailsDTO
	out.Children = getMenuChildren(out, list)

	return out.Children, err
}

func (l *UserService) GetUserApis(reqCtx *request.Context, req interface{}) (data []*dto.ApiDetailsDTO, err error) {
	//查询用户信息
	account, err := l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", reqCtx.Uid)
	if err != nil {
		return nil, err
	}

	//查询用户角色
	roles, err := l.svcCtx.RoleRepository.FindUserRoles(reqCtx, account.Id)
	if err != nil {
		return nil, err
	}

	//查询角色权限,取交集
	menuMaps := make(map[int64]*entity.Api)
	for _, item := range roles {
		menus, err := l.svcCtx.RoleRepository.FindRoleApis(reqCtx, item.Id)
		if err != nil {
			return nil, err
		}
		// 去重
		for _, m := range menus {
			if _, ok := menuMaps[m.Id]; !ok {
				menuMaps[m.Id] = m
			}
		}
	}

	var list []*entity.Api
	for _, v := range menuMaps {
		list = append(list, v)
	}

	var out dto.ApiDetailsDTO
	out.Children = getApiChildren(out, list)

	return out.Children, err
}
