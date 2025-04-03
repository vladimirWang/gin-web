package main

import (
	"exchange_app/config"
	"exchange_app/router"
)

func main() {
	config.InitConfig()
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r := router.SetupRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	r.Run(port) // listen and serve on 0.0.0.0:8080
}
