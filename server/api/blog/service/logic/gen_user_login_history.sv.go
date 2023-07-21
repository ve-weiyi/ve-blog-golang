package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type UserLoginHistoryService struct {
	svcCtx *svc.ServiceContext
}

func NewUserLoginHistoryService(svcCtx *svc.ServiceContext) *UserLoginHistoryService {
	return &UserLoginHistoryService{
		svcCtx: svcCtx,
	}
}

// 创建UserLoginHistory记录
func (s *UserLoginHistoryService) CreateUserLoginHistory(reqCtx *request.Context, userLoginHistory *entity.UserLoginHistory) (data *entity.UserLoginHistory, err error) {
	return s.svcCtx.UserLoginHistoryRepository.CreateUserLoginHistory(reqCtx, userLoginHistory)
}

// 删除UserLoginHistory记录
func (s *UserLoginHistoryService) DeleteUserLoginHistory(reqCtx *request.Context, userLoginHistory *entity.UserLoginHistory) (rows int64, err error) {
	return s.svcCtx.UserLoginHistoryRepository.DeleteUserLoginHistory(reqCtx, userLoginHistory)
}

// 更新UserLoginHistory记录
func (s *UserLoginHistoryService) UpdateUserLoginHistory(reqCtx *request.Context, userLoginHistory *entity.UserLoginHistory) (data *entity.UserLoginHistory, err error) {
	return s.svcCtx.UserLoginHistoryRepository.UpdateUserLoginHistory(reqCtx, userLoginHistory)
}

// 查询UserLoginHistory记录
func (s *UserLoginHistoryService) GetUserLoginHistory(reqCtx *request.Context, userLoginHistory *entity.UserLoginHistory) (data *entity.UserLoginHistory, err error) {
	return s.svcCtx.UserLoginHistoryRepository.FindUserLoginHistory(reqCtx, userLoginHistory.ID)
}

// 批量删除UserLoginHistory记录
func (s *UserLoginHistoryService) DeleteUserLoginHistoryByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.UserLoginHistoryRepository.DeleteUserLoginHistoryByIds(reqCtx, ids)
}

// 分页获取UserLoginHistory记录
func (s *UserLoginHistoryService) FindUserLoginHistoryList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.UserLoginHistory, total int64, err error) {
	return s.svcCtx.UserLoginHistoryRepository.FindUserLoginHistoryList(reqCtx, page)
}
