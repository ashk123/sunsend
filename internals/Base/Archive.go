package Base

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sunsend/internals/DB"
	"sunsend/internals/Data"
	"time"

	"github.com/klauspost/compress/zstd"
)

/*
	The main structue of this API is the server will encode the old
	messages and archive them inside the program for saving it as zstd
	compression, and if these message Will not call again, server will archive/remove
	These messages forever.

	Archiving messages is a new way of saving data when system
	has more than a lot of data (more than limit) or users don't want their messages.
	Those message will removed forever or archive as a encoded data
	If some conditions will be true.

	- There are 3 condition for deleting files
	1. Each 5 minute, server will try to get some old data by their Date
		And if user don't want to remove it, those data will archive/lose forever.
	2. When user delete a message that message will be a archive encoding data
		And only remove when user want to remove it completely.
	3. manager of server can do these steps manually

	You can turn on or off this feature in the configuration file
	If you wanna know more about the Archiving feature
	Read the Archive.md documentation.
*/

// type ArchiveMessage struct {
// 	A_CID int
// 	A_MID int
// 	A_Author
// 	A_Content string
// 	A_Date    string
// }

func Compress(src []byte) []byte {
	var encoder, err = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	if err != nil {
		log.Println(err.Error())
	}
	return encoder.EncodeAll(src, make([]byte, 0, len(src)))
}
func Decompress(src []byte) ([]byte, error) {
	var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))
	return decoder.DecodeAll(src, nil)
}

func AddArchive(message *Data.Message) {
	var res_CID int = message.CID
	var res_MID int = message.MID
	var res_Author string = message.Author
	var res_Content string = message.Content
	var res_Date string = message.Date
	var res_ReplyID int = message.ReplyID
	encoded_text_org := Itsr(
		res_CID,
	) + "-" + Itsr(
		res_MID,
	) + "-" + res_Author + "-" + res_Content + "-" + res_Date + "-" + Itsr(
		res_ReplyID,
	) + "+"
	encoded_res := Compress([]byte(encoded_text_org))
	fmt.Printf("%d-%d-%d\n", time.Now().Year(), time.Now().Weekday(), time.Now().Day())
	f, err := os.OpenFile(
		"Archive_"+fmt.Sprintf(
			"%d_%d_%d",
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
		)+".arc",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		return
	}
	defer f.Close()
	f.Write(encoded_res)
	log.Println("[INFO] System Created a Archive file of old message with ID:", res_MID)
}

func GetMessageDate() (*[]Data.Message, int) {
	var data []Data.Message
	count := getRowsLength()
	if count < 1 || count == -1 {
		return nil, -1 // error there is not enought messages in the DB
	}
	//if count < Data.ARCHIVE_LIMIT {
	//	return nil, 1
	//}
	//data, res := getMessagDate()
	//if res != 0 {
	//	return nil, -2
	//}
	return &data, 0

}

func SubMonths(date string) string {
	t, err := time.Parse("2006/1/2", date)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%d/%d/%d", t.Year(), t.Month()-1, t.Day())
}
func GetOldMessages() (*[]Data.Message, int) {
	t := time.Now()
	ftime := fmt.Sprintf("%d/%d/%d", t.Year(), t.Month(), t.Day())
	//stime := SubMonths(ftime)
	//fmt.Println(ftime, stime)
	row, res := DB.QueryRows(
		fmt.Sprintf(
			"SELECT * FROM Messages WHERE Date < '%s' AND Date > '%s';",
			ftime,
			t.AddDate(0, -3, 0),
		),
	)
	if res != 0 {
		log.Println("Error code is:", res)
		return nil, -1
	}
	data, err := Unmarshal(row)
	if err != nil {
		log.Println(err.Error())
		return nil, -1
	}

	if len(data) <= 0 {
		return nil, 1
	}
	// TODO: Replace the chcekCount function with a simple Scan
	//return checkCount(rows)
	return &data, 0
}

func getRowsLength() int {
	row := DB.QueryRow("SELECT COUNT(*) as count FROM Messages")
	var count int
	row.Scan(&count)
	// TODO: Replace the chcekCount function with a simple Scan
	//return checkCount(rows)
	return count
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Println(err.Error())
			return -1
		}
	}
	return count
}

func ArchivCheckSystem() {
	msgs, res := GetOldMessages()
	if res != 0 {
		log.Println("there is not any old messages")
	}
	for i, v := range *msgs {
		fmt.Println(i, v)
	}
}

func OpenArchive(date string) ([]*Data.Message, error) {
	var result []*Data.Message
	f, err := os.ReadFile("Archive_" + date + ".arc")
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("there is not any Archive_" + date + ".arc file!")
	}
	// fmt.Println(f)
	res, err := Decompress(f)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("can't compress the archive file")
	}
	data := string(res)

	// *result* slice can be a simple map, to be faster
	for _, v := range strings.Split(data[:len(data)-1], "+") {
		sub_data := strings.Split(v, "-")
		user_CID, _ := strconv.Atoi(sub_data[0])
		user_MID, _ := strconv.Atoi(sub_data[1])
		user_ReplyID, _ := strconv.Atoi(sub_data[5])
		result = append(result, &Data.Message{
			CID:     user_CID,
			MID:     user_MID,
			Author:  sub_data[2],
			Content: sub_data[3],
			Date:    sub_data[4],
			ReplyID: user_ReplyID,
		})
	}
	if len(result) < 1 {
		return nil, errors.New("there is not any data in the slice")
	}
	// Temp the result for better performance
	temp := Data.GetTemp()
	temp.Add("archive", result)
	return result, nil
}

// // Generate New archive object, if there is not any object
// func SetArchiveObj() *Archive {
// 	if archive == nil {
// 		archive = &Archive{
// 			Last:   nil,
// 			Data:   []*Data.Message{},
// 			Length: 0,
// 		}
// 		return archive
// 	} else {
// 		return archive
// 	}
// }
