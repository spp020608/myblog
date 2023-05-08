package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"myblog/internal/repository/cache"
	"myblog/internal/repository/dao"
	midware "myblog/internal/router/middleware"

	"go.uber.org/zap"
)

//
// NewHttpServer 创建 Gin 服务器, 初始化 mysql 数据库连接和 redis 连接
//	@return	*gin.Engine
//
func NewHttpServer() *gin.Engine {

	r := gin.New()

	r.Use(
		cors.Default(),                // 跨域中间件
		midware.ZapLoggerMiddleware(), // 自定义 zapLog 中间件取代 gin 框架默认的日志中间件
		gin.Recovery(),                // gin 框架默认的错误恢复中间件)
	)

	// 解决 405 错误
	r.HandleMethodNotAllowed = true
	// 上传文件的时候, 需要指定的缓冲区大小
	r.MaxMultipartMemory = 8 << 20 // 8MB

	// 初始化 mysql 数据库连接
	if err := dao.New(); err != nil {
		zap.S().Errorf("初始化 mysql 数据库连接失败: %s", err)
	}

	// 初始化 redis 连接
	if err := cache.New(); err != nil {
		zap.S().Errorf("初始化 redis 连接失败: %s", err)
	}

	// 设置 api 路由
	setupApiRouter(r)

	// 设置静态资源路由
	setupStaticRouter(r)

	// 设置 websocket 路由
	setupWebsocketRouter(r)

	return r
}
