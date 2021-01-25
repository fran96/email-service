package internal


type Config struct {
    MailGun struct {
		Domain string `yaml:"domain"`
		ApiKey string `yaml:"apiKey"`
		From string `yaml:"from"`
	} `yaml:"mailgun"`
	SendGrid struct {
		ApiKey string `yaml:"apiKey"`
		From string `yaml:"from"`
	} `yaml:"sendgrid"`
}


type Mail struct {
	To string
	Subject string
	Message string
}
