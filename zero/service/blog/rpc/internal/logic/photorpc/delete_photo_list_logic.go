package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoListLogic {
	return &DeletePhotoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除照片
func (l *DeletePhotoListLogic) DeletePhotoList(in *blog.IdsReq) (*blog.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &blog.BatchResp{}, nil
}
