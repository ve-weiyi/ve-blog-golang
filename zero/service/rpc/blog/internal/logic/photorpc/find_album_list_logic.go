package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlbumListLogic {
	return &FindAlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询相册列表
func (l *FindAlbumListLogic) FindAlbumList(in *photorpc.FindAlbumListReq) (*photorpc.FindAlbumListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)
	sorts = in.Sorts

	result, err := l.svcCtx.AlbumModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*photorpc.AlbumDetails
	for _, v := range result {
		list = append(list, ConvertAlbumOut(v))
	}

	return &photorpc.FindAlbumListResp{
		List: list,
	}, nil
}
