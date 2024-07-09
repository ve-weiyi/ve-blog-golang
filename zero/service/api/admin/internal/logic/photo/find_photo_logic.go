package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

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

func (l *FindPhotoLogic) FindPhoto(req *types.IdReq) (resp *types.Photo, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.PhotoRpc.FindPhoto(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoTypes(out), nil
}
