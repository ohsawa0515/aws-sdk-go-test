package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	region = "ap-northeast-1"
)

func main() {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(region),
		}))

	// EC2
	ec2Client := NewClient(ec2.New(sess))
	instances, err := ec2Client.ListIds()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instances)

	// EC2metadata
	ec2MetadataSvc := ec2metadata.New(sess)
	instanceID, err := GetInstanceID(ec2MetadataSvc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instanceID)
}
