package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRemarkLogic {
	return &AddRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建留言
func (l *AddRemarkLogic) AddRemark(in *blog.Remark) (*blog.Remark, error) {
	entity := convert.ConvertRemarkPbToModel(in)

	_, err := l.svcCtx.RemarkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertRemarkModelToPb(entity), nil
}
