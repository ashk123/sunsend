package Base

func IsEqulTo(message string) bool {
	return message == "321"
}

func CheckMessage(message string) int {
	if len(message) > 30 {
		return 15
	} else if IsEqulTo(message) {
		return 12
	}
	return 0
}
