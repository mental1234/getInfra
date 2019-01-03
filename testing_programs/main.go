package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type awsConfig struct {
	Tags map[string]tagsArray 
	Servers map[string]server
}

type tagsArray struct {
	Key string
	Value string
}

type server struct {
	IP string
	DC string
}

func main() {
	var config awsConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	for tag, tagsArray := range config.Tags {
		fmt.Printf("Server: %s (%s, %s)\n", tag, tagsArray.Key, tagsArray.Value)
	}

	for serverName, server := range config.Servers {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
	}
	
}

