package service

import (
	"2024Winter/app/api/global"
	"2024Winter/app/api/internal/dao"
	"2024Winter/app/api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var logger = global.Logger

func Register(c *gin.Context) {
	var user model.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		logger.Warn("get form failed:" + err.Error())
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(200, gin.H{
			"code": "0",
			"msg":  "用户名或密码为空",
		})
		return
	}
	if len(user.Phone) != 11 {
		c.JSON(200, gin.H{
			"code": "0",
			"msg":  "手机号不合法",
		})
		return
	}
	_, err = dao.GetUserByUsername(user.Username)
	if err != gorm.ErrRecordNotFound {
		c.JSON(200, gin.H{
			"code": "0",
			"msg":  "用户已存在",
		})
		return
	}
	err = dao.CreateUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": "0",
			"msg":  "注册失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  "注册成功",
	})
}
func LogIn(c *gin.Context) {

}
func Forget(c *gin.Context) {

}

func LogOut(c *gin.Context) {

}
