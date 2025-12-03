package controller

import (
	"BLOG/models"
	"BLOG/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func UserControllerInit(r *gin.Engine) {

	v1 := r.Group("user")
	{
		v1.POST("/register", Register)
		v1.POST("/login", Login)
		v1.GET("/logout", AuthMiddleware(), Logout)
		v1.GET("/getuserinfo", AuthMiddleware(), GetUserInfo)

	}

}

// 注册
func Register(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})

		return
	}
	c.JSON(http.StatusOK, user)

	ok, msg := services.UserService{}.Register(user)
	if ok {
		c.JSON(http.StatusOK, gin.H{"message2": "User registered successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message3": msg})
	}

}

// 登录
func Login(c *gin.Context) {
	var user models.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message1": err.Error()})

		return
	}
	c.JSON(http.StatusOK, user)

	ok, msg := services.UserService{}.Login(user.Username, user.Password)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message2": "User login successfully",
			"token":    msg,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message3": msg})
	}

}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		// 解析token
		ok, user := services.UserService{}.ParseToken(token)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
		}
		// 将解析出的用户信息存入上下文，供后续处理使用

		c.JSON(200, gin.H{
			"userId":   user.Id,
			"username": user.Username,
			"email":    user.Email,
		})
		c.Set("userId", user.Id)
		c.Set("username", user.Username)
		c.Set("email", user.Email)
		c.Set("token", token)
		c.Next()
	}
}

// 登出
func Logout(c *gin.Context) {

	token := c.GetString("token")

	ok, err := services.UserService{}.Logout(token)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "logout ok",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

}

// 获得用户信息
func GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	username, _ := c.Get("username")
	mail, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"message":   "get userinfo successful",
		"uerid":     userId,
		"username":  username,
		"useremail": mail,
	})
}
