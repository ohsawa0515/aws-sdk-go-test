package main

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// ListEC2Ids lists ids of ec2 instances.
func ListEC2Ids(svc ec2iface.EC2API) ([]string, error) {

	var instances []string
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
