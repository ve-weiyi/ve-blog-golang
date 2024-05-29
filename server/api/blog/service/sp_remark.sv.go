package service

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
)

type RemarkService struct {
	svcCtx *svc.ServiceContext
}

func NewRemarkService(svcCtx *svc.ServiceContext) *RemarkService {
	return &RemarkService{
		svcCtx: svcCtx,
	}
}

// 创建Remark记录
func (l *RemarkService) CreateRemark(reqCtx *request.Context, remark *entity.Remark) (data *entity.Remark, err error) {
	remark.IpAddress = reqCtx.IpAddress
	remark.IpSource = reqCtx.GetIpSource()

	user, err := l.svcCtx.UserInformationRepository.First(reqCtx, "user_id = ?", reqCtx.Uid)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{"791422171@qq.com"},
		Subject: "【blog】新增留言信息",
		Content: fmt.Sprintf("%v remark: %v", user.Nickname, remark.MessageContent),
		Type:    0,
	}

	err = l.svcCtx.EmailPublisher.PublishMessage([]byte(jsonconv.ObjectToJson(msg)))
	if err != nil {
		glog.Info("PublishMessage:", err)
	}

	return l.svcCtx.RemarkRepository.Create(reqCtx, remark)
}

// 更新Remark记录
func (l *RemarkService) UpdateRemark(reqCtx *request.Context, remark *entity.Remark) (data *entity.Remark, err error) {
	return l.svcCtx.RemarkRepository.Update(reqCtx, remark)
}

// 删除Remark记录
func (l *RemarkService) DeleteRemark(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.RemarkRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Remark记录
func (l *RemarkService) FindRemark(reqCtx *request.Context, req *request.IdReq) (data *entity.Remark, err error) {
	return l.svcCtx.RemarkRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Remark记录
func (l *RemarkService) DeleteRemarkList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.RemarkRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Remark记录
func (l *RemarkService) FindRemarkList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Remark, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.RemarkRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.RemarkRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
