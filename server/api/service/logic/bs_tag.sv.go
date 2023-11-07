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
func (s *TagService) CreateTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.CreateTag(reqCtx, tag)
}

// 更新Tag记录
func (s *TagService) UpdateTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.UpdateTag(reqCtx, tag)
}

// 删除Tag记录
func (s *TagService) DeleteTag(reqCtx *request.Context, id int) (rows int, err error) {
	return s.svcCtx.TagRepository.DeleteTag(reqCtx, id)
}

// 查询Tag记录
func (s *TagService) FindTag(reqCtx *request.Context, id int) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.FindTag(reqCtx, id)
}

// 批量删除Tag记录
func (s *TagService) DeleteTagByIds(reqCtx *request.Context, ids []int) (rows int, err error) {
	return s.svcCtx.TagRepository.DeleteTagByIds(reqCtx, ids)
}

// 分页获取Tag记录
func (s *TagService) FindTagList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Tag, total int64, err error) {
	list, err = s.svcCtx.TagRepository.FindTagList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.TagRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
