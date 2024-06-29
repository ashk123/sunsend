package Base

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

const (
	FILE_SIZE int64 = 1024 * 10
)

func SaveImageFile(image *multipart.FileHeader) error {
	//buf := new(strings.Builder)
	//n, err := io.Copy(buf, r)
	// check errors
	//fmt.Println(buf.String())

	f, err := os.Create("Storage/" + image.Filename)
	if err != nil {
		log.Println("ERROR: err to save file")
		return errors.New("can't create file")
	}
	actual, err := image.Open()
	if err != nil {
		log.Println("ERROR: err to save file 1")
		return errors.New("Can't Read Image")
	}
	if _, err = io.Copy(f, actual); err != nil {
		return errors.New("Can't Copy the Image Data")
	}
	defer f.Close()
	return nil
}

func HandleFiles(file *multipart.FileHeader) int {
	if len(file.Filename) > 10 {
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
