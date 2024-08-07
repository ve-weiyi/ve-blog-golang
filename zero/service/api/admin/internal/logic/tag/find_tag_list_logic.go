package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取标签列表
func NewFindTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTagListLogic {
	return &FindTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTagListLogic) FindTagList(req *types.TagQuery) (resp *types.PageResp, err error) {
	in := &articlerpc.FindTagListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    "id desc",
		TagName:  req.TagName,
	}

	out, err := l.svcCtx.ArticleRpc.FindTagList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.TagBackDTO
	for _, item := range out.List {
		list = append(list, ConvertTagTypes(item))
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = out.Total
	resp.List = list
	return
}

func ConvertTagTypes(in *articlerpc.TagDetails) (out *types.TagBackDTO) {
	return &types.TagBackDTO{
		Id:           in.Id,
		TagName:      in.TagName,
		ArticleCount: in.ArticleCount,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}
}
