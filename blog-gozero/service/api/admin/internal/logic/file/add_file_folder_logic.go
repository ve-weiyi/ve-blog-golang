package file

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/common/apiutils"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFileFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文件目录
func NewAddFileFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFileFolderLogic {
	return &AddFileFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFileFolderLogic) AddFileFolder(req *types.FileFolderNewReq) (resp *types.FileBackVO, err error) {
	in := &resourcerpc.FileUploadNewReq{
		Id:       0,
		UserId:   cast.ToString(l.ctx.Value(restx.HeaderUid)),
		FilePath: req.FilePath,
		FileName: req.FileName,
		FileType: "",
		FileSize: 0,
		FileMd5:  "",
		FileUrl:  "",
	}

	out, err := l.svcCtx.ResourceRpc.AddFileUpload(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	usm, err := apiutils.GetUserInfos(l.ctx, l.svcCtx, []string{out.UserId})
	if err != nil {
		return nil, err
	}

	resp = ConvertFileUploadTypes(out, usm)
	return resp, nil
}
