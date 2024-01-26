package dao

import (
	"2024Winter/app/api/global"
	"2024Winter/consts"
	"2024Winter/model"
	"2024Winter/utils"
	"errors"
	"github.com/jinzhu/gorm"
)

var logger = global.Logger

func CreateUser(user *model.UserInfo) error {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return tx.Error
	}
	var temp model.UserInfo
	var err error
	password, err := utils.Encrypt(user.Password)
	if err != nil {
		logger.Info("crypt failed:" + err.Error())
		return err
	}
	phone, err := utils.Encrypt(user.Phone)
	if err != nil {
		logger.Info("crypt failed:" + err.Error())
		return err
	}
	err = tx.Where("username = ?", user.Username).First(&temp).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		tx.Rollback()
		return errors.New(consts.MySQLExist)
	}
	temp = model.UserInfo{
		Username: user.Username,
		Password: password,
		Phone:    phone,
	}
	err = tx.Create(&temp).Error
	if err != nil {
		logger.Error("mysql insert failed" + err.Error())
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return err
	}
	return nil
}
func GetUserByUsername(username string) (*model.UserInfo, error) {
	tx := global.MDB.Begin() //开启事务

	if tx.Error != nil { //检查事务是否正常开启
		logger.Error("tx open failed")
		tx.Rollback()
		return nil, tx.Error
	}
	var user model.UserInfo
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		logger.Error("tx close failed")
		tx.Rollback()
		return nil, err
	}
	return &user, nil
}
func UpdateUser(user *model.UserInfo) error {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return tx.Error
	}
	if err := tx.
		Model(&model.UserInfo{}).
		Updates(user).
		Where("username = ?", user.Username).Error; err != nil {
		logger.Warn("update user failed")
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return err
	}
	return nil
}
