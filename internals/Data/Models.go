package Data

/*
	This is the Publix model of sevrers, if you see any structure here,
	that means it is public for all files
	You can see these struct everywhere in the source code
*/

// Global Constant Values

const (
	API_VERSION string = "v1"
)

// Config Model for server
type Config struct {
	Dotenv  map[string]string // .env file configs
	Uconfig UserConfigConf    // usr `Config` folder configs
}

type InputMsg struct {
	Username string `json:"username",omitempty`
	Message  string `json:"message",omitempty`
	Reply    int    `json:"reply",omitempty`
}

type InputChan struct {
	ID          int    `json:"id",omitempty`
	Name        string `json:"name",omitempty`
	Description string `json:"description",omitempty`
	Owner       string `json:"owner",omitempty`
}

// Flags model for getMessageParameters
type Flags struct {
	SetRangeMessage []string
}
type UserConfigConf struct {
	Name        string `json:"Server_Name",omitempty`
	Description string `json:"Server_Description",omitempty`
	Owner       string `json:"Server_Owner",omitempty`
	MediaRoute  string `json:"Media_Route",omitempty`
}
