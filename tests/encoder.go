package main

import (
	"fmt"

	"github.com/klauspost/compress/zstd"
)

// This is not a main test file, it only works for testing libraries
// And it will not be on release versions
// Ignore this file if you wanna fork the project
func main() {
	var encoded []byte = Compress([]byte("welcome"))
	decoded, _ := Decompress(encoded)
	fmt.Println("Encoded version of that string:", encoded)
	fmt.Println("UTF-8 version of that encoded:", string(encoded))
	fmt.Println("Main Byte of that string:", []byte("welcome"))
	fmt.Printf(string(decoded))
}

func Compress(src []byte) []byte {
	var encoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	return encoder.EncodeAll(src, make([]byte, 0, len(src)))
}

func Decompress(src []byte) ([]byte, error) {
	var decoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(0))
	return decoder.DecodeAll(src, nil)
}
