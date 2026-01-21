package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type NoticeLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewNoticeLogic(svcCtx *svctx.ServiceContext) *NoticeLogic {
	return &NoticeLogic{
		svcCtx: svcCtx,
	}
}

// 创建通知
func (s *NoticeLogic) AddNotice(reqCtx *request.Context, in *types.AddNoticeReq) (out *types.NoticeBackVO, err error) {
	// todo

	return
}

// 删除通知
func (s *NoticeLogic) DeletesNotice(reqCtx *request.Context, in *types.IdsReq) (out *types.BatchResp, err error) {
	// todo

	return
}

// 分页获取通知列表
func (s *NoticeLogic) FindNoticeList(reqCtx *request.Context, in *types.QueryNoticeReq) (out *types.PageResp, err error) {
	// todo

	return
}

// 查询用户通知列表
func (s *NoticeLogic) FindUserNoticeList(reqCtx *request.Context, in *types.QueryUserNoticeReq) (out *types.PageResp, err error) {
	// todo

	return
}

// 查询通知详情
func (s *NoticeLogic) GetNotice(reqCtx *request.Context, in *types.IdReq) (out *types.NoticeBackVO, err error) {
	// todo

	return
}

// 更新通知
func (s *NoticeLogic) UpdateNotice(reqCtx *request.Context, in *types.UpdateNoticeReq) (out *types.NoticeBackVO, err error) {
	// todo

	return
}

// 更新通知状态
func (s *NoticeLogic) UpdateNoticeStatus(reqCtx *request.Context, in *types.UpdateNoticeStatusReq) (out *types.NoticeBackVO, err error) {
	// todo

	return
}
