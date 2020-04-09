package middleware

import (
	"github.com/azy0922/flywind/controller"
	"github.com/gin-gonic/gin"
)

func SigninRequired(c *gin.Context) {
	session, _ := controller.Store.Get(c.Request, "session")
	_, ok := session.Values["username"]
	if !ok {
		controller.SendErrJson("未登录", c)
		return
	}
	c.Next()
}
