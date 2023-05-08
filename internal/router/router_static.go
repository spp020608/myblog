package router

import "github.com/gin-gonic/gin"

//
// setupStaticRouter 设置静态共享目录
//	@parameter	r
//
func setupStaticRouter(r *gin.Engine) {
	// 在网络上共享上传的文件
	r.Static("/uploads", "./uploads")
}
