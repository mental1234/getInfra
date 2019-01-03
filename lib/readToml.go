// read json configuration files
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Vendor string
	Service EC2info
}


type EC2info struct {
	EC2 string
	Region string
	Tags map[string]tags
}

func main() {
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	
	for serverName, EC2info := range config.EC2info.Tags {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
	}

}
