package lib

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type awsConfig struct {
	Tags map[string]tagsArray
}

type tagsArray struct {
	Key string
	Value string
}

func ReadFile(argFile string) map[string]tagsArray{
	var config awsConfig
	if _, err := toml.DecodeFile(argFile, &config); err != nil {
		fmt.Println(err)
	}
	
	for tag, tagsArray := range config.Tags {
		fmt.Printf("Server: %s (%s, %s)\n", tag, tagsArray.Key, tagsArray.Value)
		
	}
	return config.Tags

}
