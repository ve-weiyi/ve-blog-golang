package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindBannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindBannerListLogic {
	return &FindBannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询页面列表
func (l *FindBannerListLogic) FindBannerList(in *photorpc.FindBannerListReq) (*photorpc.FindBannerListResp, error) {
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

	result, err := l.svcCtx.BannerModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*photorpc.BannerDetails
	for _, v := range result {
		list = append(list, ConvertBannerOut(v))
	}

	return &photorpc.FindBannerListResp{
		List: list,
	}, nil
}
