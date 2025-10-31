package email

import (
	"github.com/rugi123/email-service/internal/config"
	"github.com/rugi123/email-service/internal/models"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	dialer *gomail.Dialer
}

func NewSender(cfg config.SMTP) *Sender {
	d := gomail.NewDialer(cfg.Host, cfg.Port, cfg.Password, cfg.Password)
	return &Sender{
		dialer: d,
	}
}

func (s *Sender) Send(letter models.Letter) error {
	m := gomail.NewMessage()
	m.SetHeader("From", letter.Sender)
	m.SetHeader("To", letter.Receiver)
	m.SetHeader("Subject", letter.Subject)
	m.SetBody("text/plain", letter.Body)

	err := s.dialer.DialAndSend(m)

	return err
}
