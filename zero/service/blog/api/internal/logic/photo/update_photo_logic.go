package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePhotoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新照片
func NewUpdatePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePhotoLogic {
	return &UpdatePhotoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePhotoLogic) UpdatePhoto(req *types.Photo) (resp *types.Photo, err error) {
	in := convert.ConvertPhotoPb(req)

	api, err := l.svcCtx.PhotoRpc.UpdatePhoto(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPhotoTypes(api), nil
}
