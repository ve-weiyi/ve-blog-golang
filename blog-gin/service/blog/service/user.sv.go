package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type UserService struct {
	svcCtx *svctx.ServiceContext
}

func NewUserService(svcCtx *svctx.ServiceContext) *UserService {
	return &UserService{
		svcCtx: svcCtx,
	}
}

// 删除用户绑定第三方平台账号
func (s *UserService) DeleteUserBindThirdParty(reqCtx *request.Context, in *dto.DeleteUserBindThirdPartyReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 获取用户信息
func (s *UserService) GetUserInfo(reqCtx *request.Context, in *dto.EmptyReq) (out *dto.UserInfoResp, err error) {
	// todo

	return
}

// 获取用户点赞列表
func (s *UserService) GetUserLike(reqCtx *request.Context, in *dto.EmptyReq) (out *dto.UserLikeResp, err error) {
	// todo

	return
}

// 修改用户头像
func (s *UserService) UpdateUserAvatar(reqCtx *request.Context, in *dto.UpdateUserAvatarReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户绑定邮箱
func (s *UserService) UpdateUserBindEmail(reqCtx *request.Context, in *dto.UpdateUserBindEmailReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户绑定手机号
func (s *UserService) UpdateUserBindPhone(reqCtx *request.Context, in *dto.UpdateUserBindPhoneReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户绑定第三方平台账号
func (s *UserService) UpdateUserBindThirdParty(reqCtx *request.Context, in *dto.UpdateUserBindThirdPartyReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户信息
func (s *UserService) UpdateUserInfo(reqCtx *request.Context, in *dto.UpdateUserInfoReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}

// 修改用户密码
func (s *UserService) UpdateUserPassword(reqCtx *request.Context, in *dto.UpdateUserPasswordReq) (out *dto.EmptyResp, err error) {
	// todo

	return
}
