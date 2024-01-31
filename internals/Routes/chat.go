package Routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sunsend/internals/Base"
	"sunsend/internals/DB"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func GetChatPostAction(c echo.Context) error {
	channel_id_user := c.Param("id")
	user := c.FormValue("user")
	msg := c.FormValue("message") // get the message from user
	fmt.Println("user", user, "wants to send a message to channel", channel_id_user, ":", msg)
	err := Base.CheckMessage(msg)
	if err != 0 {
		response, _ := Data.NewResponse(c, err, channel_id_user, nil)
		return c.JSON(http.StatusBadRequest, response) // should be same Statuscode as NewResponse
	}
	res_exists_channel := Base.ChannelExists(channel_id_user)
	if res_exists_channel != 0 {
		response, _ := Data.NewResponse(c, res_exists_channel, channel_id_user, nil)
		return c.JSON(http.StatusBadRequest, response) // should be same Statuscode as NewResponse
	}

	// fmt.Println(msg)
	IChannel_id_ser, _ := strconv.Atoi(channel_id_user)
	res := DB.InsertMsg(IChannel_id_ser, rand.Intn(999), user, msg, "2023", 0)
	if res != 0 {
		response, _ := Data.NewResponse(c, res, channel_id_user, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	response, _ := Data.NewResponse(c, 0, channel_id_user, msg)
	return c.JSON(http.StatusOK, response)

}

// This is the Get Function Runner
func chatActionFunc(c echo.Context) error {
	// Handle the Chat Channel with c.Param("id")
	channel_id := c.Param("id")
	var response *Data.Response
	res := Base.ChannelExists(channel_id)
	if res != 0 {
		fmt.Println(res)
		response, _ = Data.NewResponse(c, res, channel_id, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	chat_collection, res := Base.FindMsgsByChannelID(channel_id)
	if res != 0 {
		response, _ = Data.NewResponse(c, 11, channel_id, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	fmt.Println(len(chat_collection))

	response, _ = Data.NewResponse(c, res, channel_id, chat_collection)
	return c.JSON(http.StatusOK, response)
	// return c.JSON(http.StatusAccepted, map[string]interface{}{
	// 	"asd": "nice",
	// })
}
func GetChatRoute() *Route {
	chat_route_obj := NewRoute("/api/v1/chat/:id", chatActionFunc)
	return chat_route_obj
}

func GetChatPostRoute() *Route {
	chat_post_route_obj := NewRoute("/api/v1/chat/:id", GetChatPostAction)
	return chat_post_route_obj
}
