package dao

import (
	"2024Winter/app/api/global"
	"2024Winter/model"
)

func SelectCartByUsername(username string) ([]*model.Cart, error) {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return nil, tx.Error
	}
	var carts []*model.Cart
	ret := tx.Where("username = ?", username).Find(&carts)
	if ret.Error != nil {
		logger.Info("select carts failed")
		tx.Rollback()
		return nil, ret.Error
	}
	if err := tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return nil, err
	}
	return carts, nil
}

func CreateCart(cart *model.Cart) error {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return tx.Error
	}
	err := tx.Create(cart).Error
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

func DeleteCart(username string, gid int) error {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return tx.Error
	}
	if err := tx.
		Where("username = ? AND gid = ?", username, gid).
		Delete(&model.Cart{}).Error; err != nil {
		logger.Error("mysql delete failed" + err.Error())
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