package main

import (
	"github.com/klauspost/compress/zstd"
)

func Decompress(src []byte) ([]byte, error) {
	var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))
	return decoder.DecodeAll(src, nil)
}

// This is not a main test file, it only works for testing libraries
// And it will not be on release versions
// Ignore this file if you wanna fork the project
func main() {
	// var encoded []byte = Compress([]byte("welcome"))
	// decoded, _ := Decompress(encoded)
	// fmt.Println("Encoded version of that string:", encoded)
	// fmt.Println("UTF-8 version of that encoded:", string(encoded))
	// fmt.Println("Main Byte of that string:", []byte("welcome"))
	// fmt.Printf(string(decoded))
	// msg := &Data.Message{
	// 	CID:     123,
	// 	MID:     124123,
	// 	Author:  "eiko",
	// 	Content: "welcome",
	// 	Date:    "nice",
	// 	ReplyID: 0,
	// }
	// Base.AddArchive(msg)
	// msg2 := &Data.Message{
	// 	CID:     123,
	// 	MID:     124123,
	// 	Author:  "eiko",
	// 	Content: "welcome",
	// 	Date:    "nice",
	// 	ReplyID: 0,
	// }
	// Base.AddArchive(msg2)
	// msg3 := &Data.Message{
	// 	CID:     123,
	// 	MID:     124123,
	// 	Author:  "eiko",
	// 	Content: "welcome",
	// 	Date:    "nice",
	// 	ReplyID: 0,
	// }
	// Base.AddArchive(msg3)
	// msg4 := &Data.Message{
	// 	CID:     123,
	// 	MID:     124123,
	// 	Author:  "eiko",
	// 	Content: "welcome",
	// 	Date:    "nice",
	// 	ReplyID: 0,
	// }
	// Base.AddArchive(msg4)
	// fmt.Println("Level 1 - open the archive file by reading it")
	// res, err := Base.OpenArchive("2024_1_12")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// for i, v := range res {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println("=============================")
	// fmt.Println("Level 2 - open the archive file from temp")
	// for i2, v2 := range Base.GetValue("archive").([]*Data.Message) {
	// 	fmt.Println(i2, v2)
	// }
	//Base.ArchivCheckSystem()
	// defer f.Close()
	// // v, _ := r4.ReadByte()
	// // fmt.Println(v)
	// // v2, _ := r4.ReadByte()
	// // fmt.Println(v2)
	// var org []string
	// for {
	// 	value, err := r4.ReadSlice(byte('+'))
	// 	decoded_str_org, err := Decompress(value)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	// org = append(org, string(value[:len(string(value))-1]))
	// 	org = append(org, string(decoded_str_org))
	// 	if err != nil {
	// 		break
	// 	}

	// 	// fmt.Println(r4.ReadSlice(byte('+')))
	// 	// fmt.Println(r4.ReadSlice(byte('+')))
	// }
	// for index, value := range org {
	// 	fmt.Println(index, value)
	// }
	Run()
}
