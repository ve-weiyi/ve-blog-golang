package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svctx"
)

type ApiService struct {
	svcCtx *svctx.ServiceContext
}

func NewApiService(svcCtx *svctx.ServiceContext) *ApiService {
	return &ApiService{
		svcCtx: svcCtx,
	}
}

// 创建Api记录
func (l *ApiService) CreateApi(reqCtx *request.Context, api *entity.Api) (data *entity.Api, err error) {
	return l.svcCtx.ApiRepository.Create(reqCtx, api)
}

// 更新Api记录
func (l *ApiService) UpdateApi(reqCtx *request.Context, api *entity.Api) (data *entity.Api, err error) {
	return l.svcCtx.ApiRepository.Update(reqCtx, api)
}

// 删除Api记录
func (l *ApiService) DeleteApi(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.ApiRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Api记录
func (l *ApiService) FindApi(reqCtx *request.Context, req *request.IdReq) (data *entity.Api, err error) {
	return l.svcCtx.ApiRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Api记录
func (l *ApiService) DeleteApiList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.ApiRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Api记录
func (l *ApiService) FindApiList(reqCtx *request.Context, page *dto.PageQuery) (list []*entity.Api, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.ApiRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.ApiRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
