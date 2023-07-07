package logic

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/codes"
)

type UserService struct {
	svcCtx *svc.ServiceContext
}

func NewUserService(svcCtx *svc.ServiceContext) *UserService {
	return &UserService{
		svcCtx: svcCtx,
	}
}

func (s *UserService) GetUserinfo(reqCtx *request.Context, userId int) (result *response.UserDetail, err error) {
	account, err := s.svcCtx.UserAccountRepository.LoadUserByUsername(reqCtx.Username)
	if err != nil {
		return nil, codes.NewError(codes.CodeForbiddenOperation, "用户不存在！")
	}

	info, err := s.svcCtx.UserAccountRepository.GetUserinfo(account.ID)
	if err != nil {
		return nil, err
	}

	history, err := s.svcCtx.UserAccountRepository.GetLastLoginHistory(reqCtx, account.ID)
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

func (s *UserService) GetLoginHistory(reqCtx *request.Context, page *request.PageInfo) (result []*response.LoginHistory, total int64, err error) {
	//获取用户
	account, err := s.svcCtx.UserAccountRepository.GetUserAccount(reqCtx, reqCtx.UID)
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

func (s *UserService) ChangePassword(req request.ChangePasswordReq) (auth *entity.UserAccount, err error) {

	return auth, nil
}

// 分页获取UserAccount记录
func (s *UserService) GetUserList(reqCtx *request.Context, page *request.PageInfo) (list []*response.UserDetail, total int64, err error) {
	userAccounts, total, err := s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, page)

	for _, ua := range userAccounts {
		ui, err := s.GetUserinfo(reqCtx, ua.ID)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, ui)
	}

	return list, total, nil
}

// 修改用户角色
func (s *UserService) UpdateUserRoles(reqCtx *request.Context, req *request.UpdateUserRoles) (data interface{}, err error) {

	return s.svcCtx.RoleRepository.UpdateUserRoles(reqCtx, req.UserId, req.RoleIds)
}

// 修改用户状态
func (s *UserService) UpdateUserStatus(reqCtx *request.Context, req *entity.UserAccount) (data *entity.UserAccount, err error) {
	// 创建db
	account, err := s.svcCtx.UserAccountRepository.GetUserAccount(reqCtx, req.ID)
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
