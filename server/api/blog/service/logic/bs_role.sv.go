package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type RoleService struct {
	svcCtx *svc.ServiceContext
}

func NewRoleService(svcCtx *svc.ServiceContext) *RoleService {
	return &RoleService{
		svcCtx: svcCtx,
	}
}

// 创建Role记录
func (l *RoleService) CreateRole(reqCtx *request.Context, role *entity.Role) (data *entity.Role, err error) {
	return l.svcCtx.RoleRepository.Create(reqCtx, role)
}

// 更新Role记录
func (l *RoleService) UpdateRole(reqCtx *request.Context, role *entity.Role) (data *entity.Role, err error) {
	return l.svcCtx.RoleRepository.Update(reqCtx, role)
}

// 删除Role记录
func (l *RoleService) DeleteRole(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.RoleRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Role记录
func (l *RoleService) FindRole(reqCtx *request.Context, req *request.IdReq) (data *entity.Role, err error) {
	return l.svcCtx.RoleRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Role记录
func (l *RoleService) DeleteRoleList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.RoleRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Role记录
func (l *RoleService) FindRoleList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Role, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.RoleRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.RoleRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
