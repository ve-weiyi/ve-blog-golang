package tag

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/contentservice"
)

type UpdateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新标签
func NewUpdateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTagLogic {
	return &UpdateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTagLogic) UpdateTag(req *types.UpdateTagReq) (resp *types.TagVO, err error) {
	out, err := l.svcCtx.ContentService.UpdateTag(l.ctx, &contentservice.UpdateTagRequest{
		Id:      req.Id,
		TagName: req.TagName,
	})
	if err != nil {
		return nil, err
	}

	return &types.TagVO{
		Id:           out.Tag.Id,
		TagName:      out.Tag.TagName,
		ArticleCount: out.Tag.ArticleCount,
		CreatedAt:    out.Tag.CreatedAt,
		UpdatedAt:    out.Tag.UpdatedAt,
	}, nil
}
