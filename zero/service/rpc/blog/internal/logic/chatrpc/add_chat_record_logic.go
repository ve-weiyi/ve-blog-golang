package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddChatRecordLogic {
	return &AddChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建聊天记录
func (l *AddChatRecordLogic) AddChatRecord(in *blog.ChatRecord) (*blog.ChatRecord, error) {
	entity := convert.ConvertChatRecordPbToModel(in)

	_, err := l.svcCtx.ChatRecordModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertChatRecordModelToPb(entity), nil
}
