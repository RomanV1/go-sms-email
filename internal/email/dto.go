package email

import "github.com/google/uuid"

type EmailNotificationDTO struct {
	uuid   uuid.UUID
	status uint32
	email  string
}
