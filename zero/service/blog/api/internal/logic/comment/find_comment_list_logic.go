package comment

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取评论列表
func NewFindCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentListLogic {
	return &FindCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCommentListLogic) FindCommentList(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.CommentRpc.FindCommentList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.CommentBackDTO
	for _, v := range out.List {
		m := convert.ConvertCommentBackTypes(v)
		user, _ := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &blog.UserReq{UserId: v.UserId})
		if user != nil {
			m.UserNickname = user.Nickname
			m.UserAvatar = user.Avatar
		}

		if v.Type == 1 {
			article, _ := l.svcCtx.ArticleRpc.FindArticle(l.ctx, &blog.IdReq{Id: v.TopicId})
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
