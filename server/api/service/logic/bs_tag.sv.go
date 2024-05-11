package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type TagService struct {
	svcCtx *svc.ServiceContext
}

func NewTagService(svcCtx *svc.ServiceContext) *TagService {
	return &TagService{
		svcCtx: svcCtx,
	}
}

// 创建Tag记录
func (l *TagService) CreateTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return l.svcCtx.TagRepository.Create(reqCtx, tag)
}

// 更新Tag记录
func (l *TagService) UpdateTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return l.svcCtx.TagRepository.Update(reqCtx, tag)
}

// 删除Tag记录
func (l *TagService) DeleteTag(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.TagRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Tag记录
func (l *TagService) FindTag(reqCtx *request.Context, req *request.IdReq) (data *entity.Tag, err error) {
	return l.svcCtx.TagRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Tag记录
func (l *TagService) DeleteTagList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.TagRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Tag记录
func (l *TagService) FindTagList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Tag, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.TagRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.TagRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
