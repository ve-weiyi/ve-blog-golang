package file

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileUploadListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文件上传列表
func NewFindFileUploadListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileUploadListLogic {
	return &FindFileUploadListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFileUploadListLogic) FindFileUploadList(req *types.FileUploadQuery) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindFileUploadListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
		FilePath: req.FilePath,
		FileType: req.FileType,
	}

	out, err := l.svcCtx.ResourceRpc.FindFileUploadList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.FileUploadBackDTO
	for _, v := range out.List {
		m := ConvertFileUploadTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
