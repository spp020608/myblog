package service

import (
	"fmt"
	"go.uber.org/zap"
	"myblog/cmd/gorm/query"
	"myblog/model/vo"
	"time"
)

var BlogService = &blogService{}

type blogService struct{}

// GetAllBlog 查询所有文章列表
//
//	@parameter	int64 		  1为按照时间由近到远排列 0为按照由远到近排列
//	@return		*[]vo.Blog 查询到的所有博客
//	@return		err
func (blogService) GetAllBlog(sort string) *[]vo.Blog {
	var blogs []vo.Blog
	if sort == "1" {
		zap.S().Infof("开始正序查询所有博客")
		err := query.Blog.Select(query.Blog.BlogID, query.Blog.Title, query.Blog.Content, query.Blog.PostTime, query.Blog.ReadCount, query.User.UserID, query.User.Username, query.User.Tel, query.Category.CategoryID, query.Category.CategoryName).LeftJoin(query.User, query.User.UserID.EqCol(query.Blog.UserID)).LeftJoin(query.Category, query.Category.CategoryID.EqCol(query.Blog.CategoryID)).Where(query.Blog.IsDelete.Eq(0)).Order(query.Blog.PostTime).Scan(&blogs)

		fmt.Println("blogs为：", blogs)
		if err != nil {
			zap.S().Error("查询所有博客失败", err)
			return nil
		} else {
			zap.S().Info("查询所有博客成功")
			return &blogs
		}
	} else if sort == "0" {

		zap.S().Infof("开始逆序查询所有博客")
		err := query.Blog.Select(query.Blog.BlogID, query.Blog.Title, query.Blog.Content, query.Blog.PostTime, query.Blog.ReadCount, query.User.UserID, query.User.Username, query.User.Tel, query.Category.CategoryID, query.Category.CategoryName).LeftJoin(query.User, query.User.UserID.EqCol(query.Blog.UserID)).LeftJoin(query.Category, query.Category.CategoryID.EqCol(query.Blog.CategoryID)).Where(query.Blog.IsDelete.Eq(0)).Order(query.Blog.PostTime.Desc()).Scan(&blogs)
		fmt.Println("blogs为：", blogs)
		if err != nil {
			zap.S().Error("查询所有博客失败", err)
			return nil
		} else {
			zap.S().Info("查询所有博客成功")
			return &blogs
		}
	} else {
		return nil
	}

}

// BlogWithTitle 提供以文章标题为查询条件的模糊查询接口
//
//	@parameter	string 		  文章标题
//	@return		*[]vo.Blog   查询到的博客
//	@return		err
func (blogService) BlogWithTitle(title string) (blog *[]vo.Blog, err error) {
	zap.S().Infof("开始正序查询标题为%s的博客", title)
	err = query.Blog.Select(query.Blog.BlogID, query.Blog.Title, query.Blog.Content, query.Blog.PostTime, query.Blog.ReadCount, query.User.UserID, query.User.Username, query.User.Tel, query.Category.CategoryID, query.Category.CategoryName).
		LeftJoin(query.User, query.User.UserID.EqCol(query.Blog.UserID)).
		LeftJoin(query.Category, query.Category.CategoryID.EqCol(query.Blog.CategoryID)).
		Where(query.Blog.IsDelete.Eq(0), query.Blog.Title.Like("%"+title+"%")).
		Order(query.Blog.PostTime).Scan(&blog)

	if err != nil {
		zap.S().Error("查询博客失败", err)
		return nil, err
	} else {
		zap.S().Info("查询博客成功:", blog)
		return blog, nil
	}
}

// BlogWithPostTime
//
//	@Description:	以发布时间为条件的查询接口
//	@parameter		begin	string	开始时间
//	@parameter		end								string	截止时间
//	@return			blog 	查询到的博客
//	@return			err							错误信息
func (blogService) BlogWithPostTime(begin time.Time, end time.Time) (blog *[]vo.Blog, err error) {
	zap.S().Infof("开始查询时间%s到时间%s的博客", begin, end)
	err = query.Blog.Select(query.Blog.BlogID, query.Blog.Title, query.Blog.Content, query.Blog.PostTime, query.Blog.ReadCount, query.User.UserID, query.User.Username, query.User.Tel, query.Category.CategoryID, query.Category.CategoryName).
		LeftJoin(query.User, query.User.UserID.EqCol(query.Blog.UserID)).
		LeftJoin(query.Category, query.Category.CategoryID.EqCol(query.Blog.CategoryID)).
		Where(query.Blog.IsDelete.Eq(0), query.Blog.PostTime.Gt(begin), query.Blog.PostTime.Lt(end)).
		Order(query.Blog.PostTime).Scan(&blog)

	if err != nil {
		zap.S().Error("查询博客失败", err)
		return nil, err
	} else {
		zap.S().Info("查询博客成功:", blog)
		return blog, nil
	}

}

// BlogWithReadCount
//
//	@Description:	以阅读量为条件的查询接口
//	@parameter		min						string		最小阅读量
//	@parameter		max						string		最大阅读量
//	@return			blog 			查询到的博客
//	@return			err						错误信息
func (blogService) BlogWithReadCount(min int32, max int32) (blog *[]vo.Blog, err error) {
	zap.S().Infof("开始查询最小阅读量%v到最大阅读量%v的博客", min, max)
	err = query.Blog.Select(query.Blog.BlogID, query.Blog.Title, query.Blog.Content, query.Blog.PostTime, query.Blog.ReadCount, query.User.UserID, query.User.Username, query.User.Tel, query.Category.CategoryID, query.Category.CategoryName).
		LeftJoin(query.User, query.User.UserID.EqCol(query.Blog.UserID)).
		LeftJoin(query.Category, query.Category.CategoryID.EqCol(query.Blog.CategoryID)).
		Where(query.Blog.IsDelete.Eq(0), query.Blog.ReadCount.Gt(min), query.Blog.ReadCount.Lt(max)).
		Order(query.Blog.PostTime).Scan(&blog)

	if err != nil {
		zap.S().Error("查询博客失败", err)
		return nil, err
	} else {
		zap.S().Info("查询博客成功:", blog)
		return blog, nil
	}

}
