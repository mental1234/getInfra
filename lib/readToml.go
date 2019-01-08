package lib

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type awsConfig struct {
	Region string
	Tags map[string]tagsArray
}

type tagsArray struct {
	Key string
	Value string
}

func ReadFile(argFile string) (string, map[string]string) {
	var config awsConfig
	if _, err := toml.DecodeFile(argFile, &config); err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("Region: %s\n", config.Region)
	var mapTags = make(map[string]string)
	for  _, tagsArray := range config.Tags {
		mapTags[tagsArray.Key] = tagsArray.Value
	}
	return config.Region, mapTags

}
