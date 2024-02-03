package main

import (
	"log"
	"os"
	"sunsend/internals/DB"
	"sunsend/internals/Handlers"
	"sunsend/internals/Renderer"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Renderer = &Renderer.Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	Handlers.Handler(e)
	DB.PrepairDBSystem()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
