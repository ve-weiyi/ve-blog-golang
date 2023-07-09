package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
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

// 删除Tag记录
func (s *TagService) DeleteTag(reqCtx *request.Context, tag *entity.Tag) (rows int64, err error) {
	return s.svcCtx.TagRepository.DeleteTag(reqCtx, tag)
}

// 更新Tag记录
func (s *TagService) UpdateTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.UpdateTag(reqCtx, tag)
}

// 查询Tag记录
func (s *TagService) FindTag(reqCtx *request.Context, tag *entity.Tag) (data *entity.Tag, err error) {
	return s.svcCtx.TagRepository.GetTag(reqCtx, tag.ID)
}

// 批量删除Tag记录
func (s *TagService) DeleteTagByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.TagRepository.DeleteTagByIds(reqCtx, ids)
}

// 分页获取Tag记录
func (s *TagService) FindTagList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Tag, total int64, err error) {
	return s.svcCtx.TagRepository.FindTagList(reqCtx, page)
}
