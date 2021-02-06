package main

import (
	// "encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handler "github.com/klan300/exceed17/handler"
	config "github.com/klan300/exceed17/config"
)

// return the value of the key
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  "${time_rfc3339}: FROM ${remote_ip} ${method} ${uri} RETURN ${status}\n",
		Skipper: middleware.DefaultSkipper,
	}))
	e.Use(middleware.Recover())

	e.GET("/:studentId", handler.GetDataById)
	e.GET("/twitter", handler.GetTweet)
	e.POST("/twitter", handler.PostTweet)
	e.PUT("/:studentId", handler.PutDataById)
	e.PATCH("/:studentId", handler.PatchDataById)
	

	e.Logger.Fatal(e.Start(config.GoDotEnvVariable("PORT")))
}