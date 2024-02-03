package Base

import (
	"fmt"
	"sunsend/internals/DB"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
)

func GetChannelByID(c echo.Context, channel string) (*Data.Channel, int) {
	row := DB.QueryRow("SELECT * FROM Channels WHERE ID == '" + channel + "'")
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

func ChannelExists(channel string) int {
	row := DB.QueryRow("SELECT * FROM Channels WHERE ID == '" + channel + "'")
	var id int
	var res_name, res_des, res_owner string
	err := row.Scan(&id, &res_name, &res_des, &res_owner)
	if err != nil {
		fmt.Println(err.Error())
		return 11
	}
	// defer rows.Close()
	return 0
	// defer channel_rows.Close()
	// for channel_rows.Next() { // Iterate and fetch the records from result cursor
	// 	var user_ID int
	// 	var user_Name string
	// 	var user_Description string
	// 	var user_Owner string
	// 	err := channel_rows.Scan(&user_ID, &user_Name, &user_Description, &user_Owner)
	// 	if err != nil {
	// 		log.Fatal(err.Error())
	// 	}
	// 	// fmt.Println(channel, fmt.Sprintf("%d", user_ID))
	// 	if channel == fmt.Sprintf("%d", user_ID) {
	// 		return 0
	// 	}
	// }
}
