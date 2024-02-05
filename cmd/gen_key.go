package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/thanhpk/randstr"
)

func GenerateAPIKey() [32]byte {
	token := randstr.Hex(16) // generate 128-bit hex string
	return sha256.Sum256([]byte(token))
}

func main() {
	res := GenerateAPIKey()
	fmt.Println("Your API KEY is: ", fmt.Sprintf("%X", res))
	// TODO: Add new key to the .env file
}
