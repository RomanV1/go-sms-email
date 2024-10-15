package email

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

type Sender struct {
	service Service
}

func NewSender(service Service) *Sender {
	return &Sender{
		service: service,
	}
}

func (s *Sender) SendEmail(status uint32, email, message string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", message)

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		fmt.Printf("Error converting port: %v\n", err)
		return
	}

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Error sending email %s", err)
		return
	}

	s.service.CreateEmailNotification(context.Background(), status, email)
}
