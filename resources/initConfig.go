package resources

import (
	"encoding/json"
	"os"
)

var AppConfig LoadConfig

func InitConfig() {
	data, err := os.ReadFile("resources/config.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &AppConfig)
}
