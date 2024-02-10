package Routes

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sunsend/internals/Base"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
	"time"

	"github.com/labstack/echo/v4"
)

func GetChatPostAction(c echo.Context) error {
	channel_id_user := c.Param("id")
	// flags := c.Request().Header

	user := c.FormValue("user")
	msg := c.FormValue("message") // get the message from userres_api_key
	// fmt.Println("user", user, "wants to send a message to channel", channel_id_user, ":", msg)
	log.Println("user", user, "wants to send a message to channel", channel_id_user, ":", msg)
	// headers := c.Request().Header
	// apiKey, res_api_key := Base.BearerToken(headers)
	// if res_api_key != 0 {
	// 	response, _ := Data.NewResponse(c, res_api_key, channel_id_user, nil)
	// 	_, error_code_org, _ := Data.GetErrorByResult(res_api_key)
	// 	return c.JSON(error_code_org, response)
	// }
	// res_check_api := Base.ApiKeyIsValid(apiKey)
	// if res_check_api != 0 {
	// 	response2, _ := Data.NewResponse(c, res_check_api, channel_id_user, nil)
	// 	_, error_code_org2, _ := Data.GetErrorByResult(res_check_api)
	// 	return c.JSON(error_code_org2, response2)
	// }
	err := Base.CheckMessage(msg)
	if err != 0 {
		response, _ := Data.NewResponse(c, err, channel_id_user, nil)
		error_obj := Data.GetErrorByResult(err)
		return c.JSON(error_obj.StatusCode, response) // should be same Statuscode as NewResponse
	}
	res_exists_channel := Base.ChannelExists(channel_id_user)
	if res_exists_channel != 0 {
		response_org, _ := Data.NewResponse(c, res_exists_channel, channel_id_user, nil)
		error_obj2 := Data.GetErrorByResult(res_exists_channel)
		return c.JSON(error_obj2.StatusCode, response_org) // should be same Statuscode as NewResponse
	}

	// fmt.Println(msg)
	IChannel_id_ser, _ := strconv.Atoi(channel_id_user)
	res := DB.InsertMsg(IChannel_id_ser, rand.Intn(999), user, msg, time.Now().String(), 0)
	if res != 0 {
		response, _ := Data.NewResponse(c, res, channel_id_user, nil)
		error_obj3 := Data.GetErrorByResult(res_exists_channel)
		return c.JSON(error_obj3.StatusCode, response)
	}
	response, _ := Data.NewResponse(c, 0, channel_id_user, msg)
	error_obj4 := Data.GetErrorByResult(res_exists_channel)
	return c.JSON(error_obj4.StatusCode, response)

}

func StreamResponseJSON(c echo.Context, chat_data *Data.Response) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(chat_data)
}

// This is the Get Function Runner
// TODO: It's better to *Stream* the JSON message file
func chatActionFunc(c echo.Context) error {
	// Handle the Chat Channel with c.Param("id")
	flags := Data.Flags{SetRangeMessage: []string{}} // initial of flags
	channel_id := c.Param("id")
	var chat_collection []*Data.Message
	get_message_range := c.QueryParam("range")
	var response *Data.Response
	res_exists_channel := Base.ChannelExists(channel_id) // if channel exists

	if res_exists_channel != 0 {
		response, _ = Data.NewResponse(c, res_exists_channel, channel_id, nil)
		error_obj := Data.GetErrorByResult(res_exists_channel)
		return c.JSON(error_obj.StatusCode, response)
	}
	fmt.Println(res_exists_channel)
	if get_message_range != "" { // If user wants to have specific amount of data
		data_spl := strings.Split(get_message_range, "-")
		flags.SetRangeMessage = data_spl
		// fmt.Println(flags.SetRangeMessage)
	}

	chat_collection, res := Base.FindMsgsByChannelID(channel_id, flags)
	if res != 0 {
		response, _ = Data.NewResponse(c, 11, channel_id, nil)
		error_obj2 := Data.GetErrorByResult(res)
		return c.JSON(error_obj2.StatusCode, response)
	}
	response, _ = Data.NewResponse(c, res, channel_id, chat_collection)
	fmt.Println(len(chat_collection))

	// headers := c.Request().Header
	// apiKey, res_api_key := Base.BearerToken(headers)
	// if res_api_key != 0 {
	// 	response, _ = Data.NewResponse(c, res_api_key, channel_id, nil)
	// 	_, error_code_org, _ := Data.GetErrorByResult(res_api_key)
	// 	return c.JSON(error_code_org, response)
	// }
	// res_check_api := Base.ApiKeyIsValid(apiKey)
	// if res_check_api != 0 {
	// 	response, _ = Data.NewResponse(c, res_check_api, channel_id, nil)
	// 	_, error_code_org, _ := Data.GetErrorByResult(res_check_api)
	// 	return c.JSON(error_code_org, response)
	// }
	// fmt.Println("API KEY:", apiKey, "requested to server succsessfully")
	// response, _ = Data.NewResponse(c, res, channel_id, chat_collection)
	// return c.JSON(http.StatusOK, response)
	return StreamResponseJSON(c, response) // Stream JSON File
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
