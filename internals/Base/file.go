package Base

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

const (
	FILE_SIZE int64 = (1024 ^ 2) * 30 // 30MB Of data
)

func DecompressFile(filename string) ([]byte, int) {
	//f, _ := os.Create("/tmp/TEMP_" + filename)
	compresed_file, _ := os.ReadFile("Storage/" + filename + ".zst")
	file_data, err := Decompress(compresed_file)
	if err != nil {
		return nil, 26
	}
	return file_data, 0
	//f.Write(file_daata)
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func IsFileExists(filePath string) bool {

	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}
func SaveImageFile(image *multipart.FileHeader) error {
	//buf := new(strings.Builder)
	// check errors
	//fmt.Println(buf.String())

	f, err := os.Create("Storage/" + image.Filename + ".zst")
	if err != nil {
		log.Println("ERROR: err to save file")
		return errors.New("can't create file")
	}
	actual, err := image.Open()
	if err != nil {
		log.Println("ERROR: err to save file 1")
		return errors.New("Can't Read Image")
	}
	image_bytes_org := StreamToByte(actual)
	//io.Copy(buf, actual)
	//fmt.Println(image_bytes_org)
	compresed_image_data := Compress(image_bytes_org)

	//if _, err = io.Copy(f, actual); err != nil {
	//	return errors.New("Can't Copy the Image Data")
	//}
	f.Write(compresed_image_data)
	defer f.Close()
	return nil
}

func HandleFiles(file *multipart.FileHeader) int {
	if len(file.Filename) > 50 {
		return 23
	}
	switch value := file.Filename; {
	case strings.Contains(value, ".png"),
		strings.Contains(value, ".jpg"),
		strings.Contains(value, ".bmp"):
		break
	default:
		return 24
	}

	if file.Size >= FILE_SIZE {
		return 25
	}
	return 0
}
