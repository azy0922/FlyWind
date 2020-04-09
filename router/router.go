package router

import (
	"net/http"

	"github.com/azy0922/flywind/controller"
	"github.com/azy0922/flywind/middleware"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFile("/", "./web/index.html")
	router.StaticFS("/static", http.Dir("./web/static"))
	router.StaticFS("/js", http.Dir("./web/js"))
	router.StaticFS("/css", http.Dir("./web/css"))

	router.GET("/math", controller.ShowMaths)
	router.GET("/test", controller.TestMaths)
	router.GET("/term", controller.ShowTerms)

	api := router.Group("/api")
	{
		api.GET("/auth/checkLogin", controller.CheckLogin)
		api.POST("/auth/login", controller.Login)
		api.DELETE("/auth/logout", controller.Logout)

		api.GET("/msg", controller.ListProblems)
		api.GET("/msg/:id", middleware.SigninRequired, controller.ListProblem)
		api.GET("/authors", controller.ListAuthors)
		api.POST("/msg", middleware.SigninRequired, controller.AddProblem)
		api.PUT("/msg/:id", middleware.SigninRequired, controller.UpdateProblem)
		api.DELETE("/msg/:id", middleware.SigninRequired, controller.DeleteProblem)

		api.GET("/term", controller.ListTerms)
		api.GET("/term/:id", middleware.SigninRequired, controller.ListTerm)
		api.GET("/tauthors", controller.ListTAuthors)
		api.POST("/term", middleware.SigninRequired, controller.AddTerm)
		api.PUT("/term/:id", middleware.SigninRequired, controller.UpdateTerm)
		api.DELETE("/term/:id", middleware.SigninRequired, controller.DeleteTerm)
	}
}
