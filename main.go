package main

import (
	"fmt"
	"net/http"
	"waverley/wave-seo/app"
	"waverley/wave-seo/global"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
func main() {
	app.InitApp()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run(fmt.Sprintf(":%s", global.WAVE_CONFIG.Server.Port)) // listen and serve on 0.0.0.0:8080
}
