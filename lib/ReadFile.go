// read json configuration files
package lib

import (
	"fmt"
	"os"
	"encoding/json"
)


type Config struct {
	Region  []string
	Keys	[]string
	Values  []string
}

func ReadFile(argFile string) ([]string, []string, []string) {
	file, _ := os.Open(argFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return configuration.Region, configuration.Keys, configuration.Values
}
