package midware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//
// ZapLoggerMiddleware 自定义 Logger 中间件
//	@return	gin.HandlerFunc
//
func ZapLoggerMiddleware() gin.HandlerFunc {
	// 要输出的信息 日期 日志级别 函数调用 自定义信息
	return func(ctx *gin.Context) {
		start := time.Now() // 获取当前时间, 用于记录此次请求的开始时间

		path := ctx.Request.URL.Path      // 请求路径
		query := ctx.Request.URL.RawQuery // 请求的原始 Query 信息

		ctx.Next() // 将请求继续向处理链的下游传递

		cost := time.Since(start) // 获取当前时间和请求开始时间的时间差, 这个时间差就是此次请求耗费的时长

		// 使用 zap 输出日志
		zap.L().Info(path, // 请求路径
			zap.Int("status", ctx.Writer.Status()),   // 状态码
			zap.String("method", ctx.Request.Method), // 请求方法
			zap.String("path", path),                 // 请求路径
			zap.String("query", query),               // 查询
			zap.String("ip", ctx.ClientIP()),         // 谁发起的请求
			//zap.String("user-agent", context.Request.UserAgent()),                      // 用户代理, 一般输出用户使用的浏览器信息
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()), //错误信息
			zap.Duration("cost", cost), // 此次请求耗费的时长
		)

	}

}
