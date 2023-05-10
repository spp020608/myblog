package service

import (
	"fmt"
	"go.uber.org/zap"
	"myblog/cmd/gorm/query"
	"myblog/model/vo"
)

var BlogService = &blogService{}

type blogService struct{}

// GetAllBlog 查询所有文章列表
//
//	@parameter	int64 		  1为按照时间由近到远排列 0为按照由远到近排列
//	@return		[]blogService 查询到的所有博客
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
