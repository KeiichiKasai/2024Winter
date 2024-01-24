package service

import (
	"2024Winter/app/api/global"
	"2024Winter/app/api/internal/dao"
	"2024Winter/app/api/internal/middleware"
	"2024Winter/model"
	"2024Winter/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var logger = global.Logger

func Register(c *gin.Context) {
	var user model.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		logger.Warn("shouldBind failed:" + err.Error())
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "用户名或密码为空",
		})
		return
	}
	if len(user.Phone) != 11 {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "手机号不合法",
		})
		return
	}
	_, err = dao.GetUserByUsername(user.Username)
	if err != gorm.ErrRecordNotFound {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "用户已存在",
		})
		return
	} else if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "注册失败",
		})
		return
	}
	err = dao.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{
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
	var user model.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		logger.Warn("shouldBind failed:" + err.Error())
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "用户名或密码为空",
		})
		return
	}
	var temp *model.UserInfo
	temp, err = dao.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "用户不存在",
		})
		return
	}
	password, err := utils.Decrypt(temp.Password)
	if err != nil {
		logger.Warn("解密错误：" + err.Error())
		return
	}
	if password != user.Password {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "密码错误",
		})
		return
	}
	token, err := middleware.GenToken(user.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "无法生成token",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  "1",
		"msg":   "登录成功",
		"token": token,
	})
}

func Forget(c *gin.Context) {
	newPass := c.PostForm("newpassword")
	var user model.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		logger.Warn("shouldBind failed:" + err.Error())
	}
	if user.Username == "" {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "用户名或密码为空",
		})
		return
	}
	if len(user.Phone) != 11 {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "手机号不合法",
		})
		return
	}
	usr, err := dao.GetUserByUsername(user.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(500, gin.H{
				"code": "0",
				"msg":  "用户不存在",
			})
			return
		}
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "查找用户出错",
		})
		return
	}
	phone, err := utils.Decrypt(usr.Phone)
	if err != nil {
		logger.Warn("解密错误：" + err.Error())
		return
	}
	if phone != usr.Phone {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "验证失败",
		})
		return
	}
	newPassStr, err := utils.Encrypt(newPass)
	if err != nil {
		logger.Warn("加密错误：" + err.Error())
		return
	}
	temp := &model.UserInfo{
		Uid:      usr.Uid,
		Username: usr.Username,
		Password: newPassStr,
		Phone:    usr.Phone,
	}
	err = dao.UpdateUser(temp)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "修改失败",
		})
		return
	}
}

func LogOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  "clear jwt-token",
		"user": c.Get("username"),
	})
}
