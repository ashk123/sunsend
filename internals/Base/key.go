package Base

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"strings"
	"sunsend/internals/Data"
)

// Doc: https://dev.to/caiorcferreira/implementing-a-safe-and-sound-api-key-authorization-middleware-in-go-3g2c

func GenerateAPIKey(rawkey string) [32]byte {
	// token := randstr.Hex(16) // generate 128-bit hex string
	return sha256.Sum256([]byte(rawkey))
}

func BearerToken(headers map[string][]string) (string, int) {
	fmt.Println(headers)
	api_key_org, ok := headers["Api_key"]
	if !ok {
		return "", 17 // invalid API KE	Y
	}
	if len(strings.SplitN(api_key_org[0], " ", 2)) > 2 {
		return "", 17
	}
	return strings.TrimSpace(api_key_org[0]), 0

}

// apiKeyIsValid checks if the given API key is valid and returns the principal if it is.
func ApiKeyIsValid(user_api_key string) int {
	// hash := sha256.Sum256([]byte(user_api_key))
	// key := hash[:]
	temp := GetTemp()
	if subtle.ConstantTimeCompare([]byte(temp.Get("config").(*Data.Config).Server.Key), []byte(user_api_key)) == 1 {
		return 0
	}

	return 17
}
