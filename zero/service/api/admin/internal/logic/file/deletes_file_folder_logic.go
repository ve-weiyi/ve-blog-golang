package file

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesFileFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文件目录
func NewDeletesFileFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesFileFolderLogic {
	return &DeletesFileFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesFileFolderLogic) DeletesFileFolder(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &resourcerpc.IdsReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.ResourceRpc.DeleteFileUpload(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
