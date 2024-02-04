package templates

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}
func LogIn(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}
func No(c *gin.Context) {
	c.HTML(200, "404.html", gin.H{})
}
func Info(c *gin.Context) {
	c.HTML(200, "user_info.html", gin.H{})
}
func Product(c *gin.Context) {
	c.HTML(200, "product.html", gin.H{})
}
