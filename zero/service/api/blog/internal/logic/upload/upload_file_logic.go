package upload

import (
	"context"
	"net/http"
	"path"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传文件
func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq, r *http.Request) (resp *types.UploadFileResp, err error) {
	f, h, _ := r.FormFile("file")
	defer f.Close()

	label := req.Label
	up, err := l.svcCtx.Uploader.UploadFile(path.Join(cast.ToString(l.ctx.Value("uid")), label), h)
	if err != nil {
		return nil, err
	}
	in := &syslogrpc.UploadLogReq{
		UserId:   cast.ToInt64(l.ctx.Value("uid")),
		Label:    label,
		FileName: h.Filename,
		FileSize: h.Size,
		FileMd5:  crypto.Md5v(h.Filename, ""),
		FileUrl:  up,
	}

	out, err := l.svcCtx.SyslogRpc.AddUploadLog(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUploadTypes(out), nil
}
