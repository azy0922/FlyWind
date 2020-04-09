package main

import (
	"io"
	"os"

	"github.com/azy0922/flywind/config"
	"github.com/azy0922/flywind/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	logFile, _ := os.Create("fwlog.log")
	gin.DefaultWriter = io.MultiWriter(logFile)
	app := gin.New()
	router.Route(app)
	app.Run(":" + config.String("port"))
}
