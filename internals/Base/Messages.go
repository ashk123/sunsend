package Base

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
)

func IsEqulTo(message string) bool {
	for _, v := range *Data.LoadedWordList {
		if message == v || strings.Contains(message, v) {
			return true
		}
	}
	return false
}

func CheckMessage(message string) int {
	if len(message) > 30 {
		return 15 //  response 15 -> length error
	} else if IsEqulTo(message) {
		return 12 // response 12 -> Word error
	}
	return 0
}

// Very simple int to string function
func Itsr(data int) string {
	return fmt.Sprintf("%d", data)
}

// GetmessageByOffset will gives you, your own limit and offset for recieve  messages
// For example, if you wanna show 10 messages for your first page, it's better to get *those* 10 first message
// From Server and make another request for the rest of messages for your other pages
// If you wanan have all of the data, you can use `FindMsgsByChannelID` Function
// func GetMessageByOffset(channel_id string, start int, finish int) ([]*Data.Message, int) {
// 	var org []*Data.Message
// 	srow, sres := DB.QueryRows("SELECT * FROM Messages WHERE CID == " + channel_id + " LIMIT " + Itsr(start) + " OFFSET " + Itsr(finish))
// 	if sres != 0 {
// 		return nil, sres
// 	}
// 	for srow.Next() {
// 		var user_CID int
// 		var user_MID int
// 		var user_Author string
// 		var user_Content string
// 		var user_Date string
// 		var user_ReplyID int
// 		err := srow.Scan(&user_CID, &user_MID, &user_Author, &user_Content, &user_Date, &user_ReplyID)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return nil, 16
// 		}
// 		message_obj_result := &Data.Message{
// 			CID:     user_CID,
// 			MID:     user_MID,
// 			Author:  user_Author,
// 			Content: user_Content,
// 			Date:    user_Date,
// 			ReplyID: user_ReplyID,
// 		}
// 		org = append(org, message_obj_result)
// 	}
// 	return org, 0

// }

func FindMsgByChannelID(CID string, MID string, flags *Data.Flags) (*Data.Message, int) {
	message_rows := DB.QueryRow("SELECT * FROM Messages WHERE CID == " + CID + " AND MID == " + MID)
	var user_cid, user_mid, user_ReplyID int
	var user_Author, user_Content, user_Date string
	err := message_rows.Scan(&user_cid, &user_mid, &user_Author, &user_Content, &user_Date, &user_ReplyID)
	if err != nil {
		return nil, 19
	}
	msg_obj := &Data.Message{
		CID:     user_cid,
		MID:     user_mid,
		Author:  user_Author,
		Content: user_Content,
		Date:    user_Date,
		ReplyID: user_ReplyID,
	}
	return msg_obj, 0
}

func FindMsgsByChannelID(ID string, flags *Data.Flags) ([]*Data.Message, int) {
	var message_rows *sql.Rows
	var res int
	if len(flags.SetRangeMessage) >= 1 {
		message_rows, res = DB.QueryRows("SELECT * FROM Messages WHERE CID == " + ID + " LIMIT " + flags.SetRangeMessage[0] + " OFFSET " + flags.SetRangeMessage[1])
	} else {
		message_rows, res = DB.QueryRows("SELECT * FROM Messages WHERE CID == " + ID)
	}
	data := []*Data.Message{}
	// message_rows, res := DB.QueryRows("SELECT * FROM Messages WHERE CID == " + ID)
	if res != 0 {
		return nil, res
	}
	defer message_rows.Close()
	// fmt.Println(channel_rows)
	for message_rows.Next() { // Iterate and fetch the records from result cursor
		var user_CID int
		var user_MID int
		var user_Author string
		var user_Content string
		var user_Date string
		var user_ReplyID int
		err := message_rows.Scan(&user_CID, &user_MID, &user_Author, &user_Content, &user_Date, &user_ReplyID)
		if err != nil {
			log.Println(err.Error())
			return nil, 16
		}
		Chat := &Data.Message{
			CID:     user_CID,
			MID:     user_MID,
			Author:  user_Author,
			Content: user_Content,
			Date:    user_Date,
			ReplyID: user_ReplyID,
		}
		data = append(data, Chat)
	}
	return data, 0 // handle the error
}
