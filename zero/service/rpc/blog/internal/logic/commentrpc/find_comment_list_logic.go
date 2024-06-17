package commentrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindCommentListLogic {
	return &FindCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取评论列表
func (l *FindCommentListLogic) FindCommentList(in *blog.PageQuery) (*blog.CommentPageResp, error) {
	page, size, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.CommentModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.Comment
	for _, v := range result {
		m := convert.ConvertCommentModelToPb(v)
		// 用户信息
		if v.UserId != 0 {
			user, _ := l.svcCtx.UserAccountModel.FindOne(l.ctx, v.UserId)
			if user != nil {
				m.User = convert.ConvertCommentUserInfoToPb(user)
			}
		}
		// 回复用户信息
		if v.ReplyUserId != 0 {
			user, _ := l.svcCtx.UserAccountModel.FindOne(l.ctx, v.ReplyUserId)
			if user != nil {
				m.ReplyUser = convert.ConvertCommentUserInfoToPb(user)
			}
		}

		list = append(list, m)
	}

	return &blog.CommentPageResp{
		List: list,
	}, nil
}
