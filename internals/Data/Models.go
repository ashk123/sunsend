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
