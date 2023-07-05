package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type RoleApiService struct {
	svcCtx *svc.ServiceContext
}

func NewRoleApiService(svcCtx *svc.ServiceContext) *RoleApiService {
	return &RoleApiService{
		svcCtx: svcCtx,
	}
}

// 创建RoleApi记录
func (s *RoleApiService) CreateRoleApi(reqCtx *request.Context, roleApi *entity.RoleApi) (data *entity.RoleApi, err error) {
	return s.svcCtx.RoleApiRepository.CreateRoleApi(roleApi)
}

// 删除RoleApi记录
func (s *RoleApiService) DeleteRoleApi(reqCtx *request.Context, roleApi *entity.RoleApi) (rows int64, err error) {
	return s.svcCtx.RoleApiRepository.DeleteRoleApi(roleApi)
}

// 更新RoleApi记录
func (s *RoleApiService) UpdateRoleApi(reqCtx *request.Context, roleApi *entity.RoleApi) (data *entity.RoleApi, err error) {
	return s.svcCtx.RoleApiRepository.UpdateRoleApi(roleApi)
}

// 根据id获取RoleApi记录
func (s *RoleApiService) FindRoleApi(reqCtx *request.Context, id int) (data *entity.RoleApi, err error) {
	return s.svcCtx.RoleApiRepository.FindRoleApi(id)
}

// 批量删除RoleApi记录
func (s *RoleApiService) DeleteRoleApiByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.RoleApiRepository.DeleteRoleApiByIds(ids)
}

// 分页获取RoleApi记录
func (s *RoleApiService) GetRoleApiList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.RoleApi, total int64, err error) {
	return s.svcCtx.RoleApiRepository.GetRoleApiList(page)
}
