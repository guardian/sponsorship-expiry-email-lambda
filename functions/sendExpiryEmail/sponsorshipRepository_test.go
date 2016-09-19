package main

import (
	"testing"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
)

var dynamo = dynamodb.New(session.New(&aws.Config{
	Region: aws.String("eu-west-1"),
}))

func TestLoadExpiringSoon(t *testing.T) {
	spons, err := LoadExpiringSoon(dynamo)

	if (err != nil) {
		t.Error("error getting sponsorships", err)
		return
	}

	if (len(spons) == 0) {
		t.Error("expected to see some results, you may have just got unlucky")
	}

	fmt.Println("expiring soon:")
	for _, s := range spons {
		fmt.Println(s.SponsorName)
	}
}

func TestLoadExpiredRecently(t *testing.T) {
	spons, err := LoadExpiredRecently(dynamo)

	if (err != nil) {
		t.Error("error getting sponsorships", err)
		return
	}

	if (len(spons) == 0) {
		t.Error("expected to see some results, you may have just got unlucky")
	}

	fmt.Println("expired recently:")
	for _, s := range spons {
		fmt.Println(s.SponsorName)
	}
}
