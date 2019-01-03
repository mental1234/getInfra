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

func ReadFile(argFile string) string{
	var config awsConfig
	if _, err := toml.DecodeFile(argFile, &config); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello World")
	return "HELLO"
}
