package dto

//
// LoginInfo 用户登录功能, 前端需要传递的表单数据
//
type LoginInfo struct {
	Username string `json:"username" form:"username" biding:"required,min=1,max=32"` // 用户名
	Password string `json:"password" form:"password" biding:"required,min=1,max=128"` // 密码
}
