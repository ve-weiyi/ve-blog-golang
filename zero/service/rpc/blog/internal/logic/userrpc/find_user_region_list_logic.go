package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserRegionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserRegionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserRegionListLogic {
	return &FindUserRegionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户分布区域
func (l *FindUserRegionListLogic) FindUserRegionList(in *blog.EmptyReq) (*blog.FindUserRegionListResp, error) {
	// todo: add your logic here and delete this line

	return &blog.FindUserRegionListResp{}, nil
}
