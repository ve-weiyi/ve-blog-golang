package upload

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oss"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
)

type MultiUploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件列表
func NewMultiUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MultiUploadFileLogic {
	return &MultiUploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MultiUploadFileLogic) MultiUploadFile(req *types.MultiUploadFileReq, r *http.Request) (resp []*types.FileInfoVO, err error) {
	uid := cast.ToString(l.ctx.Value(restx.HeaderUid))

	// 获取文件切片
	files := r.MultipartForm.File["files"]
	for _, h := range files {
		f, err := h.Open()
		if err != nil {
			return nil, err
		}

		up, err := l.svcCtx.Uploader.UploadFile(f, req.FilePath, oss.NewFileNameWithDateTime(h.Filename))
		if err != nil {
			return nil, err
		}

		in := &syslogrpc.UploadLogNewReq{
			UserId:   uid,
			FilePath: req.FilePath,
			FileName: h.Filename,
			FileType: filepath.Ext(h.Filename),
			FileSize: h.Size,
			FileMd5:  crypto.Md5v(h.Filename, ""),
			FileUrl:  up,
		}

		out, err := l.svcCtx.SyslogRpc.AddUploadLog(l.ctx, in)
		if err != nil {
			return nil, err
		}

		resp = append(resp, &types.FileInfoVO{
			FilePath:  req.FilePath,
			FileName:  out.FileName,
			FileType:  out.FileType,
			FileSize:  out.FileSize,
			FileUrl:   out.FileUrl,
			UpdatedAt: out.UpdatedAt,
		})
	}

	return resp, nil
}
