package photorpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoListLogic {
	return &FindPhotoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取照片列表
func (l *FindPhotoListLogic) FindPhotoList(in *photorpc.FindPhotoListReq) (*photorpc.FindPhotoListResp, error) {
	page, size, sorts, conditions, params := convertPhotoQuery(in)

	records, total, err := l.svcCtx.TPhotoModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*photorpc.PhotoDetails
	for _, v := range records {
		list = append(list, convertPhotoOut(v))
	}

	return &photorpc.FindPhotoListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertPhotoQuery(in *photorpc.FindPhotoListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	if in.AlbumId != 0 {
		conditions += " album_id = ?"
		params = append(params, in.AlbumId)
	}

	return
}
