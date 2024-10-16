package file

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新文件目录
func NewUpdateFileFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileFolderLogic {
	return &UpdateFileFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFileFolderLogic) UpdateFileFolder(req *types.FileFolderNewReq) (resp *types.FileFolderBackDTO, err error) {
	in := ConvertFileFolderPb(req)
	in.UserId = cast.ToInt64(l.ctx.Value("uid"))

	out, err := l.svcCtx.ResourceRpc.UpdateFileFolder(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertFileFolderTypes(out)
	return resp, nil
}
