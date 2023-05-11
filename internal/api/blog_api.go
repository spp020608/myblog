package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myblog/internal/pkg/resp"
	"myblog/internal/service"
	"myblog/pkg/ctime"
	"strconv"
	"time"
)

var BlogApi = new(blogApi)

type blogApi struct{}

// GetAllBlog 获取所有博客
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			blogApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			sort	query		int				true	"规则: 正序为 1, 逆序为0"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/blog/getAllBlog [post]
func (blogApi) GetAllBlog(c *gin.Context) {
	query := c.Query("sort")
	zap.S().Info("query:", query)
	blog := service.BlogService.GetAllBlog(query)
	if blog != nil {
		resp.Ok(c, "查询文章成功", &blog)
	} else {
		resp.InternalServerError(c, "查询文章失败")
	}

}

// BlogWithTitle 提供以文章标题为查询条件的模糊查询接口
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			blogApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			title	query		string			true	"输入标题模糊查询"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/blog/blogWithTitle [post]
func (blogApi) BlogWithTitle(c *gin.Context) {
	title := c.Query("title")
	zap.S().Info("要查询的博客title为:", title)
	blogWithTitle, err := service.BlogService.BlogWithTitle(title)
	if err != nil {
		zap.S().Error("查询发生错误", err.Error())
	} else {
		if blogWithTitle != nil {
			resp.Ok(c, "查询文章成功", &blogWithTitle)
		} else {
			resp.Ok(c, "没有找到文章", &blogWithTitle)
		}
	}
}

// BlogWithPostTime 以发布时间为条件的查询接口
//
//	@Summary		Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Description	Post 请求, 发送 multipart/form-data 类型的表单数据, 参数在消息体中
//	@Tags			blogApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			begin	query		string			true	"开始时间"
//	@Param			end		query		string			true	"截止时间"
//	@Success		200		{object}	swagger.HttpOk	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/blog/blogWithPostTime [post]
func (blogApi) BlogWithPostTime(c *gin.Context) {
	begin := c.Query("begin")
	end := c.Query("end")
	beginTime, _ := time.ParseInLocation(ctime.TimeLayout, begin, time.Local)
	endTime, _ := time.ParseInLocation(ctime.TimeLayout, end, time.Local)
	zap.S().Infof("开始时间为%s: ,开始时间为%s:", begin, end)
	blog, err := service.BlogService.BlogWithPostTime(beginTime, endTime)
	if err != nil {
		zap.Error(err)
	} else {
		if blog != nil {
			resp.Ok(c, "查询文章成功", &blog)
		} else {
			resp.InternalServerError(c, "查询文章失败")
		}
	}

}

// BlogWithReadCount 以阅读量为条件的查询接口
//
//	@Summary		以阅读量为条件的查询接口
//	@Description	以阅读量为条件的查询接口
//	@Tags			blogApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			min	query		int32			true	"最小阅读量"
//	@Param			max	query		int32			true	"最大阅读量"
//	@Success		200	{object}	swagger.HttpOk	"OK"
//	@Failure		400	{object}	swagger.Http400	"Bad Request"
//	@Failure		404	{object}	swagger.Http404	"Page Not Found"
//	@Failure		500	{object}	swagger.Http500	"InternalError"
//	@Router			/blog/blogWithReadCount [post]
func (blogApi) BlogWithReadCount(c *gin.Context) {
	min := c.Query("min")
	max := c.Query("max")
	zap.S().Infof("开始查询最小阅读量%v到最大阅读量%v的博客", min, max)
	minR, err := strconv.Atoi(min)
	maxR, err := strconv.Atoi(max)
	blog, err := service.BlogService.BlogWithReadCount(int32(minR), int32(maxR))
	if err != nil {
		zap.Error(err)
	} else {
		if blog != nil {
			resp.Ok(c, "查询文章成功", &blog)
		} else {
			resp.InternalServerError(c, "查询文章失败")
		}
	}

}

// BlogWithCategoryName 以阅读量为条件的查询接口
//
//	@Summary		以文章类别为条件的查询接口
//	@Description	以文章类别为条件的查询接口
//	@Tags			blogApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			categoryName	query		string			true	"文章类别名"
//	@Success		200				{object}	swagger.HttpOk	"OK"
//	@Failure		400				{object}	swagger.Http400	"Bad Request"
//	@Failure		404				{object}	swagger.Http404	"Page Not Found"
//	@Failure		500				{object}	swagger.Http500	"InternalError"
//	@Router			/blog/blogWithCategoryName [post]
func (blogApi) BlogWithCategoryName(c *gin.Context) {
	categoryName := c.Query("categoryName")
	zap.S().Infof("开始查询类别为%s的博客", categoryName)
	blog, err := service.BlogService.BlogWithCategoryName(categoryName)
	if err != nil {
		zap.Error(err)
	} else {
		if blog != nil {
			resp.Ok(c, "查询文章成功", &blog)
		} else {
			resp.InternalServerError(c, "查询文章失败")
		}
	}

}

// BlogWithMoreCondition 一堆条件的查询接口
//
//	@Summary		一堆条件的查询接口
//	@Description	一堆条件的查询接口
//	@Tags			blogApi
//	@Accept			mpfd
//	@Produce		json
//	@Param			blogId			query		int				false	"博客ID"
//	@Param			title			query		string			false	"文章标题"
//	@Param			begin			query		string			false	"发布最早时间"
//	@Param			end				query		string			false	"发布最晚时间"
//	@Param			min				query		int32			false	"阅读最小次数"
//	@Param			max				query		int32			false	"阅读最多次数"
//	@Param			categoryName	query		string			false	"文章分类"
//	@Success		200				{object}	swagger.HttpOk	"OK"
//	@Failure		400				{object}	swagger.Http400	"Bad Request"
//	@Failure		404				{object}	swagger.Http404	"Page Not Found"
//	@Failure		500				{object}	swagger.Http500	"InternalError"
//	@Router			/blog/blogWithMoreCondition [post]
func (blogApi) BlogWithMoreCondition(c *gin.Context) {
	blogId := c.Query("blogId")
	blogID, _ := strconv.Atoi(blogId)

	title := c.Query("title")

	begin := c.Query("begin")
	end := c.Query("end")
	beginTime, _ := time.ParseInLocation(ctime.TimeLayout, begin, time.Local)
	endTime, _ := time.ParseInLocation(ctime.TimeLayout, end, time.Local)

	min := c.Query("min")
	max := c.Query("max")
	minR, _ := strconv.Atoi(min)
	maxR, _ := strconv.Atoi(max)

	categoryName := c.Query("categoryName")

	blog, err := service.BlogService.BlogWithMoreCondition(int32(blogID), title, beginTime, endTime, int32(minR), int32(maxR), categoryName)
	if err != nil {
		zap.Error(err)
	} else {
		if blog != nil {
			resp.Ok(c, "查询文章成功", &blog)
		} else {
			resp.InternalServerError(c, "查询文章失败")
		}
	}
}
