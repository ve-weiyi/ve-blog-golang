package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type ApiService struct {
	svcCtx *svc.ServiceContext
}

func NewApiService(svcCtx *svc.ServiceContext) *ApiService {
	return &ApiService{
		svcCtx: svcCtx,
	}
}

// 创建Api记录
func (s *ApiService) CreateApi(reqCtx *request.Context, api *entity.Api) (data *entity.Api, err error) {
	return s.svcCtx.ApiRepository.Create(reqCtx, api)
}

// 更新Api记录
func (s *ApiService) UpdateApi(reqCtx *request.Context, api *entity.Api) (data *entity.Api, err error) {
	return s.svcCtx.ApiRepository.Update(reqCtx, api)
}

// 删除Api记录
func (s *ApiService) DeleteApi(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return s.svcCtx.ApiRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Api记录
func (s *ApiService) FindApi(reqCtx *request.Context, req *request.IdReq) (data *entity.Api, err error) {
	return s.svcCtx.ApiRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Api记录
func (s *ApiService) DeleteApiList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return s.svcCtx.ApiRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Api记录
func (s *ApiService) FindApiList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Api, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = s.svcCtx.ApiRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.ApiRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
