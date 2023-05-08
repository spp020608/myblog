package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myblog/internal/pkg/resp"
	"myblog/internal/service"
	"myblog/model/dto"
	"myblog/pkg/validate"
)

var UserApi = new(userApi)

type userApi struct{}

// Register 用户注册
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			false	"要注册的用户名"	minlength(1)	maxlength(20)
//	@Param			password	formData	string			false	"用户密码"		minlength(1)	maxlength(128)
//	@Param			tel			formData	string			false	"电话号码"		minlength(11)	maxlength(11)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/register [post]
func (userApi) Register(c *gin.Context) {
	var registerParam dto.RegisterParam
	// 将请求数据绑定到结构体
	if err := c.ShouldBind(&registerParam); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}

	insert, err := service.UserService.Register(&registerParam)
	if err != nil {
		zap.S().Errorf("注册用户失败: %s\n", err)
		resp.InternalServerError(c, "注册用户失败")
		return
	}
	resp.Ok(c, "注册用户成功", insert)
}

// Login 用户登录
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			true	"要注册的用户名"	minlength(1)	maxlength(20)
//	@Param			password	formData	string			true	"用户密码"		minlength(1)	maxlength(128)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/login [post]
func (userApi) Login(c *gin.Context) {
	var loginInfo dto.LoginInfo
	if err := c.ShouldBind(&loginInfo); err != nil {
		// 翻译错误消息
		translatedErr := validate.TranslateError(err)
		// 向前端回写错误数据
		resp.InternalServerError(c, translatedErr)
		// 将错误记录到日志文件
		zap.S().Errorf("用户登录失败: %s\n", translatedErr)
		return
	}

	fmt.Printf("%#v\n", loginInfo)

	// 调用 service 层的登录函数
	token, err := service.UserService.Login(&loginInfo)
	if err != nil {
		// 向前端回写错误数据
		resp.InternalServerError(c, err.Error())
		// 将错误记录到日志文件
		zap.S().Errorf("用户登录失败: %s\n", err.Error())
		return
	}
	resp.Ok(c, "登录成功", token)
}

// UpdatePassword 用户修改密码
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			true	"用户名"	minlength(1)	maxlength(20)
//	@Param			password	formData	string			true	"用户旧密码"	minlength(1)	maxlength(128)
//	@Param			newPassword	formData	string			true	"用户新密码"	minlength(1)	maxlength(128)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/updatePassword [post]
func (userApi) UpdatePassword(c *gin.Context) {
	var updatePassword dto.UpdatePassword
	if err := c.ShouldBind(&updatePassword); err != nil {
		// 翻译错误消息
		translatedErr := validate.TranslateError(err)
		// 向前端回写错误数据
		resp.InternalServerError(c, translatedErr)
		// 将错误记录到日志文件
		zap.S().Errorf("用户修改密码失败: %s\n", translatedErr)
		return
	}

	fmt.Printf("%#v\n", updatePassword)

	login, _ := service.UserService.UpdatePassword(&updatePassword)

	resp.Ok(c, "修改密码成功", login)
}

// UpdateUser  	修改用户信息
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			false	"用户名"	minlength(1)	maxlength(20)
//	@Param			password	formData	string			false	"用户密码"	minlength(1)	maxlength(128)
//	@Param			tel			formData	string			false	"电话号码"	minlength(11)	maxlength(11)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/updateUser [post]
func (userApi) UpdateUser(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBind(&user); err != nil {
		// 翻译错误消息
		translatedErr := validate.TranslateError(err)
		// 向前端回写错误数据
		resp.InternalServerError(c, translatedErr)
		// 将错误记录到日志文件
		zap.S().Errorf("用户修改信息失败: %s\n", translatedErr)
		return
	}
	zap.S().Info("传进来的信息为", user)
	updateUser, err := service.UserService.UpdateUser(&user)
	if err != nil {
		zap.S().Errorf("用户修改信息失败: %s\n", err)
	}
	resp.Ok(c, "更新用户成功", updateUser)
}

// DelUser 用户注销
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string			false	"要注销的用户名"	minlength(1)	maxlength(20)
//	@Param			password	formData	string			false	"密码"		minlength(1)	maxlength(128)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/user/delUser [post]
func (userApi) DelUser(c *gin.Context) {
	var user dto.LoginInfo
	if err := c.ShouldBind(&user); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}
	delUser, err := service.UserService.DelUser(&user)
	if err != nil {
		zap.S().Errorf("注销用户失败: %s\n", err)
		resp.InternalServerError(c, "注销用户失败")
		return
	}
	resp.Ok(c, "注销用户成功", delUser)
}
