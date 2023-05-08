package dao

import (
	"gorm.io/gorm"
	"myblog/cmd/gorm/model"
)

const userTableName = "user"

var UserDao = new(userDao)

type userDao struct{}



type IUserDao interface {
	Insert(db *gorm.DB, user *model.User) (rowsAffected int64, err error)

}
//
// Insert 添加用户
//	@receiver	u
//	@parameter	user 要添加到数据库中的用户对象
//	@return		rowsAffected 此次操作影响的行数
//	@return		err
//
func (u userDao) Insert(db *gorm.DB, user *model.User) (rowsAffected int64, err error) {
	tx := db.Table(userTableName).Create(&user)
	if tx.Error != nil {
		return 0, err
	}
	return rowsAffected, nil
}