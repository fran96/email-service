package main

import (
	"log"
	"os"

	"github.com/fran96/email-service/src/internal"

	"gopkg.in/yaml.v2"
)

func main() {
	cfg := loadConfig()

	mailGunMailService := internal.NewMailGunService(cfg)
	sendGridMailService := internal.NewSendGridService(cfg)
	failOverService := internal.NewFailOverHandlerService(mailGunMailService, sendGridMailService)

	controller := internal.NewController(failOverService)
	controller.HttpRoutes()
}

func loadConfig() *internal.Config {
	f, err := os.Open("./config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg internal.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
