// main function for quetzal project
// First feature is describe instances
package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func descInst() {
	fmt.Println("Hello World")
}

func main() {
	instanceAws := os.Args[1]
	// Session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	// Create EC2 service client
	svc := ec2.New(sess)
	// Input
	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceAws),
		},
	}
	// Describe instances
	result, err := svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Sucess ", result)
		descInst()
	}
}
