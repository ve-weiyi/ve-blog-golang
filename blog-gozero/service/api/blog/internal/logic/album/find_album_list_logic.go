package album

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"
)

type FindAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取相册列表
func NewFindAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlbumListLogic {
	return &FindAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindAlbumListLogic) FindAlbumList(req *types.AlbumQueryReq) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindAlbumListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
	}
	out, err := l.svcCtx.ResourceRpc.FindAlbumList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.Album, 0)
	for _, v := range out.List {
		m := ConvertAlbumTypes(v)
		list = append(list, m)
	}

	_, err = l.svcCtx.SyslogRpc.AddVisitLog(l.ctx, &syslogrpc.VisitLogNewReq{
		PageName: "相册",
	})
	if err != nil {
		return nil, err
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = int64(len(list))
	resp.List = list
	return resp, nil
}

func ConvertAlbumTypes(req *resourcerpc.AlbumDetails) (out *types.Album) {

	return &types.Album{
		Id:         req.Id,
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
	}
}
