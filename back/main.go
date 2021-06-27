package main

import (
	"back/api"
	"back/docs"
	"back/ping"
	"back/portfolio"
	"back/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	DefaultAddr = "127.0.0.1"
	DefaultPort = "8080"
)

// @title Cryptio Front Exam API
// @version 1.0
// @description Hopefully this API will get you a job at Cryptio!

// @contact.name Lucas Santoni
// @contact.email lucas@cryptio.co

func main() {
	host := fmt.Sprintf(
		"%s:%s",
		utils.GetFromEnvOrDefault("ADDR", DefaultAddr),
		utils.GetFromEnvOrDefault("PORT", DefaultPort),
	)
	docs.SwaggerInfo.Host = host

	api := api.New()

	router := gin.Default()

	router.GET("/ping", ping.GET(api))

	portfolioRouter := router.Group("/portfolio")
	{
		portfolioRouter.GET("/history", portfolio.GETHistory(api))
		portfolioRouter.GET("/top-assets", portfolio.GETAssets(api))
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(host)
}
