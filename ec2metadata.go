package main

import "github.com/aws/aws-sdk-go/aws/ec2metadata"

// GetInstanceID returns EC2 instance ID getting by metadata.
func GetInstanceID(svc *ec2metadata.EC2Metadata) (string, error) {
	return svc.GetMetadata("instance-id")
}
