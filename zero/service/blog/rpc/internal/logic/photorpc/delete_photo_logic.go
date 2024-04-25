package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoLogic {
	return &DeletePhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除照片
func (l *DeletePhotoLogic) DeletePhoto(in *blog.IdReq) (*blog.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &blog.BatchResp{}, nil
}
