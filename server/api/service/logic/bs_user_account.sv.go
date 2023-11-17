package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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

// 更新UserAccount记录
func (s *UserAccountService) UpdateUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (data *entity.UserAccount, err error) {
	return s.svcCtx.UserAccountRepository.UpdateUserAccount(reqCtx, userAccount)
}

// 删除UserAccount记录
func (s *UserAccountService) DeleteUserAccount(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.UserAccountRepository.DeleteUserAccountById(reqCtx, id)
}

// 查询UserAccount记录
func (s *UserAccountService) FindUserAccount(reqCtx *request.Context, id int) (data *entity.UserAccount, err error) {
	return s.svcCtx.UserAccountRepository.FindUserAccountById(reqCtx, id)
}

// 批量删除UserAccount记录
func (s *UserAccountService) DeleteUserAccountByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.UserAccountRepository.DeleteUserAccountByIds(reqCtx, ids)
}

// 分页获取UserAccount记录
func (s *UserAccountService) FindUserAccountList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.UserAccount, total int64, err error) {
	list, err = s.svcCtx.UserAccountRepository.FindUserAccountList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.UserAccountRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
