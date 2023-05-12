package service

import (
	"go.uber.org/zap"
	"myblog/cmd/gorm/model"
	"myblog/cmd/gorm/query"
	"myblog/model/dto"
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
