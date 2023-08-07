package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
)

// 获取说说详情列表
func (s *TalkService) FindTalkListDetails(reqCtx *request.Context, page *request.PageQuery) (list []*response.TalkDetails, total int64, err error) {
	categories, total, err := s.svcCtx.TalkRepository.FindTalkList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询分类下的文章数量

	for _, talk := range categories {

		user, err := s.svcCtx.UserInformationRepository.FindUserInformation(reqCtx, talk.UserID)
		if err != nil {
			return nil, 0, err
		}

		data := &response.TalkDetails{
			ID:        talk.ID,
			UserID:    talk.UserID,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Content:   talk.Content,
			Images:    talk.Images,
			IsTop:     talk.IsTop,
			Status:    talk.Status,
			LikeCount: 10,
			CreatedAt: talk.CreatedAt,
			UpdatedAt: talk.UpdatedAt,
		}

		list = append(list, data)
	}

	return list, total, err
}

// 获取说说详情
func (s *TalkService) FindTalkDetails(reqCtx *request.Context, id int) (data *response.TalkDetails, err error) {
	// 查询api信息
	talk, err := s.svcCtx.TalkRepository.FindTalk(reqCtx, id)
	if err != nil {
		return nil, err
	}

	user, err := s.svcCtx.UserInformationRepository.FindUserInformation(reqCtx, talk.UserID)
	if err != nil {
		return nil, err
	}

	data = &response.TalkDetails{
		ID:        talk.ID,
		UserID:    talk.UserID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Content:   talk.Content,
		Images:    talk.Images,
		IsTop:     talk.IsTop,
		Status:    talk.Status,
		LikeCount: 10,
		CreatedAt: talk.CreatedAt,
		UpdatedAt: talk.UpdatedAt,
	}
	return data, nil
}
