package album

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"
)

type FindPhotoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取照片列表
func NewFindPhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoListLogic {
	return &FindPhotoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPhotoListLogic) FindPhotoList(req *types.PhotoQuery) (resp *types.PageResp, err error) {
	in := &blogrpc.PageQuery{
		Conditions: fmt.Sprintf("album_id = %v", req.AlbumId),
	}
	out, err := l.svcCtx.PhotoRpc.FindPhotoList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.PhotoRpc.FindPhotoCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.Photo
	for _, v := range out.List {
		m := convert.ConvertPhotoTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
