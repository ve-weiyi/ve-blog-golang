package comment

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论
func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.CommentNewReq) (resp *types.CommentNewReq, err error) {
	in := convert.ConvertCommentPb(req)
	// l.ctx.Value("uid")
	in.UserId = cast.ToInt64(l.ctx.Value("uid"))
	out, err := l.svcCtx.CommentRpc.AddComment(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convert.ConvertCommentTypes(out)
	return resp, nil
}
