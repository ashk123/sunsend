package Routes

import (
	"sunsend/internals/Base"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func listRouteaction(c echo.Context) error {
	data, res := Base.GetChannelList()
	if res != 0 {
		response, _ := Data.NewResponse(res, "", nil, "")
		return c.JSON(response.Code, response)
	}
	response, _ := Data.NewResponse(200, "", data, "")
	return c.JSON(response.Code, response)
}

func GetListRoute() *Route {
	return NewRoute("/api/v1/list", listRouteaction)
}
