package accessservicelogic

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/accessrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
)

type SyncMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncMenusLogic {
	return &SyncMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步菜单列表
func (l *SyncMenusLogic) SyncMenus(in *accessrpc.SyncMenusRequest) (*accessrpc.SyncMenusResponse, error) {
	count, err := l.syncMenuList(in)
	if err != nil {
		return nil, err
	}

	return &accessrpc.SyncMenusResponse{
		SuccessCount: count,
	}, nil
}

func (l *SyncMenusLogic) syncMenuList(in *accessrpc.SyncMenusRequest) (count int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("同步菜单 panic: %v", r)
		}
	}()

	for _, item := range in.Menus {
		if err := l.InsertMenu(l.ctx, item); err != nil {
			return count, err
		}
		count++
	}

	l.Infof("成功同步菜单")
	return count, nil
}

func (l *SyncMenusLogic) InsertMenu(ctx context.Context, item *accessrpc.CreateMenuRequest) (err error) {
	if item == nil || item.Meta == nil {
		return fmt.Errorf("invalid menu")
	}

	// 已存在则跳过
	parent, findErr := l.svcCtx.TMenuModel.FindOneByPathPerm(ctx, item.Path, item.Meta.Perm)
	if findErr != nil && !errors.Is(findErr, gorm.ErrRecordNotFound) {
		return findErr
	}
	if parent == nil {
		// 插入数据
		parent = convertMenuIn(item)
		_, err = l.svcCtx.TMenuModel.Insert(ctx, parent)
		if err != nil {
			return err
		}
	}

	for _, child := range item.Children {
		child.ParentId = parent.Id
		err = l.InsertMenu(ctx, child)
		if err != nil {
			return err
		}
	}

	return nil
}
