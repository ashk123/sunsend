package Routes

import (
	"fmt"
	"net/http"
	"sunsend/internals/Base"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func media_route_action(c echo.Context) error {
	media_file_name := c.Param("file_name")
	// TODO: Check for SqlInjection things
	//log.Println("File name is", media_file_name+".zst")
	if Base.IsFileExists("Storage/"+media_file_name+".zst") == false {
		response, _ := Data.NewResponse(27, "", nil, "")
		return c.JSON(response.Code, response)
	}
	data, cerr := Base.DecompressFile(media_file_name)
	if cerr != 0 {
		response1, _ := Data.NewResponse(cerr, "", nil, "")
		return c.JSON(response1.Code, response1)
	}
	// TODO: Check the format type and return the correct format instead of  indexing
	c.Blob(http.StatusOK, "image/"+media_file_name[len(media_file_name)-3:], data)
	return nil
}

func GetMediaRoute() *Route {
	temp := Data.GetTemp()
	media_route_obj := NewRoute(
		fmt.Sprintf(
			"/api/%s/%s/:file_name",
			temp.Get("config").(*Data.Config).Dotenv["ADR"],
			temp.Get("config").(*Data.Config).Uconfig.MediaRoute,
		),
		media_route_action,
	)
	return media_route_obj
}
