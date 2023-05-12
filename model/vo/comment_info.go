package vo

import (
	"myblog/pkg/ctime"
)

type CommentSearch struct {
	CommentID  int32       `gorm:"column:comment_id;primaryKey;autoIncrement:true;comment:评论ID，自增ID" json:"comment_id"`
	Content    string      `gorm:"column:content;comment:评论内容" json:"content"`
	CreateTime ctime.CTime `gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:评论时间" json:"create_time"`
	UserID     int32       `gorm:"column:user_id;comment:评论人ID" json:"user_id"`
	Username   string      `gorm:"column:username"`
	Tel        string      `gorm:"column:tel"`
	BlogID     int32       `gorm:"column:blog_id;comment:被评论的博客ID" json:"blog_id"`
	Title      string      `gorm:"column:title"`
}
