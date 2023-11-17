package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

// 获取说说详情列表
func (s *TalkService) FindTalkDetailsList(reqCtx *request.Context, page *request.PageQuery) (list []*response.TalkDetails, total int64, err error) {
	talkList, total, err := s.FindTalkList(reqCtx, page)
	if err != nil {
		return nil, 0, err
	}

	// 查询分类下的文章数量

	for _, talk := range talkList {

		user, err := s.svcCtx.UserInformationRepository.FindUserInformationById(reqCtx, talk.UserID)
		if err != nil {
			return nil, 0, err
		}

		var imgList []string
		jsonconv.JsonToObject(talk.Images, &imgList)
		data := &response.TalkDetails{
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
func (s *TalkService) FindTalkDetails(reqCtx *request.Context, id int) (data *response.TalkDetails, err error) {
	// 查询api信息
	talk, err := s.svcCtx.TalkRepository.FindTalkById(reqCtx, id)
	if err != nil {
		return nil, err
	}

	user, err := s.svcCtx.UserInformationRepository.FindUserInformationById(reqCtx, talk.UserID)
	if err != nil {
		return nil, err
	}

	var imgList []string
	jsonconv.JsonToObject(talk.Images, &imgList)
	data = &response.TalkDetails{
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
