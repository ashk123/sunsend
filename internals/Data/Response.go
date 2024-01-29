package Data

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Res      string      `json:"RES"`
	Channel  string      `json:"CHANNEL"`
	Ebody    string      `json:"ERROR_BODY,omitempty"`
	Messages interface{} `json:"DATA,omitempty"`
	Code     int         `json:"CODE"`
}

func GetErrorByResult(res_code int) (string, int, string) {
	switch res_code {
	case 0:
		return "", http.StatusOK, "SUCCSESS" // Response Status Ok
	case 10:
		return "Your Message has a problem", http.StatusBadRequest, "FAILD" // Response Status Error - Message
	case 11:
		return "There is not any channel", http.StatusBadRequest, "FAILD" // Response Status Error - No Channel
	case 12:
		return "Your Message has a bad word", http.StatusNotAcceptable, "FAILD" // response Status Error - Bad Word
	case 15:
		return "Your Messages Length is more than 30 character", http.StatusBadRequest, "FAILD" // response Status Error - message length
	default:
		return "There is a problem", http.StatusNotAcceptable, "FAILD" // Response Status Error - uknown Eror
	}
}

func NewResponse(c echo.Context, res_code int, channel string, messages interface{}) (*Response, error) {
	ebody, error_code, status_str := GetErrorByResult(res_code)
	response := &Response{
		Res:      status_str,
		Channel:  channel,
		Ebody:    ebody,
		Code:     error_code,
		Messages: messages,
	}
	// if err := c.Bind(response); err != nil {
	// 	return nil, errors.New("Can't bind the response to json")
	// }
	return response, nil
}
