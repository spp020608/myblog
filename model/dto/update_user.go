package dto

type User struct {
	Username string `json:"username" form:"username" biding:"required,min=1,max=32"`  // 用户名
	Password string `json:"password" form:"password" biding:"required,min=1,max=128"` // 密码
	Tel      string `json:"tel" form:"tel" binding:"alphanum,len=11"`
}
