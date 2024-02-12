package Base

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
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
	encoded_text_org := Itsr(res_CID) + "-" + Itsr(res_MID) + "-" + res_Author + "-" + res_Content + "-" + res_Date + "-" + Itsr(res_ReplyID) + "+"
	encoded_res := Compress([]byte(encoded_text_org))
	f, _ := os.OpenFile("Archive_"+fmt.Sprintf("%d_%d_%d", time.Now().Year(), time.Now().Weekday(), time.Now().Day())+".arc", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.Write(encoded_res)
	log.Println("System Created a archive file of old messages")
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
	temp := GetTemp()
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
