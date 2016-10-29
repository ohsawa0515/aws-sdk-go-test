package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// ListEC2Ids lists ids of ec2 instances.
func ListEC2Ids() ([]string, error) {

	region := "ap-northeast-1"
	sessOpt := session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(region)},
	}
	sess, err := session.NewSessionWithOptions(sessOpt)
	if err != nil {
		log.Fatal(err)
	}

	var instances []string
	svc := ec2.New(sess)
	params := &ec2.DescribeInstancesInput{}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		return instances, err
	}
	for _, res := range resp.Reservations {
		for _, instance := range res.Instances {
			instances = append(instances, *instance.InstanceId)
		}
	}
	return instances, nil
}
