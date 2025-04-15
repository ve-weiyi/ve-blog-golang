package websiterpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserTotalVisitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserTotalVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserTotalVisitLogic {
	return &GetUserTotalVisitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户总流量数
func (l *GetUserTotalVisitLogic) GetUserTotalVisit(in *websiterpc.EmptyReq) (*websiterpc.CountResp, error) {
	totalKey := rediskey.GetBlogTotalViewCountKey()
	total, err := l.svcCtx.Redis.Get(l.ctx, totalKey).Int64()
	if err != nil {
		l.Errorf("获取用户总访问量失败: %v", err)
		records, err := l.svcCtx.TVisitHistoryModel.FindALL(l.ctx, "")
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			total += record.ViewsCount
		}

		_, err = l.svcCtx.Redis.Set(l.ctx, totalKey, total, 0).Result()
	}

	return &websiterpc.CountResp{
		Count: total,
	}, nil
}
