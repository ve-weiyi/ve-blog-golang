package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type UserAccountService struct {
	svcCtx *svc.ServiceContext
}

func NewUserAccountService(svcCtx *svc.ServiceContext) *UserAccountService {
	return &UserAccountService{
		svcCtx: svcCtx,
	}
}

// 创建UserAccount记录
func (s *UserAccountService) CreateUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (data *entity.UserAccount, err error) {
	return s.svcCtx.UserAccountRepository.CreateUserAccount(reqCtx, userAccount)
}

// 删除UserAccount记录
func (s *UserAccountService) DeleteUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (rows int64, err error) {
	return s.svcCtx.UserAccountRepository.DeleteUserAccount(reqCtx, userAccount)
}

// 更新UserAccount记录
func (s *UserAccountService) UpdateUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (data *entity.UserAccount, err error) {
	return s.svcCtx.UserAccountRepository.UpdateUserAccount(reqCtx, userAccount)
}

// 查询UserAccount记录
func (s *UserAccountService) GetUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (data *entity.UserAccount, err error) {
	return s.svcCtx.UserAccountRepository.GetUserAccount(reqCtx, userAccount.ID)
}

// 批量删除UserAccount记录
func (s *UserAccountService) DeleteUserAccountByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.UserAccountRepository.DeleteUserAccountByIds(reqCtx, ids)
}

// 分页获取UserAccount记录
func (s *UserAccountService) FindUserAccountList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.UserAccount, total int64, err error) {
	return s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, page)
}
