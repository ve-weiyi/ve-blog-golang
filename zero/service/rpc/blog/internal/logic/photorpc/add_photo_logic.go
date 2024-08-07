package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPhotoLogic {
	return &AddPhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建照片
func (l *AddPhotoLogic) AddPhoto(in *photorpc.PhotoNew) (*photorpc.PhotoDetails, error) {
	entity := ConvertPhotoIn(in)

	_, err := l.svcCtx.PhotoModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertPhotoOut(entity), nil
}
