package Base

type ChatCollection []string

func FindMsgsByChannelID(ID string) ChatCollection {
	data := []string{"welcome", "nice", "this is so good"}
	if ID == "123" {
		return data
	}
	return nil // handle the error
}
