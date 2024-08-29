package comment

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"
)

type FindCommentBackListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户评论列表
func NewFindCommentBackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentBackListLogic {
	return &FindCommentBackListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentBackListLogic) FindCommentBackList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := ConvertPageQuery(req)
	out, err := l.svcCtx.CommentRpc.FindCommentList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.CommentRpc.FindCommentCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.CommentBackDTO
	for _, v := range out.List {
		m := ConvertCommentBackTypes(v)
		user, _ := l.svcCtx.AccountRpc.GetUserInfo(l.ctx, &accountrpc.UserIdReq{UserId: v.UserId})
		if user != nil {
			m.Nickname = user.Nickname
			m.Avatar = user.Avatar
		}

		if v.Type == 1 {
			article, _ := l.svcCtx.ArticleRpc.GetArticle(l.ctx, &articlerpc.IdReq{Id: v.TopicId})
			if article != nil {
				m.TopicTitle = article.ArticleTitle
			}
		}

		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
