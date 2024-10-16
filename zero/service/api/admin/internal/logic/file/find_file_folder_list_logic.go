package file

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileFolderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文件目录列表
func NewFindFileFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileFolderListLogic {
	return &FindFileFolderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFileFolderListLogic) FindFileFolderList(req *types.FileFolderQuery) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindFileFolderListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
		FilePath: req.FilePath,
	}

	out, err := l.svcCtx.ResourceRpc.FindFileFolderList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.FileFolderBackDTO
	for _, v := range out.List {
		m := ConvertFileFolderTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
