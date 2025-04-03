package router

import (
	"exchange_app/controllers"

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
	return r
}
