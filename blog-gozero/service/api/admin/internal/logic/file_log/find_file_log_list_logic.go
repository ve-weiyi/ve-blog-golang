package file_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/common/apiutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文件日志
func NewFindFileLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileLogListLogic {
	return &FindFileLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindFileLogListLogic) FindFileLogList(req *types.QueryFileLogReq) (resp *types.PageResp, err error) {
	in := &syslogrpc.FindFileLogListReq{
		Paginate: &syslogrpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		FilePath: req.FilePath,
		FileName: req.FileName,
		FileType: req.FileType,
	}

	out, err := l.svcCtx.SyslogRpc.FindFileLogList(l.ctx, in)
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

	var list []*types.FileLogBackVO
	for _, v := range out.List {
		m := ConvertFileLogTypes(v, usm)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = out.Pagination.Page
	resp.PageSize = out.Pagination.PageSize
	resp.Total = out.Pagination.Total
	resp.List = list
	return resp, nil
}

func ConvertFileLogTypes(in *syslogrpc.FileLogDetailsResp, usm map[string]*types.UserInfoVO) (out *types.FileLogBackVO) {
	out = &types.FileLogBackVO{
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
