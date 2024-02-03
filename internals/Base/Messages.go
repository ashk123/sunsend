package Base

import (
	"fmt"
	"log"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
)

func IsEqulTo(message string) bool {
	return message == "321"
}

func CheckMessage(message string) int {
	if len(message) > 30 {
		return 15
	} else if IsEqulTo(message) {
		return 12
	}
	return 0
}

func FindMsgsByChannelID(ID string) ([]*Data.Chat, int) {
	data := []*Data.Chat{}
	channel_rows, res := DB.QueryRows("SELECT * FROM Messages")
	if res != 0 {
		return nil, res
	}
	defer channel_rows.Close()
	// fmt.Println(channel_rows)
	for channel_rows.Next() { // Iterate and fetch the records from result cursor
		var user_CID int
		var user_MID int
		var user_Author string
		var user_Content string
		var user_Date string
		var user_ReplyID int
		err := channel_rows.Scan(&user_CID, &user_MID, &user_Author, &user_Content, &user_Date, &user_ReplyID)
		if err != nil {
			log.Fatal(err.Error())
		}
		if fmt.Sprintf("%d", user_CID) == ID {
			Chat := &Data.Chat{
				CID:     user_CID,
				MID:     user_MID,
				Author:  user_Author,
				Content: user_Content,
				Date:    user_Date,
				ReplyID: user_ReplyID,
			}
			data = append(data, Chat)
		}
	}
	return data, 0 // handle the error
}
