package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentReplyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentReplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentReplyListLogic {
	return &FindCommentReplyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取评论回复列表
func (l *FindCommentReplyListLogic) FindCommentReplyList(in *blog.PageQuery) (*blog.CommentReplyPageResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.CommentModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.CommentModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.CommentReply
	for _, v := range result {
		m := convert.ConvertCommentReplyPb(v)

		// 用户信息
		if v.UserId != 0 {
			user, _ := l.svcCtx.UserInformationModel.First(l.ctx, "user_id = ?", v.UserId)
			if user != nil {
				m.User = convert.ConvertUserInfoModelToPb(user)
			}
		}
		// 回复用户信息
		if v.ReplyUserId != 0 {
			user, _ := l.svcCtx.UserInformationModel.First(l.ctx, "user_id = ?", v.ReplyUserId)
			if user != nil {
				m.ReplyUser = convert.ConvertUserInfoModelToPb(user)
			}
		}

		m.LikeCount = 10
		list = append(list, m)
	}

	return &blog.CommentReplyPageResp{
		Total: count,
		List:  list,
	}, nil
}
