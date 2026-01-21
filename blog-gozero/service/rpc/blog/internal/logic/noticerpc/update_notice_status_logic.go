package noticerpclogic

import (
	"context"
	"database/sql"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/enums"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeStatusLogic {
	return &UpdateNoticeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新通知状态
func (l *UpdateNoticeStatusLogic) UpdateNoticeStatus(in *noticerpc.UpdateNoticeStatusReq) (*noticerpc.UpdateNoticeStatusResp, error) {
	// 准备更新字段
	columns := map[string]interface{}{
		"publish_status": in.PublishStatus,
	}

	// 根据状态设置时间
	now := time.Now()
	if in.PublishStatus == enums.NoticeStatusPublished {
		columns["publish_time"] = sql.NullTime{Time: now, Valid: true}
	} else if in.PublishStatus == enums.NoticeStatusRevoked {
		columns["revoke_time"] = sql.NullTime{Time: now, Valid: true}
	}

	_, err := l.svcCtx.TSystemNoticeModel.Updates(l.ctx, columns, "id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	// 查询更新后的数据
	entity, err := l.svcCtx.TSystemNoticeModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &noticerpc.UpdateNoticeStatusResp{
		Notice: convertNoticeOut(entity),
	}, nil
}
