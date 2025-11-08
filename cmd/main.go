package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/rugi123/email-service/internal/config"
	"github.com/rugi123/email-service/internal/domain/services"
	"github.com/rugi123/email-service/internal/infrastructure/email"
	"github.com/rugi123/email-service/internal/infrastructure/nats"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	js, err := nats.NewJetStream()
	if err != nil {
		log.Fatalf("jetstream err: %v", err)
	}

	uuid := uuid.New()

	worker, err := nats.NewWorker(*js, "EMAILS", uuid.String())
	if err != nil {
		log.Fatalf("worker err: %v", err)
	}

	sender := email.NewSender(cfg.SMTPConfig)

	es := services.NewEmailService(sender, worker)

	err = es.ProcessEmail()
	if err != nil {
		log.Fatalf("process email err: %v", err)
	}
}
