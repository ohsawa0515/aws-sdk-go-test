package main

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EC2Client -
type EC2Client interface {
	ListEC2Ids() ([]string, error)
}

// ec2Client -
type ec2Client struct {
	client ec2iface.EC2API
}

// NewEC2Client is construct of ec2 object.
func NewEC2Client(svc ec2iface.EC2API) EC2Client {
	return &ec2Client{
		client: svc,
	}
}

// ListEC2Ids lists ids of ec2 instances.
func (svc *ec2Client) ListEC2Ids() ([]string, error) {
	var instances []string
	params := &ec2.DescribeInstancesInput{}
	resp, err := svc.client.DescribeInstances(params)
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
