package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
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
	return s.svcCtx.TagRepository.CreateTag(tag)
}

// 删除Tag记录
func (s *TagService) DeleteTag(reqCtx *request.Context, tag *entity.Tag) (rows int64, err error) {
	return s.svcCtx.TagRepository.DeleteTag(tag)
}

// 更新Tag记录
func (s *TagService) UpdateTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.UpdateTag(tag)
}

// 根据id获取Tag记录
func (s *TagService) FindTag(reqCtx *request.Context, id int) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.FindTag(id)
}

// 批量删除Tag记录
func (s *TagService) DeleteTagByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.TagRepository.DeleteTagByIds(ids)
}

// 分页获取Tag记录
func (s *TagService) GetTagList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Tag, total int64, err error) {
	return s.svcCtx.TagRepository.GetTagList(page)
}
