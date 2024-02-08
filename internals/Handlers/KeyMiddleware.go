package Handlers

import (
	"fmt"
	"sunsend/internals/Base"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

// API KEY Authentication Middleware

func CheckAPIKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			headers := c.Request().Header
			channel_id := c.Param("Id")
			var response *Data.Response
			apiKey, res_api_key := Base.BearerToken(headers)
			if res_api_key != 0 {
				response, _ = Data.NewResponse(c, res_api_key, channel_id, nil)
				_, error_code_org, _ := Data.GetErrorByResult(res_api_key)
				return c.JSON(error_code_org, response)
			}
			res_check_api := Base.ApiKeyIsValid(apiKey)
			if res_check_api != 0 {
				response, _ = Data.NewResponse(c, res_check_api, channel_id, nil)
				_, error_code_org, _ := Data.GetErrorByResult(res_check_api)
				return c.JSON(error_code_org, response)
			}
			fmt.Println("API KEY:", apiKey, "requested to server succsessfully")
			return next(c)
		})
	}
}
