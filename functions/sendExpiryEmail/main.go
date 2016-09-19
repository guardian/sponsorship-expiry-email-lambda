package main

import (
	//"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	dynamo := dynamodb.New(session.New(&aws.Config{
		Region: aws.String("eu-west-1"),
	}))
	LoadExpiringSoon(dynamo)

	//apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
	//	dynamo := dynamodb.New(session.New())
	//
	//	LoadExpiringSoon(dynamo)
	//
	//	return m, nil
	//})
}