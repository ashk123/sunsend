package Base

import (
	"fmt"
	"log"
	"sunsend/internals/DB"
)

// There should be a ChannelCollection too
// that Holds the Channel structure
// Like ChatColelction

// Base Chat Structure
type ChatCollection struct {
	CID     int
	MID     int
	Author  string
	Content string
	Date    string
	ReplyID int
}

func FindMsgsByChannelID(ID string) ([]*ChatCollection, int) {
	data := []*ChatCollection{}
	channel_rows, res := DB.QueryMessages()
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
			Chat := &ChatCollection{
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

func ChannelExists(channel string) int {
	channel_rows, res := DB.QueryChannel()
	if res != 0 {
		// fmt.Println("there is a bug here")
		return res
	}
	defer channel_rows.Close()
	for channel_rows.Next() { // Iterate and fetch the records from result cursor
		var user_ID int
		var user_Name string
		var user_Description string
		var user_Owner string
		err := channel_rows.Scan(&user_ID, &user_Name, &user_Description, &user_Owner)
		if err != nil {
			log.Fatal(err.Error())
		}
		// fmt.Println(channel, fmt.Sprintf("%d", user_ID))
		if channel == fmt.Sprintf("%d", user_ID) {
			return 0
		}
	}
	return 11

}
