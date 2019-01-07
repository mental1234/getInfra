package main

import (
	L "./lib"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	
	region := os.Args[1]
	key := os.Args[2]
	key = "tag:" + key
	value := os.Args[3]
	//file := os.Args[1]
	// Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	// Create EC2 service client
	svc := ec2.New(sess)

	input := &ec2.DescribeInstancesInput{
    		Filters: []*ec2.Filter{
        	{
            		Name: aws.String(key),
            			Values: []*string{
                		aws.String(value),
            			},
        		},
    		},
	}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}
//	X := L.ReadFile("/home/ramon/Projects/quetzal/config/config.toml")
	X := L.ReadFile("/home/ramon/quetzal/config/config.toml")
	fmt.Println(X)
	for key, value := range X {
		fmt.Println("Key: ", key, "Value: ", value)
	}
//	fmt.Println(L.ReadFile("/home/ramon/Projects/quetzal/config/config.toml"))
//	fmt.Println(L.ReadFile("/home/ramon/quetzal/config/config.toml"))
}
