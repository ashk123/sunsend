package Routes

import (
	"github.com/labstack/echo/v4"
)

type Route struct {
	Path   string           // Route Path
	Runner echo.HandlerFunc // Route Action Function
}

func NewRoute(user_path string, user_runner echo.HandlerFunc) *Route {
	return &Route{
		Path:   user_path,
		Runner: user_runner,
	}
}

// unused function (maybe later)
// func (m *Route) getRunner() echo.HandlerFunc {
// 	return m.Runner
// }
