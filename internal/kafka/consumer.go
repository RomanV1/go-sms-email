package kafka

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/RomanV1/go-sms-email/internal/message"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer interface {
	ConsumeMessages()
}

type consumer struct {
	consumer  *kafka.Reader
	formatter message.Formatter
}

func NewConsumer(brokers []string, topic, groupID string, formatter message.Formatter) Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.LastOffset,
	})

	return &consumer{consumer: r, formatter: formatter}
}

func (c consumer) ConsumeMessages() {
	defer c.consumer.Close()

	ctx := context.Background()

	for {
		msg, err := c.consumer.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				log.Println("Timeout exceeded while reading msg")
				break
			}
			log.Fatalf("Error reading msg: %v", err)
		}

		c.formatter.HandleMessage(binary.BigEndian.Uint32(msg.Key), string(msg.Value))
	}
}
