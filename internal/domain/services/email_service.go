package services

import (
	"fmt"

	"github.com/rugi123/email-service/internal/domain/models"
)

type EmailSender interface {
	Send(email models.Email) error
}

type Worker interface {
	Start() (chan models.Email, error)
}

type EmailService struct {
	sender EmailSender
	worker Worker
}

func NewEmailService(sender EmailSender, worker Worker) *EmailService {
	return &EmailService{
		sender: sender,
		worker: worker,
	}
}

func (s EmailService) ProcessEmail() error {
	ch, err := s.worker.Start()
	if err != nil {
		return err
	}
	for {
		email := <-ch
		fmt.Println("sended email: ", email)
		err = s.sender.Send(email)
		if err != nil {
			fmt.Println(err)
		}
	}
}
