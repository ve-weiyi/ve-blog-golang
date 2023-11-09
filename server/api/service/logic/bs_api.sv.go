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
	return s.svcCtx.ApiRepository.CreateApi(reqCtx, api)
}

// 更新Api记录
func (s *ApiService) UpdateApi(reqCtx *request.Context, api *entity.Api) (data *entity.Api, err error) {
	return s.svcCtx.ApiRepository.UpdateApi(reqCtx, api)
}

// 删除Api记录
func (s *ApiService) DeleteApi(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.ApiRepository.DeleteApiById(reqCtx, id)
}

// 查询Api记录
func (s *ApiService) FindApi(reqCtx *request.Context, id int) (data *entity.Api, err error) {
	return s.svcCtx.ApiRepository.FindApiById(reqCtx, id)
}

// 批量删除Api记录
func (s *ApiService) DeleteApiByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.ApiRepository.DeleteApiByIds(reqCtx, ids)
}

// 分页获取Api记录
func (s *ApiService) FindApiList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Api, total int64, err error) {
	list, err = s.svcCtx.ApiRepository.FindApiList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.ApiRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
