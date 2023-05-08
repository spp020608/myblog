package dto

type UpdatePassword struct {
	Username string `json:"username" form:"username" biding:"required,min=1,max=32"` // 用户名
	Password string `json:"password" form:"password" biding:"required,min=1,max=128"` // 密码
	NewPassword string`json:"newPassword" form:"newPassword" biding:"required,min=1,max=128"` // 新密码
}