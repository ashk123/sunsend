package main

import (
	"fmt"
	"sunsend/internals/Base"
	"sunsend/internals/CoreConfig"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
	"sunsend/internals/Handlers"
	"sunsend/internals/Renderer"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	temp := Base.GetTemp()
	CoreConfig.UpdateConfigs() // load bouth .env configs and user configs
	Data.LoadWordsFromConfig() // It loads here just for test
	CoreConfig.ShowConfigInformation()
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		for {
			select {
			case value := <-ticker.C:
				fmt.Println(value)
				//fmt.Println("this is the Ticker that I had")
				Base.ControlUnit(value)
			}
		}
	}()
	// fmt.Println(Base.LimitCheck(nil))
	// log.Println("Your API Key is: " + fmt.Sprintf("%s", CoreConfig.Configs.Server.Key))
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Renderer = &Renderer.Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	Handlers.Handler(e)
	DB.PrepairDBSystem()

	e.Logger.Fatal(e.Start("127.0.0.1:" + temp.Get("config").(*Data.Config).Dotenv["PORT"]))
}
