package dao

import (
	"2024Winter/app/api/global"
	"2024Winter/model"
)

func SelectWallet(username string) (*model.Wallet, error) {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return nil, tx.Error
	}
	var w *model.Wallet
	if err := tx.
		Where("username = ?", username).
		Find(w).Error; err != nil {
		logger.Error("mysql insert failed" + err.Error())
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return nil, err
	}
	return w, nil
}

func UpdateBalance(wallet *model.Wallet) error {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return tx.Error
	}
	if err := tx.
		Where("username = ?", wallet.Username).
		Update(wallet).Error; err != nil {
		logger.Error("mysql update failed" + err.Error())
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
