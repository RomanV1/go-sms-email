package message

import (
	"fmt"
	"github.com/RomanV1/go-sms-email/internal/email"
)

const (
	CreatedAccount  uint32 = 1
	UpdatedAccount  uint32 = 2
	VerifiedAccount uint32 = 3
)

type Formatter interface {
	HandleMessage(status uint32, email string)
}

type formatter struct {
	sender *email.Sender
}

func NewFormatter(sender *email.Sender) Formatter {
	return &formatter{
		sender: sender,
	}
}

func (m *formatter) HandleMessage(status uint32, email string) {
	switch status {
	case CreatedAccount:
		m.sendAccountCreatedEmail(status, email)
	case UpdatedAccount:
		m.sendAccountUpdatedEmail(status, email)
	case VerifiedAccount:
		m.sendAccountVerifiedEmail(status, email)
	default:
		fmt.Println("Unknown status")
	}
}

func (m *formatter) sendAccountCreatedEmail(status uint32, email string) {
	emailContent := fmt.Sprintf("Your account has been successfully created.\n\nBest regards,\nSupport Team.")
	m.sender.SendEmail(status, email, emailContent)
}

func (m *formatter) sendAccountUpdatedEmail(status uint32, email string) {
	emailContent := fmt.Sprintf("Your account has been successfully updated.\n\nBest regards,\nSupport Team.")
	m.sender.SendEmail(status, email, emailContent)
}

func (m *formatter) sendAccountVerifiedEmail(status uint32, email string) {
	emailContent := fmt.Sprintf("Your account has been successfully verified.\n\nBest regards,\nSupport Team.")
	m.sender.SendEmail(status, email, emailContent)
}
