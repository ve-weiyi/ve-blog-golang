package noticerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/enums"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogic {
	return &AddNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建通知
func (l *AddNoticeLogic) AddNotice(in *noticerpc.AddNoticeReq) (*noticerpc.AddNoticeResp, error) {
	uid, _ := rpcutils.GetUserIdFromCtx(l.ctx)

	entity := &model.TSystemNotice{
		Title:         in.Title,
		Content:       in.Content,
		Type:          in.Type,
		Level:         in.Level,
		AppName:       in.AppName,
		PublisherId:   uid,
		PublishStatus: enums.NoticeStatusDraft,
	}

	_, err := l.svcCtx.TSystemNoticeModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &noticerpc.AddNoticeResp{
		Notice: convertNoticeOut(entity),
	}, nil
}
