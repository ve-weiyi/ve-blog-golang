package comment

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

type FindCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询评论列表
func NewFindCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentListLogic {
	return &FindCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentListLogic) FindCommentList(req *types.CommentQueryReq) (resp *types.PageResp, err error) {
	in := convert.ConvertCommentQueryTypes(req)
	out, err := l.svcCtx.CommentRpc.FindCommentReplyList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.CommentDTO
	for _, v := range out.List {
		m := convert.ConvertCommentDTOTypes(v)
		// 查询回复评论
		reply, _ := l.svcCtx.CommentRpc.FindCommentReplyList(l.ctx, &blog.PageQuery{
			Page:       1,
			PageSize:   3,
			Sorts:      in.Sorts,
			Conditions: "parent_id = ?",
			Args:       []string{cast.ToString(v.Id)},
		})

		for _, r := range reply.List {
			m.CommentReplyList = append(m.CommentReplyList, convert.ConvertCommentReplyTypes(r))
		}
		// 查询回复评论数
		replyCount, _ := l.svcCtx.CommentRpc.FindCommentCount(l.ctx, &blog.PageQuery{
			Conditions: "parent_id = ?",
			Sorts:      in.Sorts,
			Args:       []string{cast.ToString(v.Id)},
		})
		m.ReplyCount = replyCount.Count
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
