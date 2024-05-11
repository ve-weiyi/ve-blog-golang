package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
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

func (l *FindCommentReplyListLogic) FindCommentReplyList(reqCtx *types.RestHeader, req *types.CommentQueryReq) (resp *types.PageResp, err error) {
	in := convert.ConvertCommentQueryTypes(req)
	out, err := l.svcCtx.CommentRpc.FindCommentReplyList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.CommentReply
	for _, v := range out.List {
		m := convert.ConvertCommentReplyTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
