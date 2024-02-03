package Base

import (
	"sunsend/internals/DB"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func GetChannelByID(c echo.Context, channel string) (*Data.Channel, int) {
	row, res := DB.QueryRow("SELECT * FROM Channels WHERE ID == '" + channel + "'")
	if res != 0 {
		return nil, res
	}
	var res_id int
	var res_name, res_des, res_owner string
	err := row.Scan(&res_id, &res_name, &res_des, &res_owner) // Fetch data from query's result
	if err != nil {
		// log.Fatal(err.Error())
		return nil, 11
	}
	ret_channel := &Data.Channel{
		ID:          res_id,
		Name:        res_name,
		Description: res_des,
		Owner:       res_owner,
	}
	return ret_channel, 0
}
