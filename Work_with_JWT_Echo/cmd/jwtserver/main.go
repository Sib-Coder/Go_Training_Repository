package main

import (
	"awesomeProject2/internal/app/endpoint"
	"awesomeProject2/internal/app/midleware"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/jwt", endpoint.GetJwt)

	//e.Use(midleware.ValidateJWT)
	e.GET("api", endpoint.Home, midleware.ValidateJWT)

	e.Logger.Fatal(e.Start(":3500"))

}
