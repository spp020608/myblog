package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myblog/internal/pkg/resp"
	"myblog/internal/service"
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
//	@Param			sort	query		int				true	"规则: 正序为 1, 逆序为0"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/blog/getAllBlog [post]
func (blogApi) GetAllBlog(c *gin.Context) {
	fmt.Println("int为")

	query := c.Query("sort")
	zap.S().Infof("query:", query)
	/*	if err := c.ShouldBind(&number); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}*/

	blog := service.BlogService.GetAllBlog(query)
	if blog != nil {
		resp.Ok(c, "查询文章成功", &blog)
	} else {
		resp.InternalServerError(c, "查询文章失败")
	}

}
