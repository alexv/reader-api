package api

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Connect initializes a connection with AWS
func Connect() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Set the AWS Region that the service clients should use
	cfg.Region = endpoints.UsWest2RegionID

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.New(cfg)

	// Build the request with its input parameters
	req := svc.DescribeTableRequest(&dynamodb.DescribeTableInput{
		TableName: aws.String("myTable"),
	})

	// Send the request, and get the response or error back
	resp, err := req.Send()
	if err != nil {
		panic("failed to describe table, " + err.Error())
	}

	fmt.Println("Response", resp)
}
