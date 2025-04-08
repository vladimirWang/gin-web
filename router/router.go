package router

import (
	"exchange_app/controllers"
	"exchange_app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		// auth.POST("/login", func(ctx *gin.Context) {
		// 	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		// 		"msg": "login",
		// 	})
		// })
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := r.Group("/api/")

	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())

	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.CreateArticle)
	}
	return r
}
