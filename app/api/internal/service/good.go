package service

import (
	"2024Winter/app/api/internal/dao"
	"2024Winter/model"
	"github.com/gin-gonic/gin"
)

func GetGoods(c *gin.Context) {
	var goods []*model.Good
	var err error
	goods, err = dao.PutAllGoods()
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "获取商品失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  goods,
	})
}

func SearchGoods(c *gin.Context) {
	keyword := c.Query("keyword")
	goods, err := dao.SearchGoods(keyword)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "无法获取商品信息",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  goods,
	})
}
