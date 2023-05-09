package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myblog/internal/pkg/resp"
	"myblog/internal/service"
	"myblog/pkg/validate"
)

var BlogApi = new(blogApi)

type blogApi struct{}

// GetAllBlog 用户注册
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			userApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			sort	query		int				false	"规则: 正序为 1, 逆序为0"	minimum(1)	maximum(1)
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/blog/getAllBlog [post]
func (blogApi) GetAllBlog(c *gin.Context) {
	var number int64
	fmt.Println("int为")
	if err := c.ShouldBind(&number); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}

	blog := service.BlogService.GetAllBlog(number)
	resp.Ok(c, "查询所有用户成功", &blog)

}
