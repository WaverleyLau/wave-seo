package user

import (
	"fmt"
	"net/http"
	"waverley/wave-seo/global"
	"waverley/wave-seo/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var login model.User

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("%v", login)
	if password, ok := global.USER_CACHE[login.Username]; ok {
		if password == login.Password {
			c.JSON(http.StatusOK, gin.H{"message": "login successful"})
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
	return

}

func Register(c *gin.Context) {

	if global.USER_CACHE == nil {
		global.USER_CACHE = make(map[string]string)
	}

	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	global.USER_CACHE[user.Username] = user.Password
	c.JSON(http.StatusOK, gin.H{"message": "register successful"})
	fmt.Sprintf("%v", global.USER_CACHE)
	return
}
