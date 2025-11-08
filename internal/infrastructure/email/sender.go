package email

import (
	"crypto/tls"

	"github.com/rugi123/email-service/internal/config"
	"github.com/rugi123/email-service/internal/domain/models"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	dialer *gomail.Dialer
}

func NewSender(cfg config.SMTP) *Sender {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Username, cfg.Password)
	d.SSL = false
	d.TLSConfig = &tls.Config{
		ServerName: cfg.Host,
	}
	return &Sender{
		dialer: d,
	}
}

func (s *Sender) Send(email models.Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.dialer.Username)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/plain", email.Body)

	err := s.dialer.DialAndSend(m)

	return err
}
