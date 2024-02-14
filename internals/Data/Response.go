package Data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
	Special Model for using response
*/

type Response struct {
	Res      string      `json:"RES"`
	Channel  string      `json:"CHANNEL"`
	Ebody    string      `json:"ERROR_BODY,omitempty"`
	Messages interface{} `json:"DATA,omitempty"`
	Code     int         `json:"CODE"`
}

// A simple error model to simplify the GetErrorByResult function
type Error struct {
	Content    string
	StatusCode int
	Res        string
}

func GenerateNewError(content string, status_code int, res string) *Error {
	return &Error{
		Content:    content,
		StatusCode: status_code,
		Res:        res,
	}
}

func GetErrorByResult(res_code int) *Error { // TODO: make a struct for return values instead of return 3 values
	switch res_code {
	case 0:
		return GenerateNewError("", http.StatusOK, "SUCCSESS") // Response Status Ok
	case 10:
		return GenerateNewError("Your Message has a problem", http.StatusBadRequest, "FAILD") // Response Status Error - Message
		// return &Error{
		// 	Content:    "Your message has a problem",
		// 	StatusCode: http.StatusBadRequest,
		// 	Res:        "FAILD",
		// }
	case 11:
		return GenerateNewError("There is not any channel with this ID", http.StatusNotFound, "FAILD") // Response Status Error - No Channel
		// return &Error{
		// 	Content:    "There is not any channel",
		// 	StatusCode: http.StatusBadRequest,
		// 	Res:        "FAILD",
		// }
	case 12:
		return GenerateNewError("Your Message has a bad word", http.StatusNotAcceptable, "FAILD") // response Status Error - Bad Word
		// return &Error{
		// 	Content:    "Your message has a bad word",
		// 	StatusCode: http.StatusNotAcceptable,
		// 	Res:        "FAILD",
		// }
	case 15:
		return GenerateNewError("Your Messages Length is more than 30 character", http.StatusBadRequest, "FAILD") // response Status Error - message length
		// return &Error{
		// 	Content:    "Your message length is more than 30 character",
		// 	StatusCode: http.StatusBadRequest,
		// 	Res:        "FAILD",
		// }
	case 16:
		return GenerateNewError("There is a problem in the system", http.StatusServiceUnavailable, "FAILD") // response Status Error - Server has problem (bug)
	case 17:
		return GenerateNewError("Invalid API KEY", http.StatusBadRequest, "FAILD") // response Status Error - invalid api key
	case 18:
		return GenerateNewError("You Reached the Limit request time", http.StatusForbidden, "FAILD") // respponse Status error - Reached request limit
	case 19:
		return GenerateNewError("There is not any message with that ID", http.StatusNotFound, "FAILD")
	default:
		return GenerateNewError("There is a problem", http.StatusNotAcceptable, "FAILD") // Response Status Error - uknown Eror
	}
}

func NewResponse(c echo.Context, res_code int, channel string, messages interface{}) (*Response, error) {
	error_obj := GetErrorByResult(res_code)
	response := &Response{
		Res:      error_obj.Res,
		Channel:  channel,
		Ebody:    error_obj.Content,
		Code:     error_obj.StatusCode,
		Messages: messages,
	}
	// if err := c.Bind(response); err != nil {
	// 	return nil, errors.New("Can't bind the response to json")
	// }
	return response, nil
}
