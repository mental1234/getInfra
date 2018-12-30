// read toml configuration files
package lib

import (
	"fmt"
	"os"
	"encoding/json"
)


type Config struct {
	Keys	[]string
	Values  []string
}


func ReadFile() {
	file, _ := os.Open("/home/ramon/Projects/QuetzalProject/config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(configuration.Keys)
	fmt.Println(configuration.Values)
}
