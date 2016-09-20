package main

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/aws"
	"time"
)

func sendEmail(messageBody string, sesClient *ses.SES) error {

	date := time.Now().Format("Mon, 02 Jan 2006")

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String("steppenwells@gmail.com"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data:    aws.String(messageBody),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String("Expiring Advertisement Features " + date),
			},
		},
		Source: aws.String("stephen.wells@guardian.co.uk"),
	}

	_, err := sesClient.SendEmail(params)

	return err
}
