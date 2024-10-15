package main

import (
	"context"
	"github.com/RomanV1/go-sms-email/internal/email"
	"github.com/RomanV1/go-sms-email/internal/kafka"
	"github.com/RomanV1/go-sms-email/internal/message"
	"github.com/RomanV1/go-sms-email/pkg/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	loadEnv()

	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	client, err := postgres.NewClient(context.Background(), os.Getenv("PG_USERNAME"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_DATABASE"))
	if err != nil {
		logger.Fatalf("Failed to initialize PostgreSQL client: %v", err)
		return
	}

	repo := email.NewRepository(client, logger)

	service := email.NewService(repo)

	sender := email.NewSender(service)

	formatter := message.NewFormatter(sender)

	consumer := kafka.NewConsumer([]string{os.Getenv("KAFKA_BROKER")}, os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), formatter)
	consumer.ConsumeMessages()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
