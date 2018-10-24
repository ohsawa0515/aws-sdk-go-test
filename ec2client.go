package main

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EC2iface -
type EC2iface interface {
	ListIds() ([]string, error)
}

// Instance -
type Instance struct {
	client ec2iface.EC2API
}

// NewClient is construct of ec2 object.
func NewClient(svc ec2iface.EC2API) EC2iface {
	return &Instance{
		client: svc,
	}
}

// ListIds return list of ids.
func (i *Instance) ListIds() ([]string, error) {
	var instances []string

	resp, err := i.client.DescribeInstances(&ec2.DescribeInstancesInput{})
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
