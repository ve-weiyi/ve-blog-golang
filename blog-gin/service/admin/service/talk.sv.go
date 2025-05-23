package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type TalkService struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkService(svcCtx *svctx.ServiceContext) *TalkService {
	return &TalkService{
		svcCtx: svcCtx,
	}
}

// 创建说说
func (s *TalkService) AddTalk(reqCtx *request.Context, in *dto.TalkNewReq) (out *dto.TalkBackVO, err error) {
	// todo

	return
}

// 删除说说
func (s *TalkService) DeleteTalk(reqCtx *request.Context, in *dto.IdReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取说说列表
func (s *TalkService) FindTalkList(reqCtx *request.Context, in *dto.TalkQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 查询说说
func (s *TalkService) GetTalk(reqCtx *request.Context, in *dto.IdReq) (out *dto.TalkBackVO, err error) {
	// todo

	return
}

// 更新说说
func (s *TalkService) UpdateTalk(reqCtx *request.Context, in *dto.TalkNewReq) (out *dto.TalkBackVO, err error) {
	// todo

	return
}
