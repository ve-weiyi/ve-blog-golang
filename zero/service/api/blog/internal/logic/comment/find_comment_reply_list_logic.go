package comment

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"
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
	in := convert.ConvertCommentQueryPb(req)

	// 查找评论列表
	out, err := l.svcCtx.CommentRpc.FindCommentList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查找评论数量
	total, err := l.svcCtx.CommentRpc.FindCommentCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查找评论回复列表
	var list []*types.Comment
	for _, v := range out.List {
		m := convert.ConvertCommentTypes(v)
		// 查询回复评论
		reply, _ := l.svcCtx.CommentRpc.FindCommentList(l.ctx, &blog.PageQuery{
			Page:       1,
			PageSize:   3,
			Sorts:      in.Sorts,
			Conditions: "parent_id = ?",
			Args:       []string{cast.ToString(v.Id)},
		})

		// 查询回复评论数
		replyCount, _ := l.svcCtx.CommentRpc.FindCommentCount(l.ctx, &blog.PageQuery{
			Conditions: "parent_id = ?",
			Sorts:      in.Sorts,
			Args:       []string{cast.ToString(v.Id)},
		})

		for _, r := range reply.List {
			m.CommentReplyList = append(m.CommentReplyList, convert.ConvertCommentReplyTypes(r))
		}
		m.ReplyCount = replyCount.Count
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
