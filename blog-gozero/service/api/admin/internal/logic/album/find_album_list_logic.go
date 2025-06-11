package album

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取相册列表
func NewFindAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlbumListLogic {
	return &FindAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAlbumListLogic) FindAlbumList(req *types.AlbumQuery) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindAlbumListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Sorts:     req.Sorts,
		AlbumName: req.AlbumName,
		IsDelete:  req.IsDelete,
	}

	out, err := l.svcCtx.ResourceRpc.FindAlbumList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.AlbumBackVO
	for _, v := range out.List {
		m := ConvertAlbumTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
