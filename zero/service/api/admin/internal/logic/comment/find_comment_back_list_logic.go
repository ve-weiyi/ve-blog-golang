package comment

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/commentrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentBackListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询评论列表(后台)
func NewFindCommentBackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentBackListLogic {
	return &FindCommentBackListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentBackListLogic) FindCommentBackList(req *types.CommentQuery) (resp *types.PageResp, err error) {
	in := &commentrpc.FindCommentListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Sorts:     req.Sorts,
		SessionId: 0,
		Type:      req.Type,
	}

	// 查找评论列表
	out, err := l.svcCtx.CommentRpc.FindCommentList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var uids []int64
	for _, v := range out.List {
		uids = append(uids, v.UserId)
		uids = append(uids, v.ReplyUserId)
	}

	// 查询用户信息
	users, err := l.svcCtx.AccountRpc.FindUserList(l.ctx, &accountrpc.FindUserListReq{
		UserIds: uids,
	})
	if err != nil {
		return nil, err
	}

	usm := make(map[int64]*accountrpc.UserInfoResp)
	for _, v := range users.List {
		usm[v.UserId] = v
	}

	// 查找评论回复列表
	var list []*types.CommentBackDTO
	for _, v := range out.List {
		m := ConvertCommentTypes(v, usm)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertCommentTypes(in *commentrpc.CommentDetails, usm map[int64]*accountrpc.UserInfoResp) (out *types.CommentBackDTO) {
	out = &types.CommentBackDTO{
		Id:             in.Id,
		TopicId:        in.TopicId,
		ParentId:       in.ParentId,
		SessionId:      in.SessionId,
		UserId:         in.UserId,
		ReplyUserId:    in.ReplyUserId,
		CommentContent: in.CommentContent,
		Type:           in.Type,
		CreatedAt:      in.CreatedAt,
		LikeCount:      in.LikeCount,
		ReplyCount:     0,
	}

	// 用户信息
	if out.UserId != 0 {
		user, ok := usm[out.UserId]
		if ok && user != nil {
			out.User = ConvertCommentUserInfoToPb(user)
		}
	}
	// 回复用户信息
	if out.ReplyUserId != 0 {
		user, ok := usm[out.ReplyUserId]
		if ok && user != nil {
			out.ReplyUser = ConvertCommentUserInfoToPb(user)
		}
	}

	return
}

func ConvertCommentUserInfoToPb(in *accountrpc.UserInfoResp) (out *types.CommentUserInfo) {
	return &types.CommentUserInfo{
		Id:       in.UserId,
		Nickname: in.Nickname,
		Avatar:   in.Avatar,
		Website:  in.Info,
	}
}
