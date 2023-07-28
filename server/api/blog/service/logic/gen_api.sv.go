package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
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
func (s *ApiService) DeleteApi(reqCtx *request.Context, id int) (rows int64, err error) {
	return s.svcCtx.ApiRepository.DeleteApi(reqCtx, id)
}

// 查询Api记录
func (s *ApiService) FindApi(reqCtx *request.Context, id int) (data *entity.Api, err error) {
	return s.svcCtx.ApiRepository.FindApi(reqCtx, id)
}

// 批量删除Api记录
func (s *ApiService) DeleteApiByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.ApiRepository.DeleteApiByIds(reqCtx, ids)
}

// 分页获取Api记录
func (s *ApiService) FindApiList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Api, total int64, err error) {
	return s.svcCtx.ApiRepository.FindApiList(reqCtx, page)
}
