package service

import (
	"go.uber.org/zap"
	"myblog/cmd/gorm/model"
	"myblog/cmd/gorm/query"
	"myblog/model/dto"
	"myblog/model/vo"
	"time"
)

var CommentService = &commentService{}

type commentService struct{}

// Register 用户注册
//
//	@parameter	registerParam
//	@return		rowsAffected 此次操作影响的行数, 注册成功, 该返回值为 1
//	@return		err

// AddComment
//
//	@Description:添加评论
//	@receiver	commentService
//	@parameter	Comment 添加的评论的结构体
//	@return		rowsAffected 成功返回1 错误返回0
//	@return		err
func (commentService) AddComment(Comment *dto.Comment) (rowsAffected int64, err error) {

	zap.S().Infof("要增加评论为: %v\n", Comment)

	comment := &model.Comment{
		CommentID:  Comment.CommentID,
		Content:    Comment.Content,
		CreateTime: time.Now(),
		UserID:     Comment.UserID,
		BlogID:     Comment.BlogID,
	}

	err = query.Comment.Create(comment)
	if err != nil {
		zap.S().Error(err)
		return 0, err
	}
	return 1, nil

}

// DelComment
//
//	@Description:删除评论
//	@receiver	commentService
//	@parameter	CommentID 删除的评论的结构体的ID
//	@return		rowsAffected 成功返回1 错误返回0
//	@return		err
func (commentService) DelComment(CommentID int32) (rowsAffected int64, err error) {

	zap.S().Infof("要删除评论ID为: %v\n", CommentID)

	info, err := query.Comment.Where(query.Comment.CommentID.Eq(CommentID)).Delete()
	if err != nil {
		zap.S().Error(info.Error)
		return 0, err
	}
	if info.RowsAffected == 0 {
		zap.S().Info("未找到要删除的评论")
		return 0, info.Error
	} else {
		return info.RowsAffected, nil
	}

}

// UpdateComment
//
//	@Description:修改评论
//	@receiver	commentService
//	@parameter	*dto.Comment 修改的评论的结构体
//	@return		rowsAffected 成功返回1 错误返回0
//	@return		err
func (commentService) UpdateComment(Comment *dto.Comment) (rowsAffected int64, err error) {

	zap.S().Infof("要修改评论ID为: %v\n", Comment.CommentID)
	comment := &model.Comment{
		CommentID:  Comment.CommentID,
		Content:    Comment.Content,
		CreateTime: time.Now(),
		UserID:     Comment.UserID,
		BlogID:     Comment.BlogID,
	}
	update, err := query.Comment.Where(query.Comment.CommentID.Eq(comment.CommentID)).Update(query.Comment.Content, comment.Content)
	if err != nil {
		zap.S().Error(update.Error)
		return 0, err
	}
	if update.RowsAffected == 0 {
		zap.S().Info("未找到要修改的评论")
		return 0, update.Error
	} else {
		return update.RowsAffected, nil
	}

}

// QueryCommentByBlogId
//
//	@Description:通过BlogId查询下面的所有评论
//	@receiver	commentService
//	@parameter	BlogId 													查询博客ID下的所有评论
//	@return		*[]vo.CommentSearch   	查询到的评论
//	@return		err
func (commentService) QueryCommentByBlogId(BlogId int32) (Comments *[]vo.CommentSearch, err error) {
	zap.S().Infof("开始查询BlogId为%v的博客", BlogId)
	err = query.Comment.Select(query.Comment.CommentID, query.Comment.Content, query.Comment.CreateTime, query.Comment.UserID, query.User.Username, query.User.Tel, query.Blog.BlogID, query.Blog.Title).
		LeftJoin(query.User, query.User.UserID.EqCol(query.Comment.UserID)).
		LeftJoin(query.Blog, query.Blog.BlogID.EqCol(query.Comment.BlogID)).
		Where(query.Comment.BlogID.Eq(BlogId)).
		Order(query.Comment.CreateTime).Scan(&Comments)

	if err != nil {
		zap.S().Error("查询评论失败", err)
		return nil, err
	} else {
		zap.S().Info("查询评论成功:", Comments)
		return Comments, nil
	}

}

// CountCommentByBlogId
//
//	@Description:	开始统计BlogId下的评论数量
//	@receiver		commentService
//	@parameter		BlogId 		要统计的BlogID
//	@return			Count			统计出的数量
//	@return			err
func (commentService) CountCommentByBlogId(BlogId int32) (Count int64, err error) {
	zap.S().Infof("开始统计BlogId为%v下的评论数量", BlogId)
	count, err := query.Comment.Where(query.Comment.BlogID.Eq(BlogId)).Count()
	if err != nil {
		zap.S().Error("统计评论数量失败", err)
		return 0, err
	} else {
		zap.S().Info("统计评论数量成功:", count)
		return count, nil
	}
}

// QueryCommentByUserId
//
//	@Description:通过UserId查询下面的所有评论
//	@receiver	commentService
//	@parameter	BlogId 									查询用户ID下的所有评论
//	@return		*[]vo.CommentSearch   	查询到的评论
//	@return		err
func (commentService) QueryCommentByUserId(UserId int32) (Comments *[]vo.CommentSearch, err error) {

	zap.S().Infof("开始查询UserId为%v的博客", UserId)
	err = query.Comment.Select(query.Comment.CommentID, query.Comment.Content, query.Comment.CreateTime, query.Comment.UserID, query.User.Username, query.User.Tel, query.Blog.BlogID, query.Blog.Title).
		LeftJoin(query.User, query.User.UserID.EqCol(query.Comment.UserID)).
		LeftJoin(query.Blog, query.Blog.BlogID.EqCol(query.Comment.BlogID)).
		Where(query.Comment.UserID.Eq(UserId)).
		Order(query.Comment.CreateTime).Scan(&Comments)

	if err != nil {
		zap.S().Error("查询评论失败", err)
		return nil, err
	} else {
		zap.S().Info("查询评论成功:", Comments)
		return Comments, nil
	}
}
