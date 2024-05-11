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
func (l *UserAccountService) CreateUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (data *entity.UserAccount, err error) {
	return l.svcCtx.UserAccountRepository.Create(reqCtx, userAccount)
}

// 更新UserAccount记录
func (l *UserAccountService) UpdateUserAccount(reqCtx *request.Context, userAccount *entity.UserAccount) (data *entity.UserAccount, err error) {
	return l.svcCtx.UserAccountRepository.Update(reqCtx, userAccount)
}

// 删除UserAccount记录
func (l *UserAccountService) DeleteUserAccount(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.UserAccountRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询UserAccount记录
func (l *UserAccountService) FindUserAccount(reqCtx *request.Context, req *request.IdReq) (data *entity.UserAccount, err error) {
	return l.svcCtx.UserAccountRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除UserAccount记录
func (l *UserAccountService) DeleteUserAccountList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.UserAccountRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取UserAccount记录
func (l *UserAccountService) FindUserAccountList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.UserAccount, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.UserAccountRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.UserAccountRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
