package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新菜单
func (l *UpdateMenuLogic) UpdateMenu(in *blog.Menu) (*blog.Menu, error) {
	entity := convert.ConvertMenuPbToModel(in)

	_, err := l.svcCtx.MenuModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertMenuModelToPb(entity), nil
}
