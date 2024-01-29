package Routes

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	Path   string
	Runner echo.HandlerFunc
}

func NewRoute(user_path string, user_runner echo.HandlerFunc) *Route {
	return &Route{
		Path:   user_path,
		Runner: user_runner,
	}
}

func (m *Route) getRunner() echo.HandlerFunc {
	return m.Runner
}
