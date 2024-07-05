package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChatRecordLogic {
	return &CreateChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建聊天记录
func (l *CreateChatRecordLogic) CreateChatRecord(in *blog.ChatRecord) (*blog.ChatRecord, error) {
	entity := convert.ConvertChatRecordPbToModel(in)

	_, err := l.svcCtx.ChatRecordModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertChatRecordModelToPb(entity), nil
}
