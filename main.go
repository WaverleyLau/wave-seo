package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run(":9890")
}
