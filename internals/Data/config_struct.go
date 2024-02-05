package Data

type Config struct {
	Dotenv  map[string]string // .env file configs
	Uconfig map[string]string // usr `Config` folder configs
	Server  *Server
}
