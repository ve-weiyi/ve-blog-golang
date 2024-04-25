package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询照片
func NewFindPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoLogic {
	return &FindPhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoLogic) FindPhoto(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.Photo, err error) {
	// todo: add your logic here and delete this line

	return
}
