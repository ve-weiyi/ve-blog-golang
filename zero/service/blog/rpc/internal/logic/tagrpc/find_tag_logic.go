package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagLogic {
	return &FindTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询标签
func (l *FindTagLogic) FindTag(in *blog.IdReq) (*blog.Tag, error) {
	result, err := l.svcCtx.TagModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTagModelToPb(result), nil
}
