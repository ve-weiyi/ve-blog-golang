package logic

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type CommentService struct {
	svcCtx *svc.ServiceContext
}

func NewCommentService(svcCtx *svc.ServiceContext) *CommentService {
	return &CommentService{
		svcCtx: svcCtx,
	}
}

// 创建Comment记录
func (l *CommentService) CreateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return l.svcCtx.CommentRepository.Create(reqCtx, comment)
}

// 更新Comment记录
func (l *CommentService) UpdateComment(reqCtx *request.Context, comment *entity.Comment) (data *entity.Comment, err error) {
	return l.svcCtx.CommentRepository.Update(reqCtx, comment)
}

// 删除Comment记录
func (l *CommentService) DeleteComment(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.CommentRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Comment记录
func (l *CommentService) FindComment(reqCtx *request.Context, req *request.IdReq) (data *entity.Comment, err error) {
	return l.svcCtx.CommentRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Comment记录
func (l *CommentService) DeleteCommentList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.CommentRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Comment记录
func (l *CommentService) FindCommentList(reqCtx *request.Context, page *request.CommentQueryReq) (list []*response.CommentDTO, total int64, err error) {
	p, s, order, cond, args := ConvertCommentQueryTypes(page)

	commentList, err := l.svcCtx.CommentRepository.FindList(reqCtx, int(p), int(s), order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.CommentRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	var userIds []int
	var commentIds []int
	for _, item := range commentList {
		userIds = append(userIds, item.UserId)
		commentIds = append(commentIds, item.Id)
	}

	// 查询用户
	users, _ := l.svcCtx.UserInformationRepository.FindALL(reqCtx, "id in (?)", userIds)

	var userMap = make(map[int]*entity.UserInformation)
	for _, item := range users {
		userMap[item.Id] = item
	}

	for _, item := range commentList {
		// 查询评论下所有回复列表,只显示五条
		replyList, count, _ := l.FindCommentReplyList(reqCtx, item.Id, &request.PageQuery{
			Limit: request.PageLimit{
				Page:     1,
				PageSize: 5,
			},
		})
		// 查询当前评论下所有回复列表
		data := &response.CommentDTO{
			Id:             item.Id,
			UserId:         item.UserId,
			CommentContent: item.CommentContent,
			LikeCount:      100,
			CreatedAt:      item.CreatedAt,
			ReplyCount:     count,
			ReplyDTOList:   replyList,
		}

		// 用户信息
		info, _ := userMap[item.UserId]
		if info != nil {
			data.Nickname = info.Nickname
			data.Avatar = info.Avatar
			data.Website = info.Website
		}

		// 回复的用户信息
		//rinfo, _ := userMap[item.ReplyUserId]
		//if rinfo != nil {
		//	data.ReplyUserId = rinfo.Id
		//	data.ReplyNickname = rinfo.Nickname
		//	data.ReplyWebsite = rinfo.Website
		//}

		list = append(list, data)
	}

	return
}

// 查询Comment记录
func (l *CommentService) FindCommentReplyList(reqCtx *request.Context, commentId int, page *request.PageQuery) (list []*response.ReplyDTO, total int64, err error) {
	page.Conditions = append(page.Conditions, &request.PageCondition{Field: "parent_id", Operator: "=", Value: commentId})

	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询评论下所有回复列表
	replyList, err := l.svcCtx.CommentRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.CommentRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// 收集需要查询的用户id
	var userIds []int
	for _, item := range replyList {
		userIds = append(userIds, item.UserId)
		userIds = append(userIds, item.ReplyUserId)
	}

	// 查询用户
	users, _ := l.svcCtx.UserInformationRepository.FindALL(reqCtx, "id in (?)", userIds)

	var userMap = make(map[int]*entity.UserInformation)
	for _, item := range users {
		userMap[item.Id] = item
	}

	// 组装返回数据
	for _, item := range replyList {

		data := &response.ReplyDTO{
			Id:             item.Id,
			ParentId:       item.ParentId,
			UserId:         item.UserId,
			ReplyUserId:    item.ReplyUserId,
			CommentContent: item.CommentContent,
			LikeCount:      5,
			CreatedAt:      item.CreatedAt,
		}

		// 用户信息
		info, _ := userMap[item.UserId]
		if info != nil {
			data.Nickname = info.Nickname
			data.Avatar = info.Avatar
			data.Website = info.Website
		}

		// 回复的用户信息
		rinfo, _ := userMap[item.ReplyUserId]
		if rinfo != nil {
			data.ReplyUserId = rinfo.Id
			data.ReplyNickname = rinfo.Nickname
			data.ReplyWebsite = rinfo.Website
		}

		list = append(list, data)
	}
	return list, total, nil
}

// 查询Comment后台记录
func (l *CommentService) FindCommentBackList(reqCtx *request.Context, page *request.PageQuery) (list []*response.CommentBackDTO, total int64, err error) {
	// 使用用户昵称查询
	var cd *request.PageCondition
	for _, condition := range page.Conditions {
		if condition.Field == "username" {
			cd = condition
		}
	}

	if cd != nil {
		accounts, err := l.svcCtx.UserAccountRepository.FindALL(reqCtx, "username like ?")
		if err != nil {
			return nil, 0, err
		}

		var userIds []int
		for _, item := range accounts {
			userIds = append(userIds, item.Id)
		}
		// 替换查询条件
		cd.Field = "user_id"
		cd.Value = userIds
		cd.Operator = "in"
	}

	cond, args := page.ConditionClause()
	order := page.OrderClause()
	// 查询评论下所有回复列表
	commentList, err := l.svcCtx.CommentRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	total, err = l.svcCtx.CommentRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// 收集需要查询的用户id
	var userIds []int
	var articleIds []int
	for _, item := range commentList {
		userIds = append(userIds, item.UserId)
		userIds = append(userIds, item.ReplyUserId)
		articleIds = append(articleIds, item.TopicId)
	}

	// 查询用户
	users, _ := l.svcCtx.UserInformationRepository.FindALL(reqCtx, "id in (?)", userIds)
	var userMap = make(map[int]*entity.UserInformation)
	for _, item := range users {
		userMap[item.Id] = item
	}
	// 查询文章
	articles, _ := l.svcCtx.ArticleRepository.FindALL(reqCtx, "id in (?)", articleIds)
	var articleMap = make(map[int]*entity.Article)
	for _, item := range articles {
		articleMap[item.Id] = item
	}

	// 组装返回数据
	for _, item := range commentList {

		data := &response.CommentBackDTO{
			Id:             item.Id,
			Avatar:         "",
			Nickname:       "",
			ReplyNickname:  "",
			TopicTitle:     "",
			CommentContent: item.CommentContent,
			Type:           item.Type,
			IsReview:       item.IsReview,
			CreatedAt:      item.CreatedAt,
		}

		// 用户信息
		info, _ := userMap[item.UserId]
		if info != nil {
			data.Avatar = info.Avatar
			data.Nickname = info.Nickname
		}

		// 回复的用户信息
		rinfo, _ := userMap[item.ReplyUserId]
		if rinfo != nil {
			data.ReplyNickname = rinfo.Nickname
		}

		// 回复的文章信息
		aInfo, _ := articleMap[item.TopicId]
		if aInfo != nil {
			data.TopicTitle = aInfo.ArticleTitle
		}

		list = append(list, data)
	}
	return list, total, nil
}

// 点赞Comment
func (l *CommentService) LikeComment(reqCtx *request.Context, commentId int) (data interface{}, err error) {
	return l.svcCtx.CommentRepository.LikeComment(reqCtx, reqCtx.Uid, commentId)
}

func ConvertCommentQueryTypes(in *request.CommentQueryReq) (page int64, pageSize int64, sorts string, conditions string, args []interface{}) {
	//var page, pageSize int64
	//var sorts, conditions string
	//var args []string

	page = in.Page
	pageSize = in.PageSize

	if in.OrderBy != "" {
		sorts = fmt.Sprintf("`%s` desc", in.OrderBy)
	}

	if in.TopicId >= 0 {
		conditions = "topic_id = ? "
		args = append(args, cast.ToString(in.TopicId))
	}

	if in.ParentId >= 0 {
		conditions = conditions + "and "
		conditions = conditions + "parent_id = ? "
		args = append(args, cast.ToString(in.ParentId))
	}

	if in.Type >= 0 {
		conditions = conditions + "and "
		conditions = conditions + "type = ? "
		args = append(args, cast.ToString(in.Type))
	}

	return page, pageSize, sorts, conditions, args
}
