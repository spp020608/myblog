package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myblog/internal/pkg/resp"
	"myblog/internal/service"
	"myblog/model/dto"
	"myblog/pkg/validate"
)

var CommentApi = new(commentApi)

type commentApi struct{}

// AddComment 添加评论
//
//	@Summary		添加评论
//	@Description	添加评论
//	@Tags			commentApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			content	formData	string			true	"评论内容"
//	@Param			userID	formData	int32			true	"用户ID"
//	@Param			blogID	formData	int32			true	"BlogID"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/comment/addComment [post]
func (commentApi) AddComment(c *gin.Context) {
	var comment dto.Comment
	// 将请求数据绑定到结构体
	if err := c.ShouldBind(&comment); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}

	addComment, err := service.CommentService.AddComment(&comment)
	if err != nil {
		zap.S().Errorf("添加评论失败: %s\n", err)
		resp.InternalServerError(c, "添加评论失败")
		return
	}
	resp.Ok(c, "添加评论成功", addComment)
}
