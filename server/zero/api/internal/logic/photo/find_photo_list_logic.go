package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindPhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoListLogic {
	return &FindPhotoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoListLogic) FindPhotoList(req *types.PageQuery) (resp []types.Photo, err error) {
	// todo: add your logic here and delete this line

	return
}
