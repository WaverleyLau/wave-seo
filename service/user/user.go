package user

import (
	"fmt"
	"net/http"
	"waverley/wave-seo/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	trueUsername := "waverley"
	truePassword := "waverley"

	var login model.User

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Printf("%v", login)
	if login.Username != trueUsername || login.Password != truePassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "login successful"})
	}

}
