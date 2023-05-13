package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myblog/internal/pkg/resp"
	"myblog/internal/service"
	"myblog/model/dto"
	"myblog/pkg/validate"
	"strconv"
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

// DelComment 添加评论
//
//	@Summary		删除评论
//	@Description	删除评论
//	@Tags			commentApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			commentID	query		int32			true	"commentID"
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/comment/delComment [post]
func (commentApi) DelComment(c *gin.Context) {

	query := c.Query("commentID")
	commentID, _ := strconv.Atoi(query)
	comment, err := service.CommentService.DelComment(int32(commentID))
	if err != nil {
		zap.S().Errorf("删除评论失败: %s\n", err)
		resp.InternalServerError(c, "删除评论失败")
		return
	}
	if comment != 0 {
		resp.Ok(c, "删除评论成功", comment)
	} else {
		resp.NoData(c)
	}

}

// UpdateComment 添加评论
//
//	@Summary		修改评论
//	@Description	修改评论
//	@Tags			commentApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			commentID	formData	int32			true	"评论ID"
//	@Param			content		formData	string			true	"评论内容"
//	@Success		200			{object}	swagger.HttpOk	"OK"
//	@Failure		400			{object}	swagger.Http400	"Bad Request"
//	@Failure		404			{object}	swagger.Http404	"Page Not Found"
//	@Failure		500			{object}	swagger.Http500	"InternalError"
//	@Router			/comment/updateComment [post]
func (commentApi) UpdateComment(c *gin.Context) {

	var comment dto.Comment
	// 将请求数据绑定到结构体
	if err := c.ShouldBind(&comment); err != nil {
		trans := validate.TranslateError(err)
		zap.S().Errorf("绑定参数失败: %s\n", trans)
		resp.InternalServerError(c, trans)
		return
	}

	updateComment, err := service.CommentService.UpdateComment(&comment)
	if err != nil {
		zap.S().Errorf("修改评论失败: %s\n", err)
		resp.InternalServerError(c, "修改评论失败")
		return
	}
	resp.Ok(c, "修改评论成功", updateComment)

}

// QueryCommentByBlogId 提供以BlogId为查询条件的查询接口
//
//	@Summary		提供以BlogId为查询条件的查询接口
//	@Description	提供以BlogId为查询条件的查询接口
//	@Tags			commentApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			BlogId	query		string			true	"输入BlogId查询"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/comment/queryCommentByBlogId [post]
func (commentApi) QueryCommentByBlogId(c *gin.Context) {
	BlogId := c.Query("BlogId")
	zap.S().Info("要查询的评论的BlogId为:", BlogId)
	id, _ := strconv.Atoi(BlogId)
	comments, err := service.CommentService.QueryCommentByBlogId(int32(id))
	if err != nil {
		zap.S().Error("查询发生错误", err.Error())
	} else {
		if comments != nil {
			resp.Ok(c, "查询评论成功", &comments)
		} else {
			resp.NoData(c)
		}
	}
}

// CountCommentByBlogId 开始统计BlogId下的评论数量
//
//	@Summary		Post 开始统计BlogId下的评论数量
//	@Description	Post 开始统计BlogId下的评论数量
//	@Tags			commentApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			BlogId	query		string			true	"输入BlogId查询"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/comment/countCommentByBlogId [post]
func (commentApi) CountCommentByBlogId(c *gin.Context) {
	BlogId := c.Query("BlogId")
	zap.S().Infof("开始统计BlogId为%v下的评论数量", BlogId)
	id, _ := strconv.Atoi(BlogId)
	count, err := service.CommentService.CountCommentByBlogId(int32(id))
	if err != nil {
		zap.S().Error("统计发生错误", err.Error())
	} else {
		if count != 0 {
			resp.Ok(c, "统计评论数量成功", &count)
		} else {
			resp.NoData(c)
		}
	}
}

// QueryCommentByUserId 提供以UserId为查询条件的查询接口
//
//	@Summary		提供以UserId为查询条件的查询接口
//	@Description	提供以UserId为查询条件的查询接口
//	@Tags			commentApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			UserId	query		string			true	"输入UserId查询"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/comment/queryCommentByUserId [post]
func (commentApi) QueryCommentByUserId(c *gin.Context) {
	UserId := c.Query("UserId")
	zap.S().Info("要查询的评论的BlogId为:", UserId)
	id, _ := strconv.Atoi(UserId)
	comments, err := service.CommentService.QueryCommentByUserId(int32(id))
	if err != nil {
		zap.S().Error("查询发生错误", err.Error())
	} else {
		if comments != nil {
			resp.Ok(c, "查询评论成功", &comments)
		} else {
			resp.NoData(c)
		}
	}
}
