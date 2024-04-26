package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRemarkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRemarkListLogic {
	return &FindRemarkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取留言列表
func (l *FindRemarkListLogic) FindRemarkList(in *blog.PageQuery) (*blog.RemarkPageResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.RemarkModel.FindList(l.ctx, limit, offset, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.Remark
	for _, v := range result {
		list = append(list, convert.ConvertRemarkModelToPb(v))
	}

	return &blog.RemarkPageResp{
		List: list,
	}, nil
}
