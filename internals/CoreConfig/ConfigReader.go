package CoreConfig

import (
	"errors"
	"log"
	"os"
	"sunsend/internals/Data"

	"github.com/joho/godotenv"
)

var Configs *Data.Config
var rawapikey string

func getEnvConfig() (map[string]string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		// log.Fatal(err.Error())
		return nil, errors.New("invalid or can't read .env file")
	}
	// It can be better to list of all the needed Config and iterate thought them
	ret_values := make(map[string]string)
	ret_values["PORT"] = os.Getenv("PORT")
	ret_values["KEY"] = os.Getenv("KEY")
	// copy(rawapikey[:], ret_values["KEY"]) // fix fixing some invalid memory address - TODO: Fix in better way - nice
	rawapikey = ret_values["KEY"]
	return ret_values, nil
}

func UpdateConfigs() {
	envconfigs, err := getEnvConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: get the user configs here

	Configs = &Data.Config{
		Dotenv:  envconfigs,
		Uconfig: nil, // for now just a little cute nil ^^
		Server: &Data.Server{ // TODO: will holds data from user config file
			Name:        "test",
			Description: "test1",
			Owner:       "test",
			Date:        "test",
			Key:         rawapikey,
		},
	}
}
