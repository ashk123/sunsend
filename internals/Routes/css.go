package Routes

import "github.com/labstack/echo/v4"

func RenderResources(c echo.Context) error {
	return c.File("templates/css/" + c.Param("file"))
}

func GetCssResources() *Route {
	return &Route{
		Path:   "/css/:file",
		Runner: RenderResources,
	}
}
