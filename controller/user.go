package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const adminName = "admin"

var Store = sessions.NewCookieStore([]byte("qianjun-flywind-js"))

func CheckLogin(c *gin.Context) {
	session, _ := Store.Get(c.Request, "session")
	val, ok := session.Values["username"]
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"errMsg":  "",
			"data":    gin.H{"username": val.(string)},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"errMsg":  "",
			"data":    "",
		})
	}
}

func Login(c *gin.Context) {
	if c.PostForm("username") == adminName {
		//设置session
		session, _ := Store.Get(c.Request, "session")
		session.Values["username"] = adminName
		session.Save(c.Request, c.Writer)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"errMsg":  "",
			"data":    gin.H{"username": adminName},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"errMsg":  "",
			"data":    "",
		})
	}
}

func Logout(c *gin.Context) {
	session, _ := Store.Get(c.Request, "session")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"errMsg":  "",
		"data":    "",
	})
}
