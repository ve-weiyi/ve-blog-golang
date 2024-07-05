package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRemarkLogic {
	return &CreateRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建留言
func (l *CreateRemarkLogic) CreateRemark(in *blog.Remark) (*blog.Remark, error) {
	entity := convert.ConvertRemarkPbToModel(in)

	_, err := l.svcCtx.RemarkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertRemarkModelToPb(entity), nil
}
