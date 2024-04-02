package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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
func (s *RoleService) CreateRole(reqCtx *request.Context, role *entity.Role) (data *entity.Role, err error) {
	return s.svcCtx.RoleRepository.Create(reqCtx, role)
}

// 更新Role记录
func (s *RoleService) UpdateRole(reqCtx *request.Context, role *entity.Role) (data *entity.Role, err error) {
	return s.svcCtx.RoleRepository.Update(reqCtx, role)
}

// 删除Role记录
func (s *RoleService) DeleteRole(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return s.svcCtx.RoleRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Role记录
func (s *RoleService) FindRole(reqCtx *request.Context, req *request.IdReq) (data *entity.Role, err error) {
	return s.svcCtx.RoleRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Role记录
func (s *RoleService) DeleteRoleList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return s.svcCtx.RoleRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Role记录
func (s *RoleService) FindRoleList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Role, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = s.svcCtx.RoleRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.RoleRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
