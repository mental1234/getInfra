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

	for key, value := range Tags {
		fmt.Println("Key: ", key, "Value: ", value)
		AWSPR.AwsArg(region, key, value)
	}
}
