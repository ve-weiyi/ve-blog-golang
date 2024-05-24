package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type TalkService struct {
	svcCtx *svc.ServiceContext
}

func NewTalkService(svcCtx *svc.ServiceContext) *TalkService {
	return &TalkService{
		svcCtx: svcCtx,
	}
}

// 创建Talk记录
func (l *TalkService) CreateTalk(reqCtx *request.Context, talk *entity.Talk) (data *entity.Talk, err error) {
	talk.UserID = reqCtx.UID
	return l.svcCtx.TalkRepository.Create(reqCtx, talk)
}

// 更新Talk记录
func (l *TalkService) UpdateTalk(reqCtx *request.Context, talk *entity.Talk) (data *entity.Talk, err error) {
	if talk.UserID != reqCtx.UID {
		return nil, apierr.ErrorUserNotPermission
	}
	return l.svcCtx.TalkRepository.Update(reqCtx, talk)
}

// 删除Talk记录
func (l *TalkService) DeleteTalk(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.TalkRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Talk记录
func (l *TalkService) FindTalk(reqCtx *request.Context, req *request.IdReq) (data *entity.Talk, err error) {
	return l.svcCtx.TalkRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Talk记录
func (l *TalkService) DeleteTalkList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.TalkRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Talk记录
func (l *TalkService) FindTalkList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Talk, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.TalkRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.TalkRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// 获取说说详情列表
func (l *TalkService) FindTalkDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.TalkDetailsDTO, total int64, err error) {
	talkList, total, err := l.FindTalkList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询分类下的文章数量

	for _, talk := range talkList {

		user, err := l.svcCtx.UserInformationRepository.First(reqCtx, "id = ?", talk.UserID)
		if err != nil {
			continue
		}

		var imgList []string
		jsonconv.JsonToObject(talk.Images, &imgList)
		data := &response.TalkDetailsDTO{
			ID:           talk.ID,
			UserID:       talk.UserID,
			Nickname:     user.Nickname,
			Avatar:       user.Avatar,
			Content:      talk.Content,
			ImgList:      imgList,
			IsTop:        talk.IsTop,
			Status:       talk.Status,
			LikeCount:    10,
			CommentCount: 10,
			CreatedAt:    talk.CreatedAt,
			UpdatedAt:    talk.UpdatedAt,
		}

		list = append(list, data)
	}

	return list, total, err
}

// 获取说说详情
func (l *TalkService) FindTalkDetailsDTO(reqCtx *request.Context, req *request.IdReq) (data *response.TalkDetailsDTO, err error) {
	// 查询api信息
	talk, err := l.svcCtx.TalkRepository.First(reqCtx, "id = ?", req.Id)
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.UserInformationRepository.First(reqCtx, "id = ?", talk.UserID)
	if err != nil {
		return nil, err
	}

	var imgList []string
	jsonconv.JsonToObject(talk.Images, &imgList)
	data = &response.TalkDetailsDTO{
		ID:        talk.ID,
		UserID:    talk.UserID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Content:   talk.Content,
		ImgList:   imgList,
		IsTop:     talk.IsTop,
		Status:    talk.Status,
		LikeCount: 10,
		CreatedAt: talk.CreatedAt,
		UpdatedAt: talk.UpdatedAt,
	}
	return data, nil
}

// 点赞说说
func (l *TalkService) LikeTalk(reqCtx *request.Context, req *request.IdReq) (data interface{}, err error) {
	return l.svcCtx.TalkRepository.LikeTalk(reqCtx, reqCtx.UID, req.Id)
}
