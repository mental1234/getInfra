package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func AwsArg(region string, tagKey string, tagValue string) ([]string) {
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
    var instanceIDS []string
    for _, reservation := range result.Reservations {
        for _, instance := range reservation.Instances {
            instanceIDS = append(instanceIDS, *instance.InstanceId)
        }
    }
    return instanceIDS
}
