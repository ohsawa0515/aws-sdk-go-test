package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

var mockSession = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("mock-region"),
}))

func TestGetInstanceID(t *testing.T) {
	expectedInstanceID := "i-1234567890ab"
	metadataInstanceID := "i-1234567890ab"

	mux := http.NewServeMux()
	mux.HandleFunc("/latest/meta-data/instance-id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, metadataInstanceID)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	mockEC2MetaDataSvc := ec2metadata.New(mockSession, &aws.Config{
		Endpoint: aws.String(server.URL),
	})

	instanceID, err := GetInstanceID(mockEC2MetaDataSvc)
	if err != nil {
		t.Errorf("expected no error, but got %v.", err)
	}
	if instanceID != expectedInstanceID {
		t.Errorf("expected value of tag: %s, but got tag value: %v.", expectedInstanceID, instanceID)
	}
}
