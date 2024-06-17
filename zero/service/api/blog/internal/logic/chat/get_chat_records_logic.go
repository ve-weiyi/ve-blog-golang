package chat

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/chatrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询聊天记录
func NewGetChatRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatRecordsLogic {
	return &GetChatRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatRecordsLogic) GetChatRecords(req *types.ChatQueryReq) (resp *types.PageResp, err error) {
	in := &chatrpc.FindChatRecordListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	out, err := l.svcCtx.ChatRpc.FindChatRecordList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ChatRecord
	for _, v := range out.List {
		list = append(list, ConvertChatRecordTypes(v))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertChatRecordTypes(in *chatrpc.ChatRecordDetails) *types.ChatRecord {
	return &types.ChatRecord{
		Id:        in.Id,
		UserId:    in.UserId,
		Nickname:  in.Nickname,
		Avatar:    in.Avatar,
		Content:   in.Content,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		Type:      in.Type,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
