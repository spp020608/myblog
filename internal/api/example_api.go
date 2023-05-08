package api

import (
	"github.com/gin-gonic/gin"
	"myblog/cmd/gorm/model"
	"myblog/internal/pkg/resp"
	"net/http"

)

var ExampleApi exampleApi

type exampleApi struct{}

//
// SimpleGet 最简单的 GET 请求,前端不传递任何参数
//	@Summary		最简单独额接口测试,使用 Get 请求, 不携带参数
//	@Description	最简单独额接口测试,使用 Get 请求, 不携带参数
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.HttpOk	"OK"
//	@Failure		400	{object}	swagger.Http400	"Bad Request"
//	@Failure		404	{object}	swagger.Http404	"Page Not Found"
//	@Failure		500	{object}	swagger.Http500	"InternalError"
//	@Router			/example/simpleGet [get]
func (exampleApi) SimpleGet(c *gin.Context) {
	resp.Ok(c, "最简单的 Get 请求", nil)
}

//
// GetWithUrlQueryString 最简单的 GET 请求, 前端请求参数在 url 上
//	@Summary		使用 Get 请求, 在 URL 上携带参数
//	@Description	使用 Get 请求, 在 URL 上携带参数
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			id		query		int				true	"用户 ID"
//	@Param			name	query		string			true	"用户名称"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/example/getWithUrlQueryString [get]
func (exampleApi) GetWithUrlQueryString(c *gin.Context) {
	var id = c.Query("id")
	var name = c.Query("name")
	resp.Ok(c, "", gin.H{
		"id":   id,
		"name": name,
	})
}

//
// GetWithPathVariable Get 请求, 请求参数是 URL 的一部分
//	@Summary		Get 请求, 请求参数是 URL 的一部分
//	@Description	Get 请求, 请求参数是 URL 的一部分
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			group_id	path		string			true	"组 ID"
//	@Param			account_id	path		string			true	"账号 ID"
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/example/getWithPathVar/groups/{group_id}/accounts/{account_id} [get]
func (exampleApi) GetWithPathVariable(c *gin.Context) {
	// /getWithPathVar/groups/:group_id/accounts/:account_id
	groupId := c.Param("group_id")
	accountId := c.Param("account_id")
	resp.Ok(c, "", gin.H{
		"group_id":   groupId,
		"account_id": accountId,
	})

}

//
// GetAuthorizationHeader Get 请求, 请求参数是 URL 的一部分
//	@Summary		Get 请求, 请求参数是 URL 的一部分
//	@Description	Get 请求, 请求参数是 URL 的一部分
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string			true	"Authorization Header"
//	@Success		200				{object}	swagger.HttpOk	"OK"
//	@Failure		400				{object}	swagger.Http400	"Bad Request"
//	@Failure		404				{object}	swagger.Http404	"Page Not Found"
//	@Failure		500				{object}	swagger.Http500	"InternalError"
//	@Router			/example/getAuthorizationHeader [get]
func (exampleApi) GetAuthorizationHeader(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"header": c.GetHeader("Authorization"),
	})
}

//
// GetExtendAttribute Get 请求, 请求参数在 URL 中携带, 对参数配置了校验规则
//	@Summary		Get 请求, 请求参数在 URL 中携带, 对参数配置了校验规则
//	@Description	Get 请求, 请求参数在 URL 中携带, 对参数配置了校验规则
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			enumString	query		string			false	"字符串类型,参数必须是下面下拉列表中的值"		Enums(A, B, C)
//	@Param			enumInt		query		int				false	"整数类型, 参数必须是下面下拉列表中的值"		Enums(1, 2, 3)
//	@Param			enumNumber	query		number			false	"浮点数类型, 参数必须是下面下拉列表中的值"		Enums(1.1, 1.2, 1.3)
//	@Param			string		query		string			false	"验证规则:字符串最小长度 5, 最大长度 10"	minlength(5)	maxlength(10)
//	@Param			int			query		int				false	"验证部规则: 整数最小值为 1, 最大值为 10"	minimum(1)		maximum(10)
//	@Param			default		query		string			false	"提供默认的字符串在编辑框中"				default(A)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/example/getExtendAttribute [get]
func (exampleApi) GetExtendAttribute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"enumString": c.Query("enumString"),
		"enumInt":    c.Query("enumInt"),
		"enumNumber": c.Query("enumNumber"),
		"string":     c.Query("string"),
		"int":        c.Query("int"),
		"default":    c.Query("default"),
	})
}

////============== POST ===============

//
// SimplePost 最简单的 POST 请求, 前端不传递任何参数
//	@Summary		最简单的 POST 请求, 前端不传递任何参数
//	@Description	最简单的 POST 请求, 前端不传递任何参数
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.HttpOk	"OK"
//	@Failure		400	{object}	swagger.Http400	"Bad Request"
//	@Failure		404	{object}	swagger.Http404	"Page Not Found"
//	@Failure		500	{object}	swagger.Http500	"InternalError"
//	@Router			/example/simplePost [post]
func (exampleApi) SimplePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "最简单的 Post 请求, 前端不携带任何参数",
	})
}

