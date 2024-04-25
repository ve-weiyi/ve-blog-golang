package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindTalkDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取说说详情列表
func NewFindTalkDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkDetailLogic {
	return &FindTalkDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTalkDetailLogic) FindTalkDetail(req *types.IdReq) (resp *types.TalkDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
