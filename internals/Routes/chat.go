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
	if err == 0 {
		// fmt.Println(msg)
		IChannel_id_ser, _ := strconv.Atoi(channel_id_user)
		res := DB.InsertMsg(IChannel_id_ser, rand.Intn(999), user, msg, "2023", 0)
		if res != 0 {
			response, _ := Data.NewResponse(c, res, channel_id_user, nil)
			return c.JSON(http.StatusBadRequest, response)
		}
		response, _ := Data.NewResponse(c, 0, channel_id_user, msg)
		return c.JSON(http.StatusOK, response)
	} else {
		response, _ := Data.NewResponse(c, err, channel_id_user, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

}

func chatActionFunc(c echo.Context) error {
	// Handle the Chat Channel with c.Param("id")
	channel_id := c.Param("id")
	var response *Data.Response
	// Get the latest Messages from that channel
	chat_collection := Base.FindMsgsByChannelID(channel_id)
	if chat_collection == nil {
		response, _ = Data.NewResponse(c, 11, channel_id, nil)
	} else {
		// DB.InsertMsg(123, 847, "eiko", "this is my first message", "2023", 0)
		response, _ = Data.NewResponse(c, 0, channel_id, chat_collection)
	}
	// Render the Message with chat.html template

	// return c.Render(http.StatusOK, "chat.html", map[string]interface{}{
	// 	"msgs": chat_collectino,
	// })
	// })
	return c.JSON(http.StatusOK, response)

}

func GetChatRoute() *Route {
	chat_route_obj := NewRoute("/api/v1/chat/:id", chatActionFunc)
	return chat_route_obj
}

func GetChatPostRoute() *Route {
	chat_post_route_obj := NewRoute("/api/v1/chat/:id", GetChatPostAction)
	return chat_post_route_obj
}
