package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type UserOauthService struct {
	svcCtx *svc.ServiceContext
}

func NewUserOauthService(svcCtx *svc.ServiceContext) *UserOauthService {
	return &UserOauthService{
		svcCtx: svcCtx,
	}
}

// 创建UserOauth记录
func (s *UserOauthService) CreateUserOauth(reqCtx *request.Context, userOauth *entity.UserOauth) (data *entity.UserOauth, err error) {
	return s.svcCtx.UserOauthRepository.CreateUserOauth(userOauth)
}

// 删除UserOauth记录
func (s *UserOauthService) DeleteUserOauth(reqCtx *request.Context, userOauth *entity.UserOauth) (rows int64, err error) {
	return s.svcCtx.UserOauthRepository.DeleteUserOauth(userOauth)
}

// 更新UserOauth记录
func (s *UserOauthService) UpdateUserOauth(reqCtx *request.Context, userOauth *entity.UserOauth) (data *entity.UserOauth, err error) {
	return s.svcCtx.UserOauthRepository.UpdateUserOauth(userOauth)
}

// 根据id获取UserOauth记录
func (s *UserOauthService) FindUserOauth(reqCtx *request.Context, id int) (data *entity.UserOauth, err error) {
	return s.svcCtx.UserOauthRepository.FindUserOauth(id)
}

// 批量删除UserOauth记录
func (s *UserOauthService) DeleteUserOauthByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.UserOauthRepository.DeleteUserOauthByIds(ids)
}

// 分页获取UserOauth记录
func (s *UserOauthService) GetUserOauthList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.UserOauth, total int64, err error) {
	return s.svcCtx.UserOauthRepository.GetUserOauthList(page)
}
