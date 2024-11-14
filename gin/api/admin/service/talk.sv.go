package service

import (
	"github.com/ve-weiyi/ve-blog-golang/gin/api/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

type TalkService struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkService(svcCtx *svctx.ServiceContext) *TalkService {
	return &TalkService{
		svcCtx: svcCtx,
	}
}

// 分页获取说说列表
func (s *TalkService) FindTalkList(reqCtx *request.Context, in *dto.TalkQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 创建说说
func (s *TalkService) AddTalk(reqCtx *request.Context, in *dto.TalkNewReq) (out *dto.TalkBackDTO, err error) {
	// todo

	return
}

// 删除说说
func (s *TalkService) DeleteTalk(reqCtx *request.Context, in *dto.IdReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 查询说说
func (s *TalkService) GetTalk(reqCtx *request.Context, in *dto.IdReq) (out *dto.TalkBackDTO, err error) {
	// todo

	return
}

// 更新说说
func (s *TalkService) UpdateTalk(reqCtx *request.Context, in *dto.TalkNewReq) (out *dto.TalkBackDTO, err error) {
	// todo

	return
}
