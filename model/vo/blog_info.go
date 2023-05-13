package vo

import (
	"myblog/pkg/ctime"
)

type FirstBlog struct {
	BlogID    int32       `gorm:"column:blog_id;primaryKey;autoIncrement:true;comment:主键" json:"blog_id"`
	Title     string      `gorm:"column:title;not null;comment:标题" json:"title"`
	Content   string      `gorm:"column:content;comment:内容" json:"content"`
	PostTime  ctime.CTime `gorm:"column:post_time;default:CURRENT_TIMESTAMP;comment:发表时间" json:"post_time"`
	ReadCount int32       `gorm:"column:read_count;comment:阅读量" json:"read_count"`
	User      User        `gorm:"embedded"`
	Category  Category    `gorm:"embedded"`
}

type User struct {
	UserID   int32  `gorm:"column:user_id;primaryKey;autoIncrement:true;comment:主键" json:"user_id"`
	Username string `gorm:"column:username;not null;comment:用户名" json:"username"`
	Tel      string `gorm:"column:tel;comment:手机号" json:"tel"`
}

type Category struct {
	CategoryID   int32  `gorm:"column:category_id;primaryKey;autoIncrement:true;comment:分类ID，自增ID" json:"category_id"`
	CategoryName string `gorm:"column:category_name;not null;comment:分类名称" json:"category_name"`
}

type Blog struct {
	BlogID       uint        `gorm:"column:blog_id"`
	Title        string      `gorm:"column:title"`
	Content      string      `gorm:"column:content"`
	PostTime     ctime.CTime `gorm:"column:post_time"`
	ReadCount    int         `gorm:"column:read_count"`
	UserID       uint        `gorm:"column:user_id"`
	Username     string      `gorm:"column:username"`
	Tel          string      `gorm:"column:tel"`
	CategoryID   uint        `gorm:"column:category_id"`
	CategoryName string      `gorm:"column:category_name"`
}
