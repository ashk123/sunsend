package Handlers

import (
	"sunsend/internals/Routes"

	"github.com/didip/tollbooth_echo"
	"github.com/labstack/echo/v4"
)

// doc: https://blog.logrocket.com/rate-limiting-go-application/

func Handler(e *echo.Echo) {
	limiter := GetLimiterMiddleWare() // set the limiter middleware
	e.GET(
		Routes.GetRootRoute().Path,
		Routes.GetRootRoute().Runner,
	) // set the Root Router ('/')
	e.GET(
		Routes.GetCssResources().Path,
		Routes.GetCssResources().Runner,
	) // set the Resources Router ("/css/:file")
	e.GET(
		Routes.GetChatRoute().Path,
		Routes.GetChatRoute().Runner,
		tollbooth_echo.LimitHandler(limiter),
		CheckAPIKey(),
	) // set the chat Router ("/chat/:id")
	e.POST(
		Routes.GetChatPostRoute().Path,
		Routes.GetChatPostRoute().Runner,
		tollbooth_echo.LimitHandler(limiter),
		CheckAPIKey(),
	) // set the chat Router ("/chat/:id")
	e.GET(
		Routes.GetChannelRoute().Path,
		Routes.GetChannelRoute().Runner,
		tollbooth_echo.LimitHandler(limiter),
		CheckAPIKey(),
	) // set the channel Router ("/channel/:id")
	e.POST(
		Routes.GetChannelPostRoute().Path,
		Routes.GetChannelPostRoute().Runner,
		tollbooth_echo.LimitHandler(limiter),
		CheckAPIKey(),
	) // set the channel Router ("/channel/:id")
	e.GET(
		Routes.GetListRoute().Path,
		Routes.GetListRoute().Runner,
		tollbooth_echo.LimitHandler(limiter),
		CheckAPIKey(),
	) // set the channel Router ("/channel/:id")
	e.GET(
		Routes.GetMediaRoute().Path,
		Routes.GetMediaRoute().Runner,
		tollbooth_echo.LimitHandler(limiter),
		CheckAPIKey(),
	)
}
