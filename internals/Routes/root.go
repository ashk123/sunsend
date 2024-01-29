package Routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func rootAction(c echo.Context) error {
	return c.Render(http.StatusOK, "main.html", nil)
}

func GetRootRoute() *Route {
	root_route_obj := NewRoute("/", rootAction)
	return root_route_obj
}
