package dao

import (
	"2024Winter/app/api/global"
	"2024Winter/model"
)

func PutAllGoods() ([]*model.Good, error) {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return nil, tx.Error
	}
	var goods []*model.Good
	ret := tx.Find(&goods)
	if ret.Error != nil {
		logger.Error("select goods failed")
		tx.Rollback()
		return nil, ret.Error
	}
	if err := tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return nil, err
	}
	return goods, nil
}

func SearchGoods(keyword string) ([]*model.Good, error) {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return nil, tx.Error
	}
	var goods []*model.Good
	ret := tx.Where("gname LIKE ?", "%"+keyword+"%").Find(&goods)
	if ret.Error != nil {
		logger.Info("select goods failed")
		tx.Rollback()
		return nil, ret.Error
	}
	if err := tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return nil, err
	}
	return goods, nil
}

func SearchGoodsByOid(oid int) ([]*model.Good, error) {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return nil, tx.Error
	}
	var goods []*model.Good
	ret := tx.Where("gname = ?", oid).Find(&goods)
	if ret.Error != nil {
		logger.Info("select goods failed")
		tx.Rollback()
		return nil, ret.Error
	}
	if err := tx.Commit().Error; err != nil { //提交事务并判断是否成功提交
		logger.Error("tx close failed")
		tx.Rollback()
		return nil, err
	}
	return goods, nil
}

func AddGood(good *model.Good) error {
	tx := global.MDB.Begin()
	if tx.Error != nil {
		logger.Error("tx open failed")
		tx.Rollback()
		return tx.Error
	}
	err := tx.Create(&good).Error
	if err != nil {
		logger.Error("mysql insert failed" + err.Error())
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
