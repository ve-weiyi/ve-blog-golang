package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/commentrpc"
)

type FindCommentReplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询评论回复列表
func NewFindCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentReplyListLogic {
	return &FindCommentReplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentReplyListLogic) FindCommentReplyList(req *types.CommentQueryReq) (resp *types.PageResp, err error) {
	in := &commentrpc.FindCommentReplyListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Sorts:     "",
		TopicId:   req.TopicId,
		ParentId:  req.ParentId,
		SessionId: 0,
		Type:      req.Type,
	}

	// 查找评论列表
	out, err := l.svcCtx.CommentRpc.FindCommentReplyList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var userIds []int64
	for _, v := range out.List {
		userIds = append(userIds, v.UserId)
		userIds = append(userIds, v.ReplyUserId)
	}

	// 查询用户信息
	users, err := l.svcCtx.AccountRpc.FindUserList(l.ctx, &accountrpc.FindUserListReq{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}

	usm := make(map[int64]*accountrpc.UserDetails)
	for _, v := range users.List {
		usm[v.UserId] = v
	}

	// 查找评论回复列表
	var list []*types.CommentReply
	for _, v := range out.List {
		m := ConvertCommentReplyTypes(v, usm)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
