package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除照片
func NewDeletePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoLogic {
	return &DeletePhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePhotoLogic) DeletePhoto(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
