package photo

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

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

func (l *DeletePhotoLogic) DeletePhoto(req *types.IdReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.PhotoRpc.DeletePhoto(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
