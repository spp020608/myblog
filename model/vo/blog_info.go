package vo

import "myblog/pkg/ctime"

//目前bug 分层定义结构体的话 user和category绑定不进去输出
/*type Blog struct {
    BlogID       uint           `gorm:"column:blog_id" json:"blogId"`
    Title        string         `gorm:"column:title" json:"title"`
    Content      string         `gorm:"column:content" json:"content"`
    PostTime     ctime.CTime    `gorm:"column:post_time" json:"postTime"`
    ReadCount    int            `gorm:"column:read_count" json:"readCount"`
    User         User           `gorm:"foreignkey:UserID" json:"user"`
    Category     Category       `gorm:"foreignkey:CategoryID" json:"category"`
}

type User struct {
    UserID    uint   `gorm:"column:user_id" json:"userId"`
    Username  string `gorm:"column:username" json:"username"`
    Tel       string `gorm:"column:tel" json:"tel"`
}

type Category struct {
    CategoryID   uint   `gorm:"column:category_id" json:"categoryId"`
    CategoryName string `gorm:"column:category_name" json:"categoryName"`
}*/

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
