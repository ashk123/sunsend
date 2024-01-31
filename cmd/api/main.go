package main

import (
	"sunsend/internals/DB"
	"sunsend/internals/Handlers"
	"sunsend/internals/Renderer"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = &Renderer.Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	Handlers.Handler(e)
	DB.PrepairDBSystem()

	e.Logger.Fatal(e.Start(":3000"))
}
