package websiterpclogic

import (
	"context"
	"strings"
	"time"

	"github.com/mssola/useragent"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/rpcutil"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
)

type ReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上报
func (l *ReportLogic) Report(in *websiterpc.EmptyReq) (*websiterpc.ReportResp, error) {
	visitor, err := l.GetIdentity()
	if err != nil {
		return nil, err
	}

	key := rediskey.GetBlogVisitorKey(time.Now().Format(time.DateOnly))
	ok, err := l.svcCtx.Redis.SIsMember(l.ctx, key, visitor).Result()
	if err != nil {
		return nil, err
	}

	if !ok {
		// 数据库访问量+1
		l.AddVisit()

		// 添加总访问量
		totalKey := rediskey.GetBlogViewCountKey()
		_, err = l.svcCtx.Redis.Incr(l.ctx, totalKey).Result()
		if err != nil {
			return nil, err
		}

		// 保存访客标识
		_, err = l.svcCtx.Redis.SAdd(l.ctx, key, visitor).Result()
		if err != nil {
			return nil, err
		}
	}

	return &websiterpc.ReportResp{TerminalId: visitor}, nil
}

func (l *ReportLogic) AddVisit() error {
	// 添加当天的访问量
	day := time.Now().Format("2006-01-02")

	result := l.svcCtx.Gorm.Exec("UPDATE t_visit_history SET views_count = views_count + 1 WHERE date = ?", day)
	if result.RowsAffected == 0 {
		_, err := l.svcCtx.TVisitHistoryModel.Insert(l.ctx, &model.TVisitHistory{
			Date:       day,
			ViewsCount: 1,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// 获取身份标识
func (l *ReportLogic) GetIdentity() (string, error) {
	ci, err := rpcutil.GetRPCClientIP(l.ctx)
	if err != nil {
		return "", err
	}

	ua, err := rpcutil.GetRPCUserAgent(l.ctx)
	if err != nil {
		return "", err
	}

	// 分割字符串，提取 IP 部分
	ip := strings.Split(ci, ":")[0]
	browser, _ := useragent.New(ua).Browser()
	os := useragent.New(ua).OS()

	terminal := crypto.Md5v(ip+browser+os, "")

	return terminal, nil
}
