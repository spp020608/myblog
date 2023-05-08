package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// InternalServerError 服务器内部错误
//	@parameter	msg 消息文本
func InternalServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, NewCommonResult(http.StatusInternalServerError, msg, nil))
}

//
// Unauthorized 用户未授权的时候返回给前端的提示信息
//	@parameter	ctx *gin.Context
//	@parameter	msg 消息文本
//
func Unauthorized(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusUnauthorized, NewCommonResult(http.StatusUnauthorized, msg, nil))
}

//
// Ok 接口调用成功的时候返回给前端的信息提示
//	@parameter	ctx *gin.Context
//	@parameter	msg 消息文本
//	@parameter	data 附加数据
//
func Ok(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, NewCommonResult(http.StatusOK, msg, data))
}

//
// NoData 没有符合条件的数据
//
func NoData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewCommonResult(http.StatusOK, "接口调用成功, 没有符合条件的数据", nil))
}

