package internal


type MailService interface {
	SendMail(to, subject, message string) (string, error)
}

