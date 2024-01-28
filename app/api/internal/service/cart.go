package service

import (
	"2024Winter/app/api/internal/dao"
	"2024Winter/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCart(c *gin.Context) {
	username := c.GetString("username")
	carts, err := dao.SelectCartByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "未添加商品至购物车",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  carts,
	})
}

func AddCart(c *gin.Context) {
	username := c.GetString("username")
	var cart model.Cart
	err := c.ShouldBind(&cart)
	if err != nil {
		logger.Warn("shouldBind failed:" + err.Error())
		return
	}
	cart.Username = username
	err = dao.CreateCart(&cart)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "添加购物车失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "已添加至购物车",
	})
}

func DelCart(c *gin.Context) {
	username := c.GetString("username")
	gid, _ := strconv.Atoi(c.Query("gid"))
	err := dao.DeleteCart(username, gid)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 0,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"msg":  "删除成功",
	})
}
