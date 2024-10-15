package email

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Create(ctx context.Context, dto EmailNotificationDTO)
}

type db struct {
	connect *pgx.Conn
	logger  *logrus.Logger
}

func NewRepository(connect *pgx.Conn, logger *logrus.Logger) Repository {
	return &db{
		connect: connect,
		logger:  logger,
	}
}

func (d *db) Create(ctx context.Context, dto EmailNotificationDTO) {
	query := `INSERT INTO email_notification (uuid, status, email) VALUES ($1, $2, $3)`

	res, err := d.connect.Exec(ctx, query, dto.uuid, dto.status, dto.email)
	if err != nil {
		d.logger.Errorf("error creating record: %s", err)
		return
	}

	if res.RowsAffected() == 0 {
		d.logger.Errorf("no rows affected")
	}
}
