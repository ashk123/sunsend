package Routes

import (
	"fmt"
	"net/http"
	"sunsend/internals/Base"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func GetChatPostAction(c echo.Context) error {
	channel_id_user := c.Param("id")
	// user := c.FormValue("user")
	msg := c.FormValue("message") // get the message from user
	err := Base.CheckMessage(msg)
	if err == 0 {
		fmt.Println(msg)
		// u := &User{
		// 	Res:   "welcome",
		// 	Email: "welcome2 for email",
		// }
		// if err := c.Bind(u); err != nil {
		// 	return err
		// }
		// response := &Response{
		// 	Res:     "SUCCSESS",
		// 	Channel: channel_id_user,
		// 	Ebody:   "",
		// 	Code:    http.StatusOK,
		// }
		// if err := c.Bind(response); err != nil {
		// 	return err
		// }
		response, _ := Data.NewResponse(c, 0, channel_id_user, msg)
		return c.JSON(http.StatusOK, response)
		// return c.JSON(http.StatusOK, u)
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
