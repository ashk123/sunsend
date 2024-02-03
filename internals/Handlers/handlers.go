package Handlers

import (
	"sunsend/internals/Routes"

	"github.com/labstack/echo/v4"
)

func Handler(e *echo.Echo) {

	e.GET(Routes.GetRootRoute().Path, Routes.GetRootRoute().Runner)          // set the Root Router ('/')
	e.GET(Routes.GetCssResources().Path, Routes.GetCssResources().Runner)    // set the Resources Router ("/css/:file")
	e.GET(Routes.GetChatRoute().Path, Routes.GetChatRoute().Runner)          // set the chat Router ("/chat/:id")
	e.POST(Routes.GetChatPostRoute().Path, Routes.GetChatPostRoute().Runner) // set the chat Router ("/chat/:id")
	e.GET(Routes.GetChannelRoute().Path, Routes.GetChannelRoute().Runner)
}
