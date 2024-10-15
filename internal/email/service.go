package email

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateEmailNotification(ctx context.Context, status uint32, email string)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateEmailNotification(ctx context.Context, status uint32, email string) {
	s.repo.Create(ctx, EmailNotificationDTO{
		uuid:   uuid.New(),
		status: status,
		email:  email,
	})
}
