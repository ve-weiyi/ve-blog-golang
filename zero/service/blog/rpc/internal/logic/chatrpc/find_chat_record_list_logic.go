package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindChatRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatRecordListLogic {
	return &FindChatRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取聊天记录列表
func (l *FindChatRecordListLogic) FindChatRecordList(in *blog.PageQuery) (*blog.ChatRecordPageResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.ChatRecordModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.ChatRecord
	for _, v := range result {
		list = append(list, convert.ConvertChatRecordModelToPb(v))
	}

	return &blog.ChatRecordPageResp{
		List: list,
	}, nil
}
