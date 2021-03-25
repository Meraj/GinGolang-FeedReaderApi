package main

import (
	"./app"
	"./app/Controllers"
	"./app/Middleware"
	"github.com/gin-gonic/gin"
	"./app/Classes"
	"time"
)

func main() {
	app.Database().Close()
	Route := gin.Default()
	AuthRoutes := Route.Group("auth")
	{
		var AuthController Controllers.AuthController
		AuthRoutes.POST("login",AuthController.Login)
		AuthRoutes.POST("register",AuthController.Register)
		AuthRoutes.POST("forgot-password",AuthController.ForgotPassword)
		AuthRoutes.GET("verify/:token",AuthController.VerifyAccount)
	}
	DashboardRoutes := Route.Group("dashboard")
	DashboardRoutes.Use(Middleware.Auth())
	{
		var dashboard Controllers.DashboardController
		DashboardRoutes.GET("/",dashboard.Index)
	}
	FeedRoutes := Route.Group("feed")
	{
		var FeedController Controllers.FeedController
		FeedRoutes.GET("all",FeedController.All)
		FeedRoutes.GET("source/:source")
	}

	go updater()

	Route.Run() // listen and serve on 0.0.0.0:8080

}

func updater() {
	for {
		time.Sleep(60 * time.Second)
		Classes.GetSources()
	}
}