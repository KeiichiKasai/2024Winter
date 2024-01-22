package router

import (
	"2024Winter/app/api/internal/middleware"
	"2024Winter/app/api/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", service.Register)
	r.POST("/login", service.LogIn)
	r.POST("/logOut", service.LogOut)
	r.POST("/forget", service.Forget)

}
