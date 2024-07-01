package CoreConfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sunsend/internals/Base"
	"sunsend/internals/Data"
	"time"

	"github.com/joho/godotenv"
)

var configs *Data.Config
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
	fmt.Println("Server Name:", configs.Server.Name)
	fmt.Println("Server Description:", configs.Server.Description)
	fmt.Println("Server Owner:", configs.Server.Owner)
	fmt.Println("Server Date:", configs.Server.Date)
	fmt.Println("=========================")
}

// it will be more than these options in the future
func getTimeByMode(mode string) string {
	switch mode {
	case "full":
		return time.Now().String() // just raw current time
	case "normal":
		return fmt.Sprintf(
			"%d/%d/%d - %d:%d:%d",
			time.Now().Year(),
			time.Now().Weekday(),
			time.Now().Day(),
			time.Now().Hour(),
			time.Now().Minute(),
			time.Now().Second(),
		)
	case "small":
		return fmt.Sprintf("%d/%d/%d", time.Now().Year(), time.Now().Weekday(), time.Now().Day())
	default:
		return "" // it will be a Error model to cover this situation
	}
}

func UpdateConfigs() {
	envconfigs, err := getEnvConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO: get the user configs here
	userConfig := getUserConfig()
	// fmt.Println(userConfig)
	configs = &Data.Config{
		Dotenv:  envconfigs,
		Uconfig: nil,                             // for now just a little cute nil ^^
		Bin:     userConfig["Server_Bin"].(bool), // turn on bin option
		Server: &Data.Server{ // TODO: will holds data from user config file
			Name:        userConfig["Server_Name"].(string),
			Description: userConfig["Server_Description"].(string),
			Owner:       userConfig["Server_Owner"].(string),
			Date:        userConfig["Server_Date_Format"].(string), // TODO: Check which type of date format user wants
			Key:         rawapikey,
		},
	}
	// temp values for better performance
	temp := Base.GetTemp()
	temp.Add("config", configs)

}
