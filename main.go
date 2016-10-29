package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func runEc2() {

	region := "ap-northeast-1"
	sessOpt := session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(region)},
	}
	sess, err := session.NewSessionWithOptions(sessOpt)
	if err != nil {
		log.Fatal(err)
	}

	instances, err := ListEC2Ids(ec2.New(sess))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instances)
}

func main() {
	runEc2()
}
