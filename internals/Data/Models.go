package Data

/*
	This is the Publix model of sevrers, if you see any structure here,
	that means it is public for all files
	You can see these struct everywhere in the source code
*/

// Global Constant Values
const (
	API_VERSION   string = "v1"
	HOST          string = "http://127.0.0.1:3000/api/" + API_VERSION
	MEDIA_ROUTE   string = HOST + "/media"
	ARCHIVE_LIMIT int    = 30
)

// Config Model for server
type Config struct {
	Dotenv  map[string]string // .env file configs
	Uconfig map[string]string // usr `Config` folder configs
	Bin     bool
	Server  *Server
}

// Flags model for getMessageParameters
type Flags struct {
	SetRangeMessage []string
}

// Server model for server specific configs
type Server struct {
	Name        string
	Description string
	Owner       string
	Date        string
	Key         string
}
