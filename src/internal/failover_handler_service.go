package internal

import "errors"

type FailOverHandlerService struct {
	mailServices []MailService
}


func NewFailOverHandlerService(mailServices ...MailService) *FailOverHandlerService {
	return &FailOverHandlerService{
		mailServices: mailServices,
	}
}

func(f *FailOverHandlerService) SendMail(to, subject, message string) (string, error) {
	var err error
	for _, mailService := range f.mailServices {
		res, err3 := mailService.SendMail(to, subject, message)
		err = err3
		if err == nil {
			return res, err
		}
	}
	return "", errors.New("Mail service is down")
}
