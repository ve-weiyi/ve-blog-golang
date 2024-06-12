package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除菜单
func (l *DeleteMenuLogic) DeleteMenu(in *blog.IdReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.MenuModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	rows2, err := l.svcCtx.MenuModel.DeleteBatch(l.ctx, "parent_id = ? ", in.Id)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows + rows2,
	}, nil
}
