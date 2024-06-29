package Routes

import (
	"sunsend/internals/Base"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func channelRouteAction(c echo.Context) error {
	channel_id := c.Param("id")
	channel, res := Base.GetChannelByID(c, channel_id)
	if res != 0 {
		response, _ := Data.NewResponse(res, channel_id, nil, "")
		return c.JSON(response.Code, response)
	}
	response_good, _ := Data.NewResponse(res, channel_id, channel, "")
	return c.JSON(response_good.Code, response_good)
	// return c.JSON(http.StatusOK, response)
}

func GetChannelRoute() *Route {
	channel_route_obj := NewRoute("/api/v1/channel/:id", channelRouteAction)
	return channel_route_obj
}
