package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
)

type FindCommentRecentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询最新评论回复列表
func NewFindCommentRecentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentRecentListLogic {
	return &FindCommentRecentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentRecentListLogic) FindCommentRecentList(req *types.CommentQueryReq) (resp *types.PageResp, err error) {
	in := convert.ConvertCommentQueryPb(req)
	out, err := l.svcCtx.CommentRpc.FindCommentList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.CommentRpc.FindCommentCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.Comment
	for _, v := range out.List {
		m := convert.ConvertCommentTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
