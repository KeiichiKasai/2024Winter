package service

import (
	"2024Winter/app/api/global"
	"2024Winter/app/api/internal/dao"
	"2024Winter/app/api/internal/middleware"
	"2024Winter/model"
	"2024Winter/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
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
			"code":  "0",
			"msg":   "用户名或密码为空",
			"token": "null",
		})
		return
	}
	var temp *model.UserInfo
	temp, err = dao.GetUserByUsername(user.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "0",
			"msg":   "用户不存在",
			"token": "null",
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
			"code":  "0",
			"msg":   "密码错误",
			"token": "null",
		})
		return
	}
	token, err := middleware.GenToken(user.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "0",
			"msg":   "无法生成token",
			"token": "null",
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
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  "修改成功",
	})
}

func LogOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  "clear jwt-token",
		"user": c.GetString("username"),
	})
}

func GetInfo(c *gin.Context) {
	username := c.GetString("username")
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "获取用户个人信息失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg": model.UserInfo{
			Uid:      user.Uid,
			Username: user.Username,
			Address:  user.Address,
			Gender:   user.Gender,
			Birthday: user.Birthday,
		},
	})
}

func ChangeInfo(c *gin.Context) {
	address := c.PostForm("address")
	genderStr := c.PostForm("gender")
	birthdayStr := c.PostForm("birthday")
	if address == "" || genderStr == "" || birthdayStr == "" {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "信息不全",
		})
	}
	gender, err := strconv.Atoi(genderStr)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "性别错误",
		})
		return
	}
	birthday, err := utils.ParseBirthday(birthdayStr)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "生日格式错误",
		})
		return
	}
	username := c.GetString("username")
	usr, err := dao.GetUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "获取用户个人信息失败",
		})
		return
	}
	temp := &model.UserInfo{
		Uid:      usr.Uid,
		Username: usr.Username,
		Password: usr.Password,
		Phone:    usr.Phone,
		Address:  address,
		Gender:   gender,
		Birthday: birthday,
	}
	err = dao.UpdateUser(temp)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "0",
			"msg":  "更改用户个人信息失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "1",
		"msg":  "更改用户个人信息成功",
	})
}
