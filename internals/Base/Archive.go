package Base

import (
	"fmt"
	"math/rand"
	"sunsend/internals/CoreConfig"
	"time"

	"github.com/klauspost/compress/zstd"
)

/*
	The main structue of this API is the server will encode the old
	messages and archive them inside the program and if these message
	Will not call again, server will remove these message forever.

	Archiving messages is a new way of saving data when system
	has more than a lot of data (more than limit) or users don't want their messages.
	Those message will removed forever if some conditions are true.

	- There are 3 condition for deleting files
		1. Each 5 minute, server will try to get some old data by their Date
		And if user don't want to remove it, those data will lose forever.
		2. When user delete a message that message will be a archive encoding data
		And only remve when user want to remove it completely.
		3. manager of server can do this steps manually

	You can turn on or off this feature in the configuration file
	If you wanna know more about the Archiving feature
	Read the Archive.md documentation.
*/

type ArchiveMsg struct {
	ID      int
	Content []byte
	Date    string
}

type Archive struct {
	Last   ArchiveMsg
	Data   []ArchiveMsg // we can use map instead of slice too
	Length int          // this will be something like int32 or in64 in the future
}

var archive *Archive // Main archive object

func Compress(src []byte) []byte {
	var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	return encoder.EncodeAll(src, make([]byte, 0, len(src)))
}

func (p *Archive) AddArchive(data string) {
	encoded := Compress([]byte(data))
	gen_archive_msg := ArchiveMsg{
		ID:      rand.Intn(10),
		Content: encoded,
		Date:    fmt.Sprintf("%d/%d/%d", time.Now().Hour(), time.Now().Weekday(), time.Now().Day()),
	}
	p.Last = gen_archive_msg
	p.Length = p.Length + 1 // add new msgs to the length of the Datas slice
	p.Data = append(p.Data, gen_archive_msg)
}

func (p *Archive) RemoveArchive(ID int) {
	for _, v := range p.Data {
		if v.ID == ID {
			// remove item from here
			fmt.Println(v)
		}
	}
}

func (p *Archive) GetLength() int {
	return p.Length
}

func (p *Archive) GetLast() *ArchiveMsg {
	return &p.Last
}

func getArchiveObj() *Archive {
	if CoreConfig.Configs.Bin == true {
		return archive
	} else {
		return nil
	}
}

func (p *Archive) SaveArchiveMessages() int {
	// save the messages with zstd format to the server
	return 0 // result status code
}

func (p *Archive) UpdateMessages() {
	// get old messages from database and operate on them
	// SaveArchiveMessages will call in this function
	// Let's do some cool things ;D
}
