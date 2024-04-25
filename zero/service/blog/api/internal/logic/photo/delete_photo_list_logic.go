package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除照片
func NewDeletePhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoListLogic {
	return &DeletePhotoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePhotoListLogic) DeletePhotoList(reqCtx *types.RestHeader, req *types.IdsReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
