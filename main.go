package main

import (
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

}
