package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockEC2Iface struct {
	ec2iface.EC2API
}

func (svc *mockEC2Iface) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {

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

func TestList(t *testing.T) {
	mockSvc := &mockEC2Iface{}
	instances, err := ListEC2Ids(mockSvc)
	if err != nil {
		t.Errorf("Expected no error for empty input, but got %v.", err)
	}
	if len(instances) == 0 {
		t.Errorf("Expected list of ec2 instance id for empty input, but got empty.")
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
