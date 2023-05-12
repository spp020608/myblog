package dto

import "time"

type Comment struct {
	CommentID  int32     `gorm:"column:comment_id;primaryKey;autoIncrement:true;comment:评论ID，自增ID" json:"comment_id"`
	Content    string    `gorm:"column:content;comment:评论内容" form:"content"  json:"content"`
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:评论时间" json:"create_time"`
	UserID     int32     `gorm:"column:user_id;comment:评论人ID" form:"userID" json:"user_id"`
	BlogID     int32     `gorm:"column:blog_id;comment:被评论的博客ID"  form:"blogID" json:"blog_id"`
}
