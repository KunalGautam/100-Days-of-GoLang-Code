package main

import (
	"errors"
	"fmt"
	dotenv "github.com/joho/godotenv"
	"log"
	"os"
)

var MqttDetails = []string{
	"KG_MQTT_HOST", "KG_MQTT_PORT", "KG_MQTT_USER", "KG_MQTT_PASS",
}

func init() {
	fmt.Println("Initializing........")
	var MqttListNotPresent []string
	var DotfileListNotPresent []string
	var DotFileExists = true

	for _, list := range MqttDetails {
		_, ifPresent := os.LookupEnv(list)
		if !ifPresent {
			MqttListNotPresent = append(MqttListNotPresent, list)
		}
	}

	if len(MqttListNotPresent) > 0 {
		fmt.Println("Environment variable(s) not declared. Will try reading from .env file.")
		content, err := dotenv.Read()

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(".env file doesn't exists, creating new one.")
			DotFileExists = false
		}

		if errors.Is(err, os.ErrPermission) {
			// handle the case where the file doesn't exist
			fmt.Println("Check permission of .env file.")
			os.Exit(0)
		}

		if err == nil {
			for _, list := range MqttListNotPresent {
				if _, ok := content[list]; !ok {
					DotfileListNotPresent = append(DotfileListNotPresent, list)
				}
			}
		}

		if (len(DotfileListNotPresent) > 0) || !DotFileExists {
			f, err := os.OpenFile(".env", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			rangeList := DotfileListNotPresent

			if !DotFileExists {
				rangeList = MqttListNotPresent
			}
			// Flush Content
			for _, list := range rangeList {
				var value string
				fmt.Println("Enter Value for ", list)
				_, err := fmt.Scanln(&value)
				if err != nil {
					return
				}
				toWrite := fmt.Sprintf("%s=%s\n", list, value)
				if _, err := f.Write([]byte(toWrite)); err != nil {
					log.Fatal(err)
				}
				fmt.Println(toWrite)
			}

		}

	}
}

func main() {

}
