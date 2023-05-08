package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myblog/cmd/gorm/model"
	"myblog/cmd/gorm/query"
	"myblog/internal/repository/cache"
	"myblog/model/dto"
	"myblog/pkg/config"
	"myblog/pkg/ctime"
	"myblog/pkg/util"
	"time"
)

// 初始化Gorm Gen
func init() {
	cfg := config.Get().MySqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		zap.S().Info(err)
	}
	query.SetDefault(db)

}

var UserService = &userService{}

type userService struct{}

/*
1.用户注册前端传递到后端的数据必须包含如下信息︰用户名、密码、电话号码
2.注册的时候需判断用户名是否重名
3.验证前端传递的数据是否合法（比如:数据库中的username字段要求是长度为32的字符串，你传递了超过该长度的字符串)
4.用户注册的时候，密码需要加密存储(bcrypt方式加密)
*/

// Register 用户注册
//
//	@parameter	registerParam
//	@return		rowsAffected 此次操作影响的行数, 注册成功, 该返回值为 1
//	@return		err
func (userService) Register(registerParam *dto.RegisterParam) (rowsAffected int64, err error) {

	zap.S().Infof("开始对明文密码进行加密, 明文密码为: %s\n", registerParam.Password)
	hashedPassword, err := util.GenerateFromPassword(registerParam.Password)
	zap.S().Infof("加密之后的密码为: %s\n", hashedPassword)
	if err != nil {
		zap.S().Errorf("密码加密失败: %s\n", err)
		return 0, err
	}

	user := &model.User{
		Username:   registerParam.Username,
		Password:   hashedPassword,
		Tel:        registerParam.Tel,
		CreateTime: ctime.CTime{Time: time.Now()}.Time,
		IsDelete:   0,
	}

	/*	if err = dao.DB.Transaction(func(tx *gorm.DB) error {
			var e error
			rowsAffected, e = dao.UserDao.Insert(tx, user)
			if e != nil {
				return e
			}
			return nil
		}); err != nil {
			return 0, err
		}*/
	find, err := query.User.Where(query.User.Username.In(user.Username)).Find()
	if err != nil {
		zap.S().Error(err)
	}
	fmt.Println(find)
	if len(find) == 0 {
		zap.S().Infof("用户注册接收到的密码为:%s\n", registerParam.Password)

		// 对密码进行加密
		zap.S().Infof("开始对明文密码进行加密, 明文密码为: %s\n", registerParam.Password)
		hashedPassword, err := util.GenerateFromPassword(registerParam.Password)
		zap.S().Infof("加密之后的密码为: %s\n", hashedPassword)
		if err != nil {
			zap.S().Errorf("密码加密失败: %s\n", err)
			return 0, err
		}
		query.User.Create(user)
	} else {
		zap.S().Info(user.Username + "已存在")
	}

	return rowsAffected, nil
}

// Login 用户登录
//
//	@parameter	loginInfo 前端传递的登录信息
//	@return		token 登录成功之后生成的 token
//	@return		err
func (userService) Login(loginInfo *dto.LoginInfo) (token string, err error) {
	zap.S().Infof("用户登录前端传递过来的密码为: %s", loginInfo.Password)
	// 判断传入参数是否为空
	if loginInfo == nil {
		zap.L().Info("登录信息为空 Login 直接返回")
		return "", errors.New("登录信息为空")
	}

	// 根据用户名从数据库中查询用户
	find, err := query.User.Where(query.User.Username.In(loginInfo.Username)).Find()
	if err != nil {
		zap.S().Error(err)
	}
	user := find[0]
	if err != nil {
		zap.S().Errorf("从数据库中查询用户名为: %s 的用户失败: %s", loginInfo.Username, err)
		return "", err
	}

	zap.S().Infof("从数据库中获取的用户密码为:%s", user.Password)
	zap.S().Infof("接收到的前端用户用户密码为:%s", loginInfo.Password)

	// 比对数据库中的密码(加密之后的)和传递过来的密码(明文密码)
	err = util.CompareHashAndPassword(user.Password, loginInfo.Password)
	if err != nil {
		zap.S().Errorf("用户名或者密码错误: %s", err)
		return "", err
	}

	// 密码比对成功之后, 生成 token 字符串
	generateToken, err := util.GenerateToken(user.Username)
	if err != nil {
		zap.S().Errorf("根据用户名生成 token 失败: %s\n", err)
		return "", err
	}

	// 登录成功之后, 将登录用户信息保存到 redis 中
	if err = cache.SaveLoginUser(user); err != nil {
		zap.S().Infof("将登录用户存放到 redis 中失败: %s", err)
	}

	return generateToken, nil
}

