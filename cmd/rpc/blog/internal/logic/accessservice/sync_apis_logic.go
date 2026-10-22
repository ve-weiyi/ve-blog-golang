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

type SyncApisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncApisLogic {
	return &SyncApisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步接口列表
func (l *SyncApisLogic) SyncApis(in *accessrpc.SyncApisRequest) (*accessrpc.SyncApisResponse, error) {
	count, err := l.syncApiList(in)
	if err != nil {
		return nil, err
	}

	return &accessrpc.SyncApisResponse{
		SuccessCount: count,
	}, nil
}

func (l *SyncApisLogic) syncApiList(in *accessrpc.SyncApisRequest) (count int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("同步接口 panic: %v", r)
		}
	}()

	for _, item := range in.Apis {
		if err := l.InsertApi(l.ctx, item); err != nil {
			return count, err
		}
		count++
	}

	l.Infof("成功同步接口")
	return count, nil
}

func (l *SyncApisLogic) InsertApi(ctx context.Context, item *accessrpc.CreateApiRequest) (err error) {
	if item == nil {
		return fmt.Errorf("invalid api")
	}

	// 已存在则跳过
	parent, findErr := l.svcCtx.TApiModel.FindOneByPathMethodName(ctx, item.Path, item.Method, item.Name)
	if findErr != nil && !errors.Is(findErr, gorm.ErrRecordNotFound) {
		return findErr
	}
	if parent == nil {
		// 插入数据
		parent = convertApiIn(item)
		_, err = l.svcCtx.TApiModel.Insert(ctx, parent)
		if err != nil {
			return err
		}
	}

	for _, child := range item.Children {
		child.ParentId = parent.Id
		err = l.InsertApi(ctx, child)
		if err != nil {
			return err
		}
	}

	return nil
}
