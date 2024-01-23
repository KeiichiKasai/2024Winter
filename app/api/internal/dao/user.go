package dao

import (
	"2024Winter/app/api/global"
	"2024Winter/app/api/internal/model"
	"2024Winter/consts"
	"errors"
	"github.com/jinzhu/gorm"
)

var logger = global.Logger
var mdb = global.MDB

func CreateUser(user *model.UserInfo) error {
	tx := mdb.Begin()
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	var temp model.UserInfo
	err := mdb.Where("username = ?", user.Username).First(&temp).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		tx.Rollback()
		return errors.New(consts.MySQLExist)
	}
	err = mdb.Create(&user).Error
	if err != nil {
		logger.Error("mysql insert failed" + err.Error())
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		tx.Rollback()
		return err
	}
	return nil
}
func GetUserByUsername(username string) (*model.UserInfo, error) {
	tx := mdb.Begin() //开启事务

	if tx.Error != nil { //检查事务是否正常开启
		tx.Rollback()
		return nil, tx.Error
	}
	var user model.UserInfo
	if err := mdb.Where("username = ?", username).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &user, nil
}
