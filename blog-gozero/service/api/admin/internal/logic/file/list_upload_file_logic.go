package file

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
)

type ListUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文件列表
func NewListUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUploadFileLogic {
	return &ListUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUploadFileLogic) ListUploadFile(req *types.ListUploadFileReq) (resp *types.ListUploadFileResp, err error) {
	logx.Info("ListUploadFileLogic ListUploadFile")
	up, err := l.svcCtx.Uploader.ListFiles(req.FilePath, int(req.Limit))
	if err != nil {
		return nil, err
	}

	resp = &types.ListUploadFileResp{
		Urls: up,
	}
	return
}
