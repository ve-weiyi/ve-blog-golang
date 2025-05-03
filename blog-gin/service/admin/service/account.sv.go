package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type AccountService struct {
	svcCtx *svctx.ServiceContext
}

func NewAccountService(svcCtx *svctx.ServiceContext) *AccountService {
	return &AccountService{
		svcCtx: svcCtx,
	}
}

// 查询用户列表
func (s *AccountService) FindAccountList(reqCtx *request.Context, in *dto.AccountQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 查询在线用户列表
func (s *AccountService) FindAccountOnlineList(reqCtx *request.Context, in *dto.AccountQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 修改用户密码
func (s *AccountService) UpdateAccountPassword(reqCtx *request.Context, in *dto.UpdateAccountPasswordReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户角色
func (s *AccountService) UpdateAccountRoles(reqCtx *request.Context, in *dto.UpdateAccountRolesReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户状态
func (s *AccountService) UpdateAccountStatus(reqCtx *request.Context, in *dto.UpdateAccountStatusReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}
