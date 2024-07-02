package CoreConfig

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sunsend/internals/Data"
	"time"

	"github.com/joho/godotenv"
)

func GetEnvItems(item string) string {
	data, exists := os.LookupEnv(item)
	if exists == false {
		log.Fatal("[ERROR] Can't Read the .env file")
	}
	return data
}

// var configs *Data.Config
// var rawapikey string
func getEnvConfig() map[string]string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("[ERROR] Can't Read .env file cause: ", err.Error())
	}
	// It can be better to list of all the needed Config and iterate thought them
	ret_values := make(map[string]string)
	ret_values["PORT"] = os.Getenv(
		"PORT",
	) // TODO: User LookUpEnv instead of GetEnv
	//ret_values["HOST"] = os.Getenv("HOST")
	ret_values["HOST"] = GetEnvItems("HOST")
	ret_values["ADR"] = fmt.Sprintf("http://%s:%s", ret_values["HOST"], ret_values["PORT"]) // ~~~
	ret_values["KEY"] = os.Getenv("KEY")                                                    // ~~~
	return ret_values
}

func getUserConfig() Data.UserConfigConf {
	file, err := os.Open("Config/config.json")

	if err != nil {
		log.Println("[WARNING] Can't open .env file cause:", err.Error())
		return Data.UserConfigConf{}
	}
	defer file.Close()

	// sample_config_str := &UserConfigStr{}
	sample_config_str := Data.UserConfigConf{}
	read, err := io.ReadAll(file)
	if err != nil {
		log.Println("[WARNING] Can't read User Config File", err.Error())
		return Data.UserConfigConf{}
	}
	log.Println(sample_config_str)
	// read the json user config file
	if err := json.Unmarshal(read, &sample_config_str); err != nil {
		log.Fatal(err.Error())
		return Data.UserConfigConf{}
	}
	log.Println("[INFO] Read the UserConfigs Succsessfully")
	return sample_config_str
}

func ShowConfigInformation() {
	temp := Data.GetTemp()
	fmt.Println("=========================")
	fmt.Println("Server Config Information")
	fmt.Println("Server Name:", temp.Get("config").(*Data.Config).Uconfig.Name)
	fmt.Println("Server Description:", temp.Get("config").(*Data.Config).Uconfig.Description)
	fmt.Println("Server Owner:", temp.Get("config").(*Data.Config).Uconfig.Owner)
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
	envconfigs := getEnvConfig()
	// TODO: get the user configs here
	userConfig := getUserConfig()

	// fmt.Println(userConfig)
	configs := &Data.Config{
		Dotenv:  envconfigs,
		Uconfig: userConfig, // for now just a little cute nil ^^
	}
	// temp values for better performance
	temp := Data.GetTemp()
	temp.Add("config", configs)
	log.Println("[INFO] All the system updates relaoded successfully")

}
