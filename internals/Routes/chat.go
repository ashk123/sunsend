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
	//json_obj, res1 := Base.ReadJSONMsg(c) // Read the JSON input from user
	//if res1 != 0 {
	//	response_file, _ := Data.NewResponse(res1, channel_id_user, nil, "")
	//	return c.JSON(
	//		response_file.Code,
	//		response_file,
	//	)
	//}
	//user := json_obj.Username
	//msg := json_obj.Message
	//reply := json_obj.Reply
	user := c.FormValue("username")
	msg := c.FormValue("message")
	reply := c.FormValue("reply")
	image_file_name := ""
	image, err1 := c.FormFile("image")
	if err1 == nil {
		fres := Base.HandleFiles(image)
		if fres != 0 {
			response_file, _ := Data.NewResponse(fres, channel_id_user, nil, "")
			return c.JSON(
				response_file.Code,
				response_file,
			)
		}
		image_file_name = image.Filename
	} else {
		log.Fatal(err1.Error())
	}
	//fmt.Println("A:LKSD:LAKSD:ALKSD:AKSD:A")
	//msg := c.FormValue("message") // get the message from userres_api_key
	// fmt.Println("user", user, "wants to send a message to channel", channel_id_user, ":", msg)
	log.Println(
		"[INFO] User",
		user,
		"wants to send a message to channel",
		channel_id_user,
		":",
		msg,
	)
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
		response, _ := Data.NewResponse(err, channel_id_user, nil, "")
		return c.JSON(response.Code, response) // should be same Statuscode as NewResponse
	}
	res_exists_channel := Base.ChannelExists(channel_id_user)
	if res_exists_channel != 0 {

		response_org, _ := Data.NewResponse(res_exists_channel, channel_id_user, nil, "")
		return c.JSON(
			response_org.Code,
			response_org,
		) // should be same Statuscode as NewResponse
	}

	// fmt.Println(msg)
	IChannel_id_ser, _ := strconv.Atoi(channel_id_user)
	Ireply, _ := strconv.Atoi(reply)
	crt := time.Now()
	// TODO: response msg base on date
	res := DB.InsertMsg(
		IChannel_id_ser,
		rand.Intn(999),
		user,
		msg,
		fmt.Sprintf("%d/%d/%d", crt.Year(), crt.Month(), crt.Day()),
		//crt,
		image,
		Ireply,
	)
	if res != 0 {
		response, _ := Data.NewResponse(res, channel_id_user, nil, "")
		return c.JSON(response.Code, response)
	}
	if image != nil {
		err := Base.SaveImageFile(image)
		if err != nil {
			response_error, _ := Data.NewResponse(15, channel_id_user, nil, "")
			return c.JSON(response_error.Code, response_error)
		}

	}
	response, _ := Data.NewResponse(http.StatusOK, channel_id_user, msg, image_file_name)
	return c.JSON(response.Code, response)

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
	flags := &Data.Flags{SetRangeMessage: []string{}} // initial of flags
	channel_id := c.Param("id")
	find_id := c.QueryParam("find")
	del_id := c.QueryParam("remove")
	user_name_search := c.QueryParam("username")
	var chat_collection []Data.Message
	get_message_range := c.QueryParam("range")
	var response *Data.Response
	res_exists_channel := Base.ChannelExists(channel_id) // if channel exists

	if res_exists_channel != 0 {
		response, _ = Data.NewResponse(res_exists_channel, channel_id, nil, "")
		return c.JSON(response.Code, response)
	}
	// Otherwise user wants to have the group of messages from a channel
	// fmt.Println(res_exists_channel)
	if get_message_range != "" { // If user wants to have specific amount of data
		log.Println("I'M UISING THJIS PROGRAM FOR THIS CHT KAHAHAHAHAHAH")
		data_spl := strings.Split(get_message_range, "-")
		flags.SetRangeMessage = data_spl
		// fmt.Println(flags.SetRangeMessage)
	}
	if user_name_search != "" {
		//log.Printf("a%sa", user_name_search)
		chat_collection_by_user, res := Base.FindMsgsByUsername(channel_id, user_name_search, flags)
		if res != 0 {
			response, _ = Data.NewResponse(res, channel_id, nil, "")
			return c.JSON(response.Code, response)
		}
		response, _ = Data.NewResponse(res, channel_id, chat_collection_by_user, "")
		return c.JSON(http.StatusOK, response)
	}

	if del_id != "" {
		//log.Printf("a%sa", user_name_search)
		res := Base.DeleteMsgsByID(channel_id, del_id)
		if res != 0 {
			response, _ = Data.NewResponse(res, channel_id, nil, "")
			return c.JSON(response.Code, response)
		}
		response, _ = Data.NewResponse(
			res,
			channel_id,
			fmt.Sprintf("%s succsessfully removed", del_id),
			"",
		)
		return c.JSON(http.StatusOK, response)
	}

	// If user just wants to find a message
	if find_id != "" {
		// if the data is digit
		chat_collection, res := Base.FindMsgByMsgID(channel_id, find_id, nil)
		if res != 0 {
			response, _ = Data.NewResponse(res, channel_id, nil, "")
			return c.JSON(response.Code, response)
		}
		response, _ = Data.NewResponse(res, channel_id, chat_collection, "")
		return c.JSON(response.Code, response)
	}

	chat_collection, res := Base.FindMsgsByChannelID(channel_id, flags)
	if res != 0 {
		response, _ = Data.NewResponse(11, channel_id, nil, "")
		return c.JSON(response.Code, response)
	}
	response, _ = Data.NewResponse(http.StatusOK, channel_id, chat_collection, "")
	//fmt.Println(len(chat_collection))

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
	chat_route_obj := NewRoute(fmt.Sprintf("/api/%s/chat/:id", Data.API_VERSION), chatActionFunc)
	return chat_route_obj
}

func GetChatPostRoute() *Route {
	chat_post_route_obj := NewRoute(
		fmt.Sprintf("/api/%s/chat/:id", Data.API_VERSION),
		GetChatPostAction,
	)
	return chat_post_route_obj
}
