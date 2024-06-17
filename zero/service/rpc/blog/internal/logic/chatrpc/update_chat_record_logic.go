package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChatRecordLogic {
	return &UpdateChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新聊天记录
func (l *UpdateChatRecordLogic) UpdateChatRecord(in *blog.ChatRecord) (*blog.ChatRecord, error) {
	entity := convert.ConvertChatRecordPbToModel(in)

	_, err := l.svcCtx.ChatRecordModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertChatRecordModelToPb(entity), nil
}
