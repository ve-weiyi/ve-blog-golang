package file

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFileFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文件目录
func NewAddFileFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFileFolderLogic {
	return &AddFileFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFileFolderLogic) AddFileFolder(req *types.FileFolderNewReq) (resp *types.FileFolderBackDTO, err error) {
	in := ConvertFileFolderPb(req)
	in.UserId = cast.ToInt64(l.ctx.Value("uid"))

	out, err := l.svcCtx.ResourceRpc.AddFileFolder(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertFileFolderTypes(out)
	return resp, nil
}
