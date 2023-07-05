package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
)

type CasbinRuleService struct {
	svcCtx *svc.ServiceContext
}

func NewCasbinRuleService(svcCtx *svc.ServiceContext) *CasbinRuleService {
	return &CasbinRuleService{
		svcCtx: svcCtx,
	}
}

// 创建CasbinRule记录
func (s *CasbinRuleService) CreateCasbinRule(reqCtx *request.Context, casbinRule *entity.CasbinRule) (data *entity.CasbinRule, err error) {
	return s.svcCtx.CasbinRuleRepository.CreateCasbinRule(casbinRule)
}

// 删除CasbinRule记录
func (s *CasbinRuleService) DeleteCasbinRule(reqCtx *request.Context, casbinRule *entity.CasbinRule) (rows int64, err error) {
	return s.svcCtx.CasbinRuleRepository.DeleteCasbinRule(casbinRule)
}

// 更新CasbinRule记录
func (s *CasbinRuleService) UpdateCasbinRule(reqCtx *request.Context, casbinRule *entity.CasbinRule) (data *entity.CasbinRule, err error) {
	return s.svcCtx.CasbinRuleRepository.UpdateCasbinRule(casbinRule)
}

// 根据id获取CasbinRule记录
func (s *CasbinRuleService) FindCasbinRule(reqCtx *request.Context, id int) (data *entity.CasbinRule, err error) {
	return s.svcCtx.CasbinRuleRepository.FindCasbinRule(id)
}

// 批量删除CasbinRule记录
func (s *CasbinRuleService) DeleteCasbinRuleByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.CasbinRuleRepository.DeleteCasbinRuleByIds(ids)
}

// 分页获取CasbinRule记录
func (s *CasbinRuleService) GetCasbinRuleList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.CasbinRule, total int64, err error) {
	return s.svcCtx.CasbinRuleRepository.GetCasbinRuleList(page)
}
