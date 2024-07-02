package Base

import (
	"database/sql"
	"errors"
	"log"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
)

func CUnmarshal(Rows *sql.Rows) (*[]Data.Channel, error) {
	data := []Data.Channel{}
	defer Rows.Close()
	// fmt.Println(channel_rows)
	for Rows.Next() { // Iterate and fetch the records from result cursor
		var user_ID int
		var user_Name string
		var user_Description string
		var user_Owner string
		err := Rows.Scan(
			&user_ID,
			&user_Name,
			&user_Description,
			&user_Owner,
		)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("ERROR: can't Read from database cause: " + err.Error())
		}
		Chat := Data.Channel{
			ID:          user_ID,
			Name:        user_Name,
			Description: user_Description,
			Owner:       user_Owner,
		}
		data = append(data, Chat)
	}
	return &data, nil
}

func GetChannelList() (*[]Data.Channel, int) {
	rows, err := DB.QueryRows("SELECT * FROM Channels")
	if err != 0 {
		return nil, 30
	}
	data, err1 := CUnmarshal(rows)
	if err1 != nil {
		return nil, 30
	}
	return data, 0
}
