package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminHomeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取后台首页信息
func NewGetAdminHomeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminHomeInfoLogic {
	return &GetAdminHomeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminHomeInfoLogic) GetAdminHomeInfo(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.AdminHomeInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
