package logic

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/repository/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
)

type CommentRepository struct {
	DbEngin *gorm.DB
	Cache   *redis.Client
}

func NewCommentRepository(svcCtx *svc.RepositoryContext) *CommentRepository {
	return &CommentRepository{
		DbEngin: svcCtx.DbEngin,
		Cache:   svcCtx.Cache,
	}
}

// 创建Comment记录
func (s *CommentRepository) CreateComment(comment *entity.Comment) (out *entity.Comment, err error) {
	db := s.DbEngin
	err = db.Create(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, err
}

// 删除Comment记录
func (s *CommentRepository) DeleteComment(comment *entity.Comment) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&comment)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 更新Comment记录
func (s *CommentRepository) UpdateComment(comment *entity.Comment) (out *entity.Comment, err error) {
	db := s.DbEngin
	err = db.Save(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, err
}

// 查询Comment记录
func (s *CommentRepository) GetComment(id int) (out *entity.Comment, err error) {
	db := s.DbEngin
	err = db.Where("id = ?", id).First(&out).Error
	if err != nil {
		return nil, err
	}
	return out, err
}

// 批量删除Comment记录
func (s *CommentRepository) DeleteCommentByIds(ids []int) (rows int64, err error) {
	db := s.DbEngin
	query := db.Delete(&[]entity.Comment{}, "id in ?", ids)
	err = query.Error
	rows = query.RowsAffected
	return rows, err
}

// 分页查询Comment记录
func (s *CommentRepository) FindCommentList(page *request.PageInfo) (list []*entity.Comment, total int64, err error) {
	// 创建db
	db := s.DbEngin

	// 如果有搜索条件
	if len(page.Conditions) != 0 {
		query, args := page.WhereClause()
		db = db.Where(query, args...)
	}

	// 如果有排序参数
	if len(page.Orders) != 0 {
		db = db.Order(page.OrderClause())
	}

	// 查询总数,要在使用limit之前
	err = db.Model(&list).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 如果有分页参数
	if page.Page != 0 || page.PageSize != 0 {
		limit := page.Limit()
		offset := page.Offset()
		db = db.Limit(limit).Offset(offset)
	}

	// 查询数据
	err = db.Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// 查看评论id集合下的回复评论
func (s *CommentRepository) FindCommentReplyList(id int, page *request.PageInfo) (out []*entity.Comment, total int64, err error) {

	page.Conditions = append(page.Conditions, &request.Condition{
		Flag:  "AND",
		Field: "parent_id",
		Rule:  "=",
		Value: id,
	})

	return s.FindCommentList(page)
}

// 点赞评论
func (s *CommentRepository) LikeComment(ctx context.Context, uid int, commentId int) (data interface{}, err error) {
	// 用户点赞的评论列表
	commentUserLikeKey := constant.RedisWrapKey("comment_user_like", uid)
	// 当前评论的点赞量
	commentLikeCountKey := constant.RedisWrapKey("comment_like_count", commentId)

	// 判断是否已经点赞
	if s.Cache.SIsMember(ctx, commentUserLikeKey, commentId).Val() {
		// 点过赞则删除评论id
		s.Cache.SRem(ctx, commentUserLikeKey, commentId)
		// 评论点赞量-1
		s.Cache.Decr(ctx, commentLikeCountKey)
	} else {
		// 未点赞则增加评论id
		s.Cache.SAdd(ctx, commentUserLikeKey, commentId)
		// 评论点赞量+1
		s.Cache.Incr(ctx, commentLikeCountKey)
	}

	return data, nil
}

// 获取用户点赞记录
func (s *CommentRepository) GetUserLikeComment(ctx context.Context, uid int) (data []string, err error) {
	// 用户点赞的评论列表
	commentUserLikeKey := constant.RedisWrapKey("comment_user_like", uid)
	return global.REDIS.SMembers(ctx, commentUserLikeKey).Result()
}
