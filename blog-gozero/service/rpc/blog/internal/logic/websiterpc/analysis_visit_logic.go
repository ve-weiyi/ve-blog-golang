package websiterpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisVisitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisVisitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisVisitLogic {
	return &AnalysisVisitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户日浏览量分析
func (l *AnalysisVisitLogic) AnalysisVisit(in *websiterpc.EmptyReq) (*websiterpc.AnalysisVisitResp, error) {
	day := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	todayUvCount := l.getTodayUvCount(day)
	todayPvCount := l.getTodayPvCount(day)
	totalUvCount := l.getTotalUvCount()
	totalPvCount := l.getTotalPvCount()

	yesterdayUvCount := l.getTodayUvCount(yesterday)
	yesterdayPvCount := l.getTodayPvCount(yesterday)
	var uvGrowthRate, pvGrowthRate float64
	if yesterdayUvCount != 0 {
		uvGrowthRate = float64(todayUvCount-yesterdayUvCount) / float64(yesterdayUvCount) * 100
	}
	if yesterdayPvCount != 0 {
		pvGrowthRate = float64(todayPvCount-yesterdayPvCount) / float64(yesterdayPvCount) * 100
	}

	return &websiterpc.AnalysisVisitResp{
		TodayUvCount: todayUvCount,
		TotalUvCount: totalUvCount,
		UvGrowthRate: uvGrowthRate,
		TodayPvCount: todayPvCount,
		TotalPvCount: totalPvCount,
		PvGrowthRate: pvGrowthRate,
	}, nil
}

func (l *AnalysisVisitLogic) getTodayUvCount(day string) int64 {
	key := rediskey.GetUserViewCountSetKey()
	// 获取日访客数
	uvCount, err := l.svcCtx.Redis.ZScore(l.ctx, key, day).Result()
	if err != nil {
		l.Logger.Errorf("getTodayUvCount err: %v", err)
		return 0
	}
	return int64(uvCount)
}

func (l *AnalysisVisitLogic) getTodayPvCount(day string) int64 {
	key := rediskey.GetPageViewCountSetKey()
	// 获取日浏览量
	pvCount, err := l.svcCtx.Redis.ZScore(l.ctx, key, day).Result()
	if err != nil {
		l.Logger.Errorf("getTodayPvCount err: %v", err)
		return 0
	}

	return int64(pvCount)
}

func (l *AnalysisVisitLogic) getTotalUvCount() int64 {
	key := rediskey.GetTotalUserViewCountKey()
	// 获取总访客数
	uvCount, err := l.svcCtx.Redis.SCard(l.ctx, key).Result()
	if err != nil {
		l.Logger.Errorf("getTotalUvCount err: %v", err)
		records, err := l.svcCtx.TVisitDailyStatsModel.FindALL(l.ctx, "")
		if err != nil {
			return 0
		}

		for _, record := range records {
			uvCount += record.ViewCount
		}

		_, err = l.svcCtx.Redis.Set(l.ctx, key, uvCount, 0).Result()
		return 0
	}
	return uvCount
}
func (l *AnalysisVisitLogic) getTotalPvCount() int64 {
	key := rediskey.GetTotalPageViewCountKey()
	// 获取总浏览量
	pvCount, err := l.svcCtx.Redis.SCard(l.ctx, key).Result()
	if err != nil {
		l.Logger.Errorf("getTotalPvCount err: %v", err)
		records, err := l.svcCtx.TVisitDailyStatsModel.FindALL(l.ctx, "")
		if err != nil {
			return 0
		}

		for _, record := range records {
			pvCount += record.ViewCount
		}

		_, err = l.svcCtx.Redis.Set(l.ctx, key, pvCount, 0).Result()
		return 0
	}
	return pvCount
}
