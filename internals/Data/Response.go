package Data

import (
	"net/http"
)

/*
	Special Model for using response
*/

type Response struct {
	Res      string      `json:"RES"`
	Channel  string      `json:"CHANNEL,omitempty"`
	Ebody    string      `json:"ERROR_BODY,omitempty"`
	Messages interface{} `json:"DATA,omitempty"`
	Image    string      `json:"IMAGE,omitempty"`
	Code     int         `json:"CODE"`
}

func GenerateNewSuccsess(
	messages interface{},
	channel string,
	res string,
	code int,
	image string,
) *Response {
	return &Response{
		Res:      res,
		Channel:  channel,
		Image:    image,
		Code:     code,
		Messages: messages,
	}
}

func GenerateNewError(content string, status_code int, res string) *Response {
	return &Response{
		Ebody: content,
		Code:  status_code,
		Res:   res,
	}
}

// TODO: Add break for each case
func GetResponseByResult(
	res_code int,
) *Response { // TODO: make a struct for return values instead of return 3 values
	switch res_code {
	case 10:
		return GenerateNewError(
			"Your Message has a problem",
			http.StatusBadRequest,
			"FAILD",
		) // Response Status Error - Message
		// return &Error{
		// 	Content:    "Your message has a problem",
		// 	StatusCode: http.StatusBadRequest,
		// 	Res:        "FAILD",
		// }
	case 11:
		return GenerateNewError(
			"There is not any channel with this ID",
			http.StatusNotFound,
			"FAILD",
		) // Response Status Error - No Channel
		// return &Error{
		// 	Content:    "There is not any channel",
		// 	StatusCode: http.StatusBadRequest,
		// 	Res:        "FAILD",
		// }
	case 12:
		return GenerateNewError(
			"Your Message has a bad word",
			http.StatusNotAcceptable,
			"FAILD",
		) // response Status Error - Bad Word
		// return &Error{
		// 	Content:    "Your message has a bad word",
		// 	StatusCode: http.StatusNotAcceptable,
		// 	Res:        "FAILD",
		// }
	case 15:
		return GenerateNewError(
			"Your Messages Length is more than 30 character",
			http.StatusBadRequest,
			"FAILD",
		) // response Status Error - message length
		// return &Error{
		// 	Content:    "Your message length is more than 30 character",
		// 	StatusCode: http.StatusBadRequest,
		// 	Res:        "FAILD",
		// }
	case 16:
		return GenerateNewError(
			"There is a problem in the system",
			http.StatusServiceUnavailable,
			"FAILD",
		) // response Status Error - Server has problem (bug)
	case 17:
		return GenerateNewError(
			"Invalid API KEY",
			http.StatusBadRequest,
			"FAILD",
		) // response Status Error - invalid api key
	case 18:
		return GenerateNewError(
			"You Reached the Limit request time",
			http.StatusForbidden,
			"FAILD",
		) // respponse Status error - Reached request limit
	case 19:
		return GenerateNewError(
			"There is not any message with that ID",
			http.StatusNotFound,
			"FAILD",
		)
	case 21:
		return GenerateNewError(
			"There is not any Message with that username",
			http.StatusNotFound,
			"FAILD",
		) // response Status error - not found messages by username
	case 23:
		return GenerateNewError(
			"Your file name is too long!",
			http.StatusNotAcceptable,
			"FAILD",
		)
	case 24:
		return GenerateNewError(
			"You only allow to upload [png, jpg, bmp] formats!",
			http.StatusAccepted,
			"FAILD",
		)
	case 25:
		return GenerateNewError(
			"You reached the file limit, limit is 30MB",
			http.StatusAccepted,
			"FAILD",
		)
	case 27:
		return GenerateNewError(
			"There is not any imag file with that name",
			http.StatusNotFound,
			"FAILD",
		)
	case 30:
		return GenerateNewError(
			"Error when fetching channel list",
			http.StatusNotFound,
			"FAILD",
		)
	case 31:
		return GenerateNewError(
			"Entry Json is not correct",
			400,
			"FAILD",
		)

	case 33:
		return GenerateNewError(
			"Faild to delete the message",
			http.StatusNotAcceptable,
			"FAILD",
		)
	default:
		return GenerateNewError(
			"",
			0,
			"",
		)
	}
}

// TODO: Do the GenerateNewSuccess with GetResponseByResult
func NewResponse(
	res_code int,
	channel string,
	messages interface{},
	image string,
) (*Response, error) {
	error_obj := GetResponseByResult(res_code)
	if error_obj.Code != 0 {
		return error_obj, nil
	}
	return GenerateNewSuccsess(messages, channel, "SUCCSESS", res_code, image), nil
}
