package Base

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sunsend/internals/DB"
	"sunsend/internals/Data"

	"github.com/labstack/echo/v4"
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

func ReadJSONMsg(c echo.Context) (Data.InputMsg, int) {
	msg := Data.InputMsg{}
	err := json.NewDecoder(c.Request().Body).Decode(&msg)
	//log.Println(msg)
	if err != nil {
		log.Println("[ERROR] Can't Parse the JSON Message Input (Message) cause " + err.Error())
		return Data.InputMsg{}, 31
	}
	return msg, 0
}

func ReadJSONChan(c echo.Context) (Data.InputChan, int) {
	chl := Data.InputChan{}
	err := json.NewDecoder(c.Request().Body).Decode(&chl)
	if err != nil {
		log.Println("[ERROR] Can't Parse the JSON Message Input (Channel)")
		return Data.InputChan{}, 31
	}
	return chl, 0
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

func DeleteMsgsByID(cid string, msg_id string) int {
	_, rese := FindMsgByMsgID(cid, msg_id, nil)
	if rese != 0 {
		return rese
	}
	res := DB.QueryRow(fmt.Sprintf("DELETE FROM Messages WHERE MID = %s", msg_id))
	err := res.Scan()
	if err != nil && errors.Is(err, sql.ErrNoRows) == false {
		log.Println(err.Error())
		return 33
	}
	return 0
}

func Unmarshal(Rows *sql.Rows) ([]Data.Message, error) {
	data := []Data.Message{}
	defer Rows.Close()
	// fmt.Println(channel_rows)
	for Rows.Next() { // Iterate and fetch the records from result cursor
		var user_CID int
		var user_MID int
		var user_Author string
		var user_Content string
		var user_Date string
		var user_Image string
		var user_ReplyID int
		err := Rows.Scan(
			&user_CID,
			&user_MID,
			&user_Author,
			&user_Content,
			&user_Date,
			&user_Image,
			&user_ReplyID,
		)
		if err != nil {
			return nil, errors.New("ERROR: can't Read from database cause: " + err.Error())
		}
		if user_Image != "None" {
			user_Image = Data.GetValue("config").(*Data.Config).Dotenv["ADR"] + "/" + Data.GetValue("config").(*Data.Config).Uconfig.MediaRoute + "/" + user_Image
		}
		Chat := Data.Message{
			CID:     user_CID,
			MID:     user_MID,
			Author:  user_Author,
			Content: user_Content,
			Date:    user_Date,
			Image:   user_Image,
			ReplyID: user_ReplyID,
		}
		data = append(data, Chat)
	}
	return data, nil
}

func UnmarshalRow(Row *sql.Row) (*Data.Message, error) {

	var user_cid, user_mid, user_ReplyID int
	var user_Author, user_Content, user_Date string
	var user_image string
	err := Row.Scan(
		&user_cid,
		&user_mid,
		&user_Author,
		&user_Content,
		&user_Date,
		&user_image,
		&user_ReplyID,
	)
	if err != nil {
		return nil, errors.New("ERROR: Can't read message from Database cause: " + err.Error())
	}
	msg_obj := &Data.Message{
		CID:     user_cid,
		MID:     user_mid,
		Author:  user_Author,
		Content: user_Content,
		Date:    user_Date,
		Image:   user_image,
		ReplyID: user_ReplyID,
	}
	return msg_obj, nil
}

func FindMsgByUsername(CID string, User string, flags *Data.Flags) (*Data.Message, int) {
	message_rows := DB.QueryRow(
		"SELECT * FROM Messages WHERE CID == " + CID + " AND Author == " + User,
	)
	msg_obj, err := UnmarshalRow(message_rows)
	if err != nil {
		log.Println(err.Error())
		return nil, 31
	}
	return msg_obj, 0
}

func FindMsgByMsgID(CID string, MID string, flags *Data.Flags) (*Data.Message, int) {
	message_rows := DB.QueryRow("SELECT * FROM Messages WHERE CID == " + CID + " AND MID == " + MID)
	msg_obj, err := UnmarshalRow(message_rows)
	if err != nil {
		log.Println(err.Error())
		return nil, 31
	}
	return msg_obj, 0
}

func FindMsgsByChannelID(ID string, flags *Data.Flags) ([]Data.Message, int) {
	var message_rows *sql.Rows
	var res int
	if len(flags.SetRangeMessage) >= 1 {
		message_rows, res = DB.QueryRows(
			"SELECT * FROM Messages WHERE CID == " + ID + " LIMIT " + flags.SetRangeMessage[0] + " OFFSET " + flags.SetRangeMessage[1],
		)
	} else {
		message_rows, res = DB.QueryRows("SELECT * FROM Messages WHERE CID == " + ID)
	}
	// message_rows, res := DB.QueryRows("SELECT * FROM Messages WHERE CID == " + ID)
	if res != 0 {
		return nil, res
	}

	data, err := Unmarshal(message_rows)
	if err != nil {
		return nil, 16
	}
	return data, 0 // handle the error
}

func FindMsgsByUsername(ID string, User string, flags *Data.Flags) ([]*Data.Message, int) {
	var message_rows *sql.Rows
	var res int
	if len(flags.SetRangeMessage) >= 1 {
		message_rows, res = DB.QueryRows(
			"SELECT * FROM Messages WHERE CID == " + ID + " LIMIT " + flags.SetRangeMessage[0] + " OFFSET " + flags.SetRangeMessage[1],
		)
	} else {
		message_rows, res = DB.QueryRows("SELECT * FROM Messages WHERE Author == " + fmt.Sprintf("'%s'", User))
	}
	data := []*Data.Message{}
	// message_rows, res := DB.QueryRows("SELECT * FROM Messages WHERE CID == " + ID)
	if res != 0 {
		return nil, res
	}
	// var asd []byte
	// message_rows.Scan(asd)
	// log.Println("The result if your SQL is:", asd)
	defer message_rows.Close()
	// fmt.Println(channel_rows)
	for message_rows.Next() { // Iterate and fetch the records from result cursor
		var user_CID int
		var user_MID int
		var user_Author string
		var user_Content string
		var user_Date string
		var user_Image string
		var user_ReplyID int
		err := message_rows.Scan(
			&user_CID,
			&user_MID,
			&user_Author,
			&user_Content,
			&user_Date,
			&user_Image,
			&user_ReplyID,
		)
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
			Image:   user_Image,
			ReplyID: user_ReplyID,
		}
		data = append(data, Chat)
	}
	if len(data) == 0 {
		return nil, 21
	}
	return data, 0 // handle the error
}
