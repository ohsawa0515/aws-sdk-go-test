package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockEC2iface struct {
	ec2iface.EC2API
}

func (m *mockEC2iface) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {

	return &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{
			{
				Instances: []*ec2.Instance{
					{
						InstanceId: aws.String("i-12345678"),
					},
					{
						InstanceId: aws.String("i-abcdefgh"),
					},
				},
			},
		},
	}, nil
}

func TestListEC2Ids(t *testing.T) {
	mockClient := NewClient(&mockEC2iface{})
	instances, err := mockClient.ListIds()
	if err != nil {
		t.Errorf("Expected no error, but got %v.", err)
	}
	if len(instances) == 0 {
		t.Errorf("Expected list of ec2 instance id, but got empty.")
	}
	expectedInstances := []string{
		"i-12345678",
		"i-abcdefgh",
	}
	for i, instance := range instances {
		if expectedInstances[i] != instance {
			t.Errorf("Expected %s, but got %s.", expectedInstances[i], instance)
		}
	}
}
