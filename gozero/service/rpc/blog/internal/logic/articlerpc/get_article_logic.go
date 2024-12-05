package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章
func (l *GetArticleLogic) GetArticle(in *articlerpc.IdReq) (*articlerpc.ArticleDetails, error) {
	helper := NewArticleHelperLogic(l.ctx, l.svcCtx)
	record, err := l.svcCtx.TArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	acm, err := helper.convertArticleDetails([]*model.TArticle{record})
	if err != nil {
		return nil, err
	}

	return acm[0], nil
}
