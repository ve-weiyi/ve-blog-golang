package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询标签
func NewFindTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagLogic {
	return &FindTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTagLogic) FindTag(req *types.IdReq) (resp *types.Tag, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.TagRpc.FindTag(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertTagTypes(out), nil
}
