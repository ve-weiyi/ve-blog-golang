package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type UserInformationService struct {
	svcCtx *svc.ServiceContext
}

func NewUserInformationService(svcCtx *svc.ServiceContext) *UserInformationService {
	return &UserInformationService{
		svcCtx: svcCtx,
	}
}

// 创建UserInformation记录
func (s *UserInformationService) CreateUserInformation(reqCtx *request.Context, userInformation *entity.UserInformation) (data *entity.UserInformation, err error) {
	return s.svcCtx.UserInformationRepository.CreateUserInformation(userInformation)
}

// 删除UserInformation记录
func (s *UserInformationService) DeleteUserInformation(reqCtx *request.Context, userInformation *entity.UserInformation) (rows int64, err error) {
	return s.svcCtx.UserInformationRepository.DeleteUserInformation(userInformation)
}

// 更新UserInformation记录
func (s *UserInformationService) UpdateUserInformation(reqCtx *request.Context, userInformation *entity.UserInformation) (data *entity.UserInformation, err error) {
	return s.svcCtx.UserInformationRepository.UpdateUserInformation(userInformation)
}

// 查询UserInformation记录
func (s *UserInformationService) GetUserInformation(reqCtx *request.Context, userInformation *entity.UserInformation) (data *entity.UserInformation, err error) {
	return s.svcCtx.UserInformationRepository.GetUserInformation(userInformation.ID)
}

// 批量删除UserInformation记录
func (s *UserInformationService) DeleteUserInformationByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.UserInformationRepository.DeleteUserInformationByIds(ids)
}

// 分页获取UserInformation记录
func (s *UserInformationService) FindUserInformationList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.UserInformation, total int64, err error) {
	return s.svcCtx.UserInformationRepository.FindUserInformationList(page)
}
