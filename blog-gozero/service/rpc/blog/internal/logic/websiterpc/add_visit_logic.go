package websiterpclogic

import (
	"context"
	"fmt"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVisitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVisitLogic {
	return &AddVisitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加用户访问记录
func (l *AddVisitLogic) AddVisit(in *websiterpc.AddVisitReq) (*websiterpc.AddVisitResp, error) {
	visitor := in.Visitor
	if visitor == "" {
		return nil, fmt.Errorf("visitor is empty")
	}

	key := rediskey.GetBlogDailyVisitorKey(time.Now().Format(time.DateOnly))
	ok, err := l.svcCtx.Redis.SIsMember(l.ctx, key, visitor).Result()
	if err != nil {
		return nil, err
	}

	if !ok {
		// 添加当天的访问量+1
		day := time.Now().Format("2006-01-02")

		result := l.svcCtx.Gorm.Exec("UPDATE t_visit_history SET views_count = views_count + 1 WHERE date = ?", day)
		if result.RowsAffected == 0 {
			_, err := l.svcCtx.TVisitHistoryModel.Insert(l.ctx, &model.TVisitHistory{
				Date:       day,
				ViewsCount: 1,
			})
			if err != nil {
				return nil, err
			}
		}

		// 添加总访问量
		totalKey := rediskey.GetBlogTotalViewCountKey()
		_, err = l.svcCtx.Redis.Incr(l.ctx, totalKey).Result()
		if err != nil {
			return nil, err
		}

		// 添加总访问量
		dailyKey := rediskey.GetBlogViewCountKey()
		_, err = l.svcCtx.Redis.ZIncrBy(l.ctx, dailyKey, 1, day).Result()
		if err != nil {
			return nil, err
		}

		// 保存访客标识
		_, err = l.svcCtx.Redis.SAdd(l.ctx, key, visitor).Result()
		if err != nil {
			return nil, err
		}
	}

	return &websiterpc.AddVisitResp{Count: 0}, nil
}
