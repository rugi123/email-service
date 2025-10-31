package main

import (
	"fmt"

	"github.com/rugi123/email-service/internal/config"
	"github.com/rugi123/email-service/internal/email"
	"github.com/rugi123/email-service/internal/models"
)

func main() {
	cfg, err := config.Load("internal/config/config.yaml")
	fmt.Println(cfg, err)

	sender := email.NewSender(cfg.SMTPConfig)
	letter := models.Letter{
		Receiver: "lik500mlg@mail.ru",
		Subject:  "тестовое письмо",
		Body:     "тестовое письмо",
	}
	err = sender.Send(letter)
	fmt.Println(err)
}
