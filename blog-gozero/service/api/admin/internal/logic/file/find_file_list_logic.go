package file

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/common/apiutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/resourcerpc"
)

type FindFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文件列表
func NewFindFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileListLogic {
	return &FindFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFileListLogic) FindFileList(req *types.FileQuery) (resp *types.PageResp, err error) {
	in := &resourcerpc.FindFileUploadListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
		FilePath: req.FilePath,
		FileType: req.FileType,
	}

	out, err := l.svcCtx.ResourceRpc.FindFileUploadList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var uids []string
	for _, v := range out.List {
		uids = append(uids, v.UserId)
	}

	// 获取用户信息
	usm, err := apiutils.GetUserInfos(l.ctx, l.svcCtx, uids)
	if err != nil {
		return nil, err
	}

	var list []*types.FileBackVO
	for _, v := range out.List {
		m := ConvertFileUploadTypes(v, usm)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertFileUploadTypes(in *resourcerpc.FileUploadDetails, usm map[string]*types.UserInfoVO) (out *types.FileBackVO) {
	out = &types.FileBackVO{
		Id:        in.Id,
		UserId:    in.UserId,
		FilePath:  in.FilePath,
		FileName:  in.FileName,
		FileType:  in.FileType,
		FileSize:  in.FileSize,
		FileMd5:   in.FileMd5,
		FileUrl:   in.FileUrl,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}

	// 用户信息
	if in.UserId != "" {
		user, ok := usm[in.UserId]
		if ok && user != nil {
			out.Creator = user
		}
	}

	return
}
