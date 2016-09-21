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
				aws.String("graham.hayday@theguardian.com"),
				aws.String("stephen.tolfree@theguardian.com"),
				aws.String("emma.wevill@theguardian.com"),
				aws.String("tim.hughes@theguardian.com"),
				aws.String("marcus.browne@theguardian.com"),
				aws.String("matthew.caines@theguardian.com"),
				aws.String("guardian.adops@theguardian.com"),
			},
			CcAddresses: []*string{
				aws.String("commercial.dev@theguardian.com"),

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
		Source: aws.String("commercial.dev@theguardian.com"),
	}

	_, err := sesClient.SendEmail(params)

	return err
}
