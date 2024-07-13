package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTagLogic {
	return &AddTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建标签
func (l *AddTagLogic) AddTag(in *blog.Tag) (*blog.Tag, error) {
	entity := convert.ConvertTagPbToModel(in)

	_, err := l.svcCtx.TagModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTagModelToPb(entity), nil
}
