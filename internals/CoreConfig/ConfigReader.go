package CoreConfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sunsend/internals/Data"
	"time"

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
	// copy(rawapikey[:], ret_values["KEY"])
	rawapikey = ret_values["KEY"] // fix fixing some invalid memory address - TODO: Fix in better way - nice
	return ret_values, nil
}

func getUserConfig() map[string]interface{} {
	file, err := os.Open("Config/config.json")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// sample_config_str := &UserConfigStr{}
	var sample_config_str map[string]interface{}
	read, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(read, &sample_config_str) // read the json user config file

	return sample_config_str
}

func ShowConfigInformation() {
	fmt.Println("=========================")
	fmt.Println("Server Config Information")
	fmt.Println("Server Name:", Configs.Server.Name)
	fmt.Println("Server Description:", Configs.Server.Description)
	fmt.Println("Server Owner:", Configs.Server.Owner)
	fmt.Println("Server Date:", Configs.Server.Date)
	fmt.Println("=========================")
}

func UpdateConfigs() {
	envconfigs, err := getEnvConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: get the user configs here
	userConfig := getUserConfig()
	// fmt.Println(userConfig)
	Configs = &Data.Config{
		Dotenv:  envconfigs,
		Uconfig: nil, // for now just a little cute nil ^^
		Server: &Data.Server{ // TODO: will holds data from user config file
			Name:        userConfig["Server_Name"].(string),
			Description: userConfig["Server_Description"].(string),
			Owner:       userConfig["Server_Owner"].(string),
			Date:        time.Now().String(), // TODO: Check which type of date format user wants
			Key:         rawapikey,
		},
	}
}