package main

import (
	AWSPR "./aws"
	FTOML "./lib"
	"fmt"
	"os"
)

func main() {
	File := os.Args[1]
	region, Tags := FTOML.ReadEC2(File)
    //var mapaEC2 = make(map[string][]string)

	for key, value := range Tags {
		fmt.Println("Key: ", key, "Value: ", value)
        
        fmt.Println(AWSPR.AwsArg(region, key, value))
	}
}
