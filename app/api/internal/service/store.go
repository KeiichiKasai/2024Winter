package service

import (
	"2024Winter/app/api/internal/dao"
	"2024Winter/model"
	"github.com/gin-gonic/gin"
)

func PushGood(c *gin.Context) {
	username := c.GetString("username")
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "获取个人信息失败",
		})
		return
	}
	oid := user.Uid
	var good model.Good
	err = c.ShouldBind(&good)
	if err != nil {
		logger.Warn("shouldBind failed:" + err.Error())
		return
	}
	var temp = &model.Good{
		Gname:    good.Gname,
		Category: good.Category,
		Picture:  good.Picture,
		Price:    good.Price,
		OwnerId:  oid,
	}
	err = dao.AddGood(temp)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "添加商品失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  "添加商品成功",
	})
}

func GetOwnGood(c *gin.Context) {
	username := c.GetString("username")
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "获取个人信息失败",
		})
		return
	}
	oid := user.Uid
	goods, err := dao.SearchGoodsByOid(oid)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "获取商品信息失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  goods,
	})
}
