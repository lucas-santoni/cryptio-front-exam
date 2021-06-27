package main

import (
	"back/api"
	"back/ping"
	"back/portfolio"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	ADDR = "127.0.0.1"
	PORT = 8080
)

func main() {
	api := api.New()

	router := gin.Default()

	router.GET("/ping", ping.GET(api))

	portfolioRouter := router.Group("portfolio")
	{
		portfolioRouter.GET("/history", portfolio.GETHistory(api))
		portfolioRouter.GET("/top-assets", portfolio.GETAssets(api))
	}

	router.Run(fmt.Sprintf("%s:%d", ADDR, PORT))
}
