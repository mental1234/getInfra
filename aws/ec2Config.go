package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func AwsArg(region string, tagKey string, tagValue string) {
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
	result, _ := svc.DescribeInstances(input)
    
    for _, reservation := range result.Reservations {
        for _, instance := range reservation.Instances {
            fmt.Println(*instance.InstanceId)
        }
    }
}