// UpdatePassword 修改密码
//
//	@parameter	updatePassword 前端传递的登录信息
//	@return		rowAffected
//	@return		err
func (userService) UpdatePassword(updatePassword *dto.UpdatePassword) (rowsAffected int64, err error) {

	zap.S().Infof("用户修改密码前端传递过来的新密码为: %s", updatePassword.NewPassword)
	// 判断传入参数是否为空
	if updatePassword == nil {
		zap.L().Info("登录信息为空 Login 直接返回")
		return 0, errors.New("登录信息为空")
	}

	// 根据用户名从数据库中查询用户
	find, err := query.User.Where(query.User.Username.In(updatePassword.Username)).Find()
	if err != nil {
		zap.S().Error(err)
	}
	user := find[0]
	if err != nil {
		zap.S().Errorf("从数据库中查询用户名为: %s 的用户失败: %s", updatePassword.Username, err)
		return 0, err
	}
	zap.S().Infof("从数据库中获取的用户密码为:%s", user.Password)
	zap.S().Infof("接收到的前端用户用户密码为:%s", updatePassword.Password)

	// 比对数据库中的密码(加密之后的)和传递过来的密码(明文密码)
	err = util.CompareHashAndPassword(user.Password, updatePassword.Password)
	if err != nil {
		zap.S().Errorf("用户名或者密码错误: %s", err)
		return 0, err
	}

	hashedPassword, err := util.GenerateFromPassword(updatePassword.NewPassword)
	if err != nil {
		zap.S().Errorf("密码加密失败: %s\n", err)
		return 0, err
	}

	user.Password = hashedPassword
	query.User.Save(user)
	zap.L().Info("修改密码成功" + updatePassword.NewPassword)
	return 0, nil
}

// UpdateUser 修改用户信息
//
//	@parameter	User 前端传递的登录信息
//	@return		string
//	@return		err
func (userService) UpdateUser(user *dto.User) (s string, err error) {
	zap.S().Info("用户修改的信息为:", user)

	hashedPassword, err := util.GenerateFromPassword(user.Password)
	if err != nil {
		zap.S().Errorf("密码加密失败: %s\n", err)
		return "", err
	}

	_, err = query.User.Where(query.User.Username.Eq(user.Username)).Updates(model.User{
		Username: user.Username,
		Password: hashedPassword,
		Tel:      user.Tel,
	})
	if err != nil {
		zap.L().Info(err.Error())
		return "", err
	}
	return "更新成功", nil
}

// DelUser 注销用户
//
//	@parameter	loginInfo 前端传递的登录信息
//	@return		rowAffected
//	@return		err
func (userService) DelUser(info *dto.LoginInfo) (rowsAffected int64, err error) {
	zap.L().Info("要注销的用户为:" + info.Username)
	first, err := query.User.Where(query.User.Username.Eq(info.Username), query.User.Password.Eq(info.Password)).First()
	if err != nil {
		zap.S().Errorf("注销用户验证用户名&密码出先错误" + err.Error())
	}
	if first == nil {
		zap.S().Info("查询的用户为空")
		return 0, nil
	} else {
		first.IsDelete = 1
		query.User.Save(first)
		zap.S().Infof("用户注销成功")
		rowsAffected = 1
	}
	return rowsAffected, nil
}
