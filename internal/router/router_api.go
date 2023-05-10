package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"myblog/internal/api"
)

func setupApiRouter(r *gin.Engine) {
	// http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	exampleGroup := r.Group("/example")
	{
		// 最简单的 GET 请求, 不携带任何参数
		exampleGroup.GET("/simpleGet", api.ExampleApi.SimpleGet)
		// GET 请求, 使用 URL 传递参数
		exampleGroup.GET("/getWithUrlQueryString", api.ExampleApi.GetWithUrlQueryString)
		// GET 请求, 使用 URL 路径传递参数 PathVariable 形式
		exampleGroup.GET("/getWithPathVar/groups/:group_id/accounts/:account_id", api.ExampleApi.GetWithPathVariable)
		// GET 请求, 用 Header 传递数据
		exampleGroup.GET("/getAuthorizationHeader", api.ExampleApi.GetAuthorizationHeader)
		// GET 请求, 使用 swagger 中的扩展属性
		exampleGroup.GET("/getExtendAttribute", api.ExampleApi.GetExtendAttribute)

		// 最简单的 POST 请求, 不携带任何参数
		exampleGroup.POST("/simplePost", api.ExampleApi.SimplePost)
		// POST 请求, 请求参数在 URL 中
		exampleGroup.POST("/postWithUrlQuery", api.ExampleApi.PostWithUrlQuery)
		// POST 请求, 使用 URL 路径传递参数 PathVariable 形式
		exampleGroup.POST("/postWithPathVariable/groups/:group_id/accounts/:account_id", api.ExampleApi.PostWithPathVariable)
		// POST 请求, 用 Header 传递数据
		exampleGroup.POST("/postAuthorizationHeader", api.ExampleApi.PostAuthorizationHeader)
		// POST 请求, 使用 swagger 中的扩展属性
		exampleGroup.POST("/postExtendAttribute", api.ExampleApi.PostExtendAttribute)
		// POST 请求, 发送 JSON 数据, 数据在消息体中
		exampleGroup.POST("/postJsonData", api.ExampleApi.PostJsonData)
		// POST 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
		exampleGroup.POST("/postFormData", api.ExampleApi.PostFormData)
		// POST 请求, 发送 x-www-form-urlencodered 类型的表单数据, 参数在消息体中, 编码方式为 urlencoder
		exampleGroup.POST("/postUrlEncodedFormData", api.ExampleApi.PostUrlEncodedFormData)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", api.UserApi.Register)
		userGroup.POST("/login", api.UserApi.Login)
		userGroup.POST("/updatePassword", api.UserApi.UpdatePassword)
		userGroup.POST("/updateUser", api.UserApi.UpdateUser)
		userGroup.POST("/delUser", api.UserApi.DelUser)
	}
	blogGroup := r.Group("/blog")
	{
		blogGroup.POST("/getAllBlog", api.BlogApi.GetAllBlog)
		blogGroup.POST("/blogWithTitle", api.BlogApi.BlogWithTitle)
		blogGroup.POST("/blogWithPostTime", api.BlogApi.BlogWithPostTime)
		blogGroup.POST("/blogWithReadCount", api.BlogApi.BlogWithReadCount)
	}
}
