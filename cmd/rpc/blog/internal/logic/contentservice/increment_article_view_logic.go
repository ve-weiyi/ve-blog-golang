package contentservicelogic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/contentrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/cachekey"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

type IncrementArticleViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIncrementArticleViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IncrementArticleViewLogic {
	return &IncrementArticleViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加文章浏览量
func (l *IncrementArticleViewLogic) IncrementArticleView(in *contentrpc.IncrementArticleViewRequest) (*contentrpc.IncrementArticleViewResponse, error) {
	// 只有公开且未删除的文章才计数，草稿/删稿的浏览量不再被刷
	rows, err := l.svcCtx.TArticleModel.UpdateFields(l.ctx, map[string]interface{}{
		"view_count": gorm.Expr("view_count + 1"),
	}, "id = ? AND status = ? AND is_delete = ?", in.Id, enums.ArticleStatusPublic, 0)
	if err != nil {
		return nil, err
	}

	// 数据库计数已提交，缓存失败只影响展示。此处返回错误会让客户端重试，造成重复计数
	if rows > 0 {
		key := cachekey.ArticleViewCountKey
		if err := l.svcCtx.Redis.ZIncrBy(l.ctx, key, 1, fmt.Sprintf("%d", in.Id)).Err(); err != nil {
			l.Errorf("IncrementArticleView ZIncrBy error: %v", err)
		}
	}

	return &contentrpc.IncrementArticleViewResponse{Success: true}, nil
}
