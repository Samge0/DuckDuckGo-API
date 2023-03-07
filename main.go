package main

import (
	"github.com/acheong08/DuckDuckGo-API/app/config"
	"github.com/acheong08/DuckDuckGo-API/app/controllers"
	"github.com/acheong08/DuckDuckGo-API/app/middlewares"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var ddgController = controllers.NewDDGController()

func main() {
	HOST := "0.0.0.0"
	PORT := config.LoadConfig().Port
	handler := gin.Default()
	handler.Use(middlewares.Cors())
	apiRouter := handler.Group("search").Use(middlewares.TokenJWTAuth())
	{
		apiRouter.POST("/duck", ddgController.HandlerDDGSearchPost)
		apiRouter.GET("/duck", ddgController.HandlerDDGSearchGet)
	}
	endless.ListenAndServe(HOST+":"+PORT, handler)
}
