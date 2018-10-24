package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
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

	client := NewClient(ec2.New(sess))
	instances, err := client.ListIds()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(instances)
}
