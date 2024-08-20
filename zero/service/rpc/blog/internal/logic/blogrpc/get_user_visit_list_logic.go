package blogrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserVisitListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVisitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVisitListLogic {
	return &GetUserVisitListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传文件
func (l *GetUserVisitListLogic) GetUserVisitList(in *blog.EmptyReq) (*blog.UserVisitPageRsp, error) {
	result, err := l.svcCtx.VisitHistoryModel.FindList(l.ctx, 1, 100, "id desc", "")
	if err != nil {
		return nil, err
	}

	var list []*blog.UserVisit
	for _, v := range result {
		uv := &blog.UserVisit{
			Date:      v.CreatedAt.Format(time.DateTime),
			ViewCount: v.ViewsCount,
		}
		list = append(list, uv)
	}

	return &blog.UserVisitPageRsp{
		List: list,
	}, nil
}
