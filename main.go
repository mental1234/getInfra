package main

import (
	FTOML "./lib"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)


func awsArg(region string, tagKey string, tagValue string) {
	tagKey = "tag:" + tagKey
	// Session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	// Create EC2 Service client
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
		{
			Name: aws.String(tagKey),
				Values: []*string{
				aws.String(tagValue),
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

func main() {
	File := os.Args[1]
	region, Tags := FTOML.ReadFile(File)

	//fmt.Println(region)
	for key, value := range Tags {
		fmt.Println("Key: ", key, "Value: ", value)
		awsArg(region, key, value)
	}
}
