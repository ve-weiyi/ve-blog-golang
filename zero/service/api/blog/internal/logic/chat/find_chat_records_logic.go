package chat

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindChatRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询聊天记录
func NewFindChatRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindChatRecordsLogic {
	return &FindChatRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindChatRecordsLogic) FindChatRecords(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.ChatRpc.FindChatRecordList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ChatRecord
	for _, v := range out.List {
		list = append(list, convert.ConvertChatRecordTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
