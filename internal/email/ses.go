package email

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/aws/aws-sdk-go/aws/session"
)

type SimpleEmailSender struct {
	sesClient *ses.SES
}

func NewSimpleEmailSender() (*SimpleEmailSender, error) {
	awsSession, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf("impossible to create aws session: %w", err)
	}
	sesClient := ses.New(awsSession)

	return &SimpleEmailSender{
		sesClient: sesClient,
	}, nil
}

func (s *SimpleEmailSender) Send(in SendInput) error {
	input := ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				&in.ToAddress,
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    &in.HtmlBody,
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    &in.TextBody,
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    &in.Subject,
			},
		},
		Source: &in.FromAddress,
	}
	log.Println(input.String())
	_, err := s.sesClient.SendEmail(&input)
	if err != nil {
		return fmt.Errorf("impossible to call SendEmail: %w", err)
	}
	return nil
}
