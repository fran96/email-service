package internal

import (
	"github.com/mailgun/mailgun-go"
)

type mailGunService struct {
	Domain string
	ApiKey string
	From string
}

func NewMailGunService(config *Config) *mailGunService {
	return &mailGunService{
		Domain: config.MailGun.Domain,
		ApiKey: config.MailGun.ApiKey,
		From: config.MailGun.From,
	}
}

func(mg *mailGunService) SendMail(to, subject, message string) (string, error) {
	mgInstance := mailgun.NewMailgun(mg.Domain, mg.ApiKey)
	m := mgInstance.NewMessage(
	  to,
	  subject,
	  message,
	  mg.From,
	)
	_, id, err := mgInstance.Send(m)
	if err != nil {
		return "", err
	}
	return id, err
}
