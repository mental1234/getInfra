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
	file := os.Args[1]
	fmt.Println(file)
	Region, keys, values := L.ReadFile(file)
	fmt.Println(Region[0])
	fmt.Println(keys)
	fmt.Println(values)
	region := os.Args[2]
	key := os.Args[3]
	key = "tag:" + key
	value := os.Args[4]
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
	X, Y, Z := L.ReadFile("/home/ramon/Projects/QuetzalProject/config/config.json")
	fmt.Println(X, Y, Z)

}