//
// PostWithUrlQuery 使用 POST 请求, 参数拼接到 URL 后面
//	@Summary		使用 POST 请求, 参数拼接到 URL 后面
//	@Description	使用 POST 请求, 参数拼接到 URL 后面
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			id		query		int				true	"用户 ID"
//	@Param			name	query		string			true	"用户名称"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/example/postWithUrlQuery [post]
func (exampleApi) PostWithUrlQuery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   c.Query("id"),
		"name": c.Query("name"),
	})
}

//
// PostWithPathVariable Post 请求, 请求参数是 URL 的一部分
//	@Summary		Post 请求, 请求参数是 URL 的一部分
//	@Description	Post 请求, 请求参数是 URL 的一部分
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			group_id	path		string			true	"组 ID"
//	@Param			account_id	path		string			true	"账号 ID"
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/example/postWithPathVariable/groups/{group_id}/accounts/{account_id} [post]
func (exampleApi) PostWithPathVariable(c *gin.Context) {
	// /getWithPathVar/groups/:group_id/accounts/:account_id
	groupId := c.Param("group_id")
	accountId := c.Param("account_id")

	c.JSON(http.StatusOK, gin.H{
		"group_id":   groupId,
		"account_id": accountId,
	})
}

//
// PostAuthorizationHeader Post 请求, 参数在 Header 中
//	@Summary		Post 请求, 参数在 Header 中
//	@Description	Post 请求, 参数在 Header 中
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string			true	"Authorization Header"
//	@Success		200				{object}	swagger.HttpOk	"OK"
//	@Failure		400				{object}	swagger.Http400	"Bad Request"
//	@Failure		404				{object}	swagger.Http404	"Page Not Found"
//	@Failure		500				{object}	swagger.Http500	"InternalError"
//	@Router			/example/postAuthorizationHeader [post]
func (exampleApi) PostAuthorizationHeader(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"header": c.GetHeader("Authorization"),
	})
}

//
// PostExtendAttribute Post 请求, 请求参数在 URL 中携带, 对参数配置了校验规则
//	@Summary		Post 请求, 请求参数在 URL 中携带, 对参数配置了校验规则
//	@Description	Post 请求, 请求参数在 URL 中携带, 对参数配置了校验规则
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			enumString	query		string			false	"字符串类型,参数必须是下面下拉列表中的值"		Enums(A, B, C)
//	@Param			enumInt		query		int				false	"整数类型, 参数必须是下面下拉列表中的值"		Enums(1, 2, 3)
//	@Param			enumNumber	query		number			false	"浮点数类型, 参数必须是下面下拉列表中的值"		Enums(1.1, 1.2, 1.3)
//	@Param			string		query		string			false	"验证规则:字符串最小长度 5, 最大长度 10"	minlength(5)	maxlength(10)
//	@Param			int			query		int				false	"验证部规则: 整数最小值为 1, 最大值为 10"	minimum(1)		maximum(10)
//	@Param			default		query		string			false	"提供默认的字符串在编辑框中"				default(A)
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/example/postExtendAttribute [post]
func (exampleApi) PostExtendAttribute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"enumString": c.Query("enumString"),
		"enumInt":    c.Query("enumInt"),
		"enumNumber": c.Query("enumNumber"),
		"string":     c.Query("string"),
		"int":        c.Query("int"),
		"default":    c.Query("default"),
	})
}

//
// PostJsonData Post 请求, 发送 Json 数据, 参数在消息体中
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中
//	@Tags			exampleApi
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.User		true	"前端发送的 Json 对象"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/example/postJsonData [post]
func (exampleApi) PostJsonData(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "绑定 User 对象失败" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, u)
}

//
// PostFormData Post 请求, 发送 Json 数据, 参数在消息体中
//	@Summary		Post 请求, 发送 Json 数据, 参数在消息体中
//	@Description	Post 请求, 发送 Json 数据, 参数在消息体中
//	@Tags			exampleApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			name	formData	string			false	"前端发送的表单中的 name 参数"
//	@Param			age		formData	int				false	"前端发送的表单中的 age 参数"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/example/postFormData [post]
func (exampleApi) PostFormData(c *gin.Context) {
	// multipart/form-data
	name := c.PostForm("name")
	age := c.PostForm("age")

	c.JSON(http.StatusOK, gin.H{
		"msg": "获取表单数据成功",
		"data": gin.H{
			"name": name,
			"age":  age,
		},
	})
}

//
// PostUrlEncodedFormData Post 请求, 发送表单数据, 表单类型为: x-www-form-urlencoded, 参数在消息体中
//	@Summary		Post 请求, 发送表单数据, 表单类型为: x-www-form-urlencoded, 参数在消息体中
//	@Description	Post 请求, 发送表单数据, 表单类型为: x-www-form-urlencoded, 参数在消息体中
//	@Tags			exampleApi
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			name	formData	string			false	"前端发送的表单中的 name 参数"
//	@Param			age		formData	int				false	"前端发送的表单中的 age 参数"
//	@Success		200		{array}		model.User		"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/example/postUrlEncodedFormData [post]
func (exampleApi) PostUrlEncodedFormData(c *gin.Context) {
	// x-www-form-urlencoded
	name := c.PostForm("name")
	age := c.PostForm("age")

	c.JSON(http.StatusOK, gin.H{
		"msg": "获取表单数据成功",
		"data": gin.H{
			"name": name,
			"age":  age,
		},
	})
}
