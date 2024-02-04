package cmd

import (
	"gostart/core"
	"gostart/views"
	"os"

	"github.com/labstack/echo/v4"
)

func RunServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return core.Render(c, views.App())
	})
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
