package Routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"sunsend/internals/Base"
	"sunsend/internals/DB"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func channelPostAction(c echo.Context) error {
	data, res1 := Base.ReadJSONChan(c) // TODO: Return error number code instead of err
	if res1 != 0 {
		response, _ := Data.NewResponse(res1, "", nil, "")
		return c.JSON(response.Code, response)
	}
	channel_id := rand.Intn(999)

	res := DB.InsertChannel(
		channel_id,
		data.Name,
		data.Description,
		data.Owner,
	) // Create default chat Channel
	if res != 0 {
		response, _ := Data.NewResponse(res, "", nil, "")
		return c.JSON(response.Code, response)
	}
	response, _ := Data.NewResponse(http.StatusOK, Base.Itsr(channel_id), nil, "")
	return c.JSON(response.Code, response)
}

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
	channel_route_obj := NewRoute(
		fmt.Sprintf("/api/%s/channel/:id", Data.API_VERSION),
		channelRouteAction,
	)
	return channel_route_obj
}

func GetChannelPostRoute() *Route {
	channel_post_route_obj := NewRoute(
		fmt.Sprintf("/api/%s/channel/create", Data.API_VERSION), channelPostAction)
	return channel_post_route_obj
}
