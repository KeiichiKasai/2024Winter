package service

import (
	"2024Winter/app/api/internal/dao"
	"2024Winter/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Buy(c *gin.Context) {
	username := c.GetString("username")
	carts, err := dao.SelectCartByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "查询购物车失败",
		})
		return
	}
	var sum float64 = 0
	for _, cart := range carts {
		sum += cart.Price * float64(cart.Count)
	}
	wallet, err := dao.SelectWallet(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "查询钱包失败",
		})
		return
	}
	if sum > wallet.Balance {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "余额不足",
		})
		return
	}
	now := wallet.Balance - sum
	temp := model.Wallet{
		Username: username,
		Balance:  now,
	}
	err = dao.UpdateWallet(&temp)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "购买失败",
		})
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "购买成功",
	})
}

func Recharge(c *gin.Context) {
	moneyStr := c.PostForm("money")
	username := c.GetString("username")
	money, err := strconv.ParseFloat(moneyStr, 64)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "金额错误",
		})
		return
	}
	wallet, err := dao.SelectWallet(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "查询钱包失败",
		})
		return
	}
	now := wallet.Balance + money
	temp := model.Wallet{
		Username: username,
		Balance:  now,
	}
	err = dao.UpdateWallet(&temp)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "充值失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "充值成功",
	})
}
