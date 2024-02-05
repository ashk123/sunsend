package main

import (
	"fmt"
	"log"
	"sunsend/internals/CoreConfig"
	"sunsend/internals/DB"
	"sunsend/internals/Handlers"
	"sunsend/internals/Renderer"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	CoreConfig.UpdateConfigs() // load bouth .env configs and user configs
	log.Println("Your API Key is: " + fmt.Sprintf("%s", CoreConfig.Configs.Server.Key))
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Renderer = &Renderer.Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	Handlers.Handler(e)
	DB.PrepairDBSystem()

	e.Logger.Fatal(e.Start(":" + CoreConfig.Configs.Dotenv["PORT"]))
}
