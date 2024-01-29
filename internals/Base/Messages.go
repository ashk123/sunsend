package Base

func IsEqulTo(message string) bool {
	if message == "321" {
		return true
	}
	return false
}

func CheckMessage(message string) int {
	if len(message) > 30 {
		return 15
	} else if IsEqulTo(message) {
		return 12
	}
	return 0
}
