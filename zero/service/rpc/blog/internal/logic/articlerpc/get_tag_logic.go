package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagLogic {
	return &GetTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询标签
func (l *GetTagLogic) GetTag(in *blog.IdReq) (*blog.TagDetails, error) {
	// todo: add your logic here and delete this line

	return &blog.TagDetails{}, nil
}
