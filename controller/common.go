package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrJson(msg string, args ...interface{}) {
	if len(args) == 0 {
		log.Panicln("缺少 *gin.Context")
	}
	var c *gin.Context
	if len(args) == 1 {
		ctx, ok := args[0].(*gin.Context)
		if !ok {
			log.Panicln("缺少 *gin.Context")
		}
		c = ctx
	}
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"errMsg":  msg,
		"data":    gin.H{},
	})
	// 终止请求链
	c.Abort()
}
