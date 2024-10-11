package main

import (
	"context"
	"tmplexample/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "andres@maksdf")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
