package internal

import (
	"errors"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGridService struct {
	ApiKey string
	From string
}

func NewSendGridService(config *Config) *sendGridService {
	return &sendGridService{
		ApiKey: config.SendGrid.ApiKey,
		From: config.SendGrid.From,
	}
}

func(sg *sendGridService) SendMail(to, subject, message string) (string, error) {
	sender := mail.NewEmail(subject, sg.From)
	receipient := mail.NewEmail(subject, to)
	msg := mail.NewSingleEmail(sender, subject, receipient, message, "")
	client := sendgrid.NewSendClient(sg.ApiKey)
	response, err := client.Send(msg)
	if err != nil {
		return "", err
	} else if response.StatusCode < 200 || response.StatusCode > 299 { //Reason for this check: send grid sometimes responds with an error in response.Body instead of a returning normal error.
		return "", errors.New(response.Body)
	} else {
		return response.Body, nil
	}
}
